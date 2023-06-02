package main

import (
	"fmt"
	"log"
	"net/http"
)

type nextChan chan int

func (ch nextChan) handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>You got %d</h1>", <-ch)
	// Each request will increase a counter by 1.
}

func feedToChannel() {
	var nextID nextChan = make(chan int)
	go func() {
		for i := 0; ; i++ {
			nextID <- i
		}
	}()
	go http.HandleFunc("/", nextID.handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
