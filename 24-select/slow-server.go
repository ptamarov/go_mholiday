package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Create a slow server to test the new implementation of Section 23.
func slow(w http.ResponseWriter, r *http.Request) {
	log.Println("Waiting...")
	time.Sleep(7 * time.Second)
	fmt.Fprintf(w, "Sorry, I am slow!")
	log.Println("Responded...")
}

func startSlowServer() {
	http.HandleFunc("/", slow)

	err := http.ListenAndServe("localhost:8000", nil)

	if err != nil {
		log.Fatal(err)
	}

}
