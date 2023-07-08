package main

import (
	"fmt"
	"sync"
)

func do() int {
	var m sync.Mutex
	var n int64
	var w sync.WaitGroup

	for i := 0; i < 1000; i++ {
		w.Add(1)

		go func() {
			m.Lock() // More efficient than using a counting semaphore with a channel.
			n++      // DATA RACE
			m.Unlock()
			w.Done()
		}()
	}

	w.Wait()

	return int(n)
}

func main() {
	fmt.Println(do())
	fmt.Println(doAtomic())
}
