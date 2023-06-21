package main

import "fmt"

func main() {
	messages := make(chan string, 2)

	messages <- "Buffered"
	messages <- "channel!"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

	fmt.Println("*****************")

	badExample() // all false: receive ends before sends finishes, no time to set bools to true

	fmt.Println("*****************")

	alsoBadExample() // all true: sends are not blocking, so sleep time allows for modification (but still bad)
}
