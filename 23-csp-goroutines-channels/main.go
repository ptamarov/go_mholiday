package main

import "log"

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
	feedToChannel()
	// This does not stop, so coming lines will not execute.

	// Example 3.
	sieve(100)
}
