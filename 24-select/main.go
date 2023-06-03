package main

import (
	"log"
	"time"
)

const serve = false

func main() {
	// Example 1.
	if serve {
		startSlowServer()
	}

	// Example 2.
	const tickRate = 2 * time.Second

	stopper := time.After(5 * tickRate)
	ticker := time.NewTicker(tickRate).C

	log.Println("*** Start ***")
loop:
	for {
		select {
		case <-ticker:
			log.Println("Tick!")
		case <-stopper:
			break loop // Good style to name loop while running.
		}
	}

	log.Println("*** Finish ***")
}
