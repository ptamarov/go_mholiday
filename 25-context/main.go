package main

import (
	"context"
	"log"
	"net/http"
	"runtime"
	"time"
)

// A context is usually a tree.
// Has a root context. Nodes are created and are inmutable.
// Subtrees are also contexts, and can have shorter (but not longer!) timeouts.
// Timeout will apply only to this subtree, not to anything above it.

type result struct {
	url     string
	err     error
	latency time.Duration
}

func main() {
	list := []string{
		"https://amazon.com",
		"https://google.com",
		"https://stackoverflow.com",
		"https://youtube.com",
		"https://wsj.com",
		"https://nytimes.com",
		"http://localhost:8000",
	}

	r, _ := first(context.Background(), list)

	if r.err != nil {
		log.Printf("%-20s %s\n", r.url, r.err)
	} else {
		log.Printf("%-20s %s\n", r.url, r.latency)
	}

	time.Sleep(9 * time.Second)
	log.Println("Quit anyway...", runtime.NumGoroutine(), "still running.")
}

// Convention: pass Context as first parameter (usually ctx).
func first(ctx context.Context, urls []string) (*result, error) {
	results := make(chan result, len(urls))
	// BAD: unbuffered channel, other coroutines will get stuck. No one will be there to read.
	// Memory leak -> memory which is no longer needed is not released.
	// Solution: buffer to avoid leaking; coroutines will send to channel and terminate.

	ctx, cancel := context.WithCancel(ctx)
	defer cancel() // Must cancel even if a timeout does not happen to release resources.

	// Start one coroutine for each url.
	for _, url := range urls {
		go get(ctx, url, results)
	}

	select {
	case r := <-results:
		return &r, nil // Will trigger defered cancel.
	case <-ctx.Done():
		// It could happen the context you received is done before cancel is issued.
		// Responsible to keep track of inherited contexts and their cancelations.
		return nil, ctx.Err()
	}
}

func get(ctx context.Context, url string, ch chan<- result) { // only write to the channel
	var r result

	start := time.Now()
	ticker := time.NewTicker(1 * time.Second).C
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if resp, err := http.DefaultClient.Do(req); err != nil {
		r = result{url, err, 0}
	} else {
		t := time.Since(start).Round(time.Millisecond)
		r = result{url, nil, t}
		resp.Body.Close()
	}

	for {
		select {
		case ch <- r: // Ready to send. Coroutine ends.
			return
		case <-ticker:
			log.Println("tick", r)
		}
	}

}
