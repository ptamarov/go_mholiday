package main

import (
	"fmt"
	"time"
)

func Example1() {
	chans := []chan int{
		make(chan int),
		make(chan int),
	}

	for i := range chans {
		go func(j int, ch chan<- int) {
			for {
				time.Sleep(time.Duration(j) * time.Second)
				ch <- j
			}
		}(i+1, chans[i])
	}

	for i := 0; i < 12; i++ {
		select {
		case a := <-chans[0]:
			fmt.Println(a)
		case a := <-chans[1]:
			fmt.Println(a)
		}
	}
}
