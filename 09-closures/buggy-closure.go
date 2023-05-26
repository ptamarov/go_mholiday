package main

import "fmt"

func badClosure() {

	fmt.Println("*** badClosure says ***")
	s := make([]func(), 4)

	for i := 0; i < 4; i++ {
		s[i] = func() {
			fmt.Printf("%d @ %p\n", i, &i)
		}
	}

	for j := 0; j < 4; j++ {
		s[j]() // <- all calls will refer to last i in loop
	}

	fmt.Println("*** end badClosure ***")
}
