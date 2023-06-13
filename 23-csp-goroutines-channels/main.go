package main

import "log"

// Constrol what executes.
const doFeed = false
const doSieve = false

func main() {
	// Example 1.
	list := []string{
		"https://amazon.com",
		"https://google.com",
		"https://stackoverflow.com",
		"https://youtube.com",
		"https://wsj.com",
		"https://nytimes.com",
		"http://localhost:8000",
	}

	ping(list)
	log.Println("*** Done with pings. ***")

	// Example 2.
	if doFeed {
		feedToChannel()
	}

	// Example 3.
	if doSieve {
		sieve(100)
	}
}
