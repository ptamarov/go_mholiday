package main

import (
	"sync"
	"sync/atomic"
)

func doAtomic() int {
	var n int64
	var w sync.WaitGroup

	for i := 0; i < 1000; i++ {
		w.Add(1)

		go func() {
			atomic.AddInt64(&n, 1) // Done at the hardware level.
			w.Done()
		}()
	}

	w.Wait()

	return int(n)
}
