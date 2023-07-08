package main

import "fmt"

func main() {
	fmt.Println(sum(1, 4, 10))

	s := []int{1, 4, 10}
	fmt.Println(sum(s...)) // unpack operator

}
