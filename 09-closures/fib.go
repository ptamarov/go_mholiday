package main

import "fmt"

func fib() func() int {
	a, b := 0, 1

	return func() int {
		a, b = b, a+b
		return b
	}
}

func callFib() {
	fmt.Println("*** callFib says ***")
	f, g, h := fib(), fib(), fib()

	for x := h(); x < 100; x = h() {
		fmt.Print(fmt.Sprint(x) + ", ")
	}
	fmt.Println()

	// a, b different for f, g and h (different environment pointers)
	fmt.Println(f(), f(), f(), f())
	fmt.Println(g(), g(), g(), g())
	fmt.Println("*** end callFib ***")

}
