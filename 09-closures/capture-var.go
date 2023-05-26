package main

import "fmt"

func captureVar() {

	fmt.Println("*** captureVar says ***")
	s := make([]func(), 4)

	for i := 0; i < 4; i++ {
		j := i // capture var
		s[i] = func() {
			fmt.Printf("%d @ %p\n", j, &j)
		}
	}

	for j := 0; j < 4; j++ {
		s[j]()
	}

	fmt.Println("*** end captureVar ***")
}
