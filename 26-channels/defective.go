package main

import (
	"fmt"
	"time"
)

type T struct {
	i byte
	b bool
}

func send(i int, ch chan<- *T) {
	t := &T{i: byte(i)} // passing the address, not the object
	ch <- t

	t.b = true // UNSAFE AT ANY SPEED, race condition
	// once you give something to a channel
	// you have GIVEN UP OWNERSHIP of that object
	// do not modify it
}

func badExample() {
	vs := make([]T, 5)
	ch := make(chan *T) // unbuffered, rendezvous behaviour

	for i := range vs {
		go send(i, ch)
	}

	time.Sleep(1 * time.Second) // all goroutines started

	for i := range vs {
		vs[i] = *<-ch
	}

	for _, v := range vs {
		fmt.Println(v)
	}
}

func alsoBadExample() {
	vs := make([]T, 5)
	ch := make(chan *T, 5) // buffered

	for i := range vs {
		go send(i, ch)
	}

	time.Sleep(1 * time.Second) // all goroutines started

	for i := range vs {
		vs[i] = *<-ch
	}

	for _, v := range vs {
		fmt.Println(v)
	}
}
