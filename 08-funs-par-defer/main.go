package main

import "fmt"

var t = Tree{
	Left: &Tree{
		Left:  nil,
		Value: 1,
		Right: &Tree{
			Left:  nil,
			Value: 2,
			Right: nil,
		}},
	Value: 3,
	Right: &Tree{
		Left:  nil,
		Value: 4,
		Right: nil,
	},
}

func main() {
	fmt.Println(walkTree(&t)) // ((1 (2)) 3 (4))
	fmt.Println(do())         // 2
}
