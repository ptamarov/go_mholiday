package main

import (
	"fmt"
	"os"

	"github.com/ptamarov/go_mholiday/simple-example/hello"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println(hello.Say(os.Args[1]))
	} else {
		fmt.Println(hello.Say("world!"))
	}
}
