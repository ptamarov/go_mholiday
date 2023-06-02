package main

import "fmt"

func generator(limit int, src chan<- int) {
	for i := 2; i < limit; i++ {
		src <- i
	}
	close(src)

}

type TestNumber int

type LabelledChan struct {
	channel *(chan int)
	label   int
}

func (lc LabelledChan) String() string {
	return fmt.Sprintf("Channel for number %d.", lc.label)
}

func (prime TestNumber) filter(src <-chan int, dst chan<- int, labelsrc int, labeldst int) {
	// Will block for generator.
	fmt.Printf("Filter for %d is checking with src %d and dst %d.\n", prime, labelsrc, labeldst)
	for i := range src {
		fmt.Printf("Filter for %d is checking %d coming from src %d.\n", prime, i, labelsrc)
		if i%int(prime) != 0 {
			fmt.Printf("Filter for %d found %d coming from src %d.\n", prime, i, labelsrc)
			dst <- i
		}
	}
	close(dst)
}

func sieve(limit int) {
	src := make(chan int)

	labelledSrc := LabelledChan{&src, 0}

	go generator(limit, src)

	for {
		prime, ok := <-*labelledSrc.channel

		if !ok {
			fmt.Printf("Src channel for number %d is closed.", labelledSrc.label)
			break
		}

		dst := make(chan int)
		labelledDst := LabelledChan{&dst, prime}

		f := TestNumber(prime).filter
		go f(*labelledSrc.channel, *labelledDst.channel, labelledSrc.label, labelledDst.label)
		labelledSrc = labelledDst // What is going on here?
		fmt.Println(prime)
	}

}
