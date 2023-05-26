package main

import (
	"fmt"
	"os"

	"github.com/ptamarov/go_mholiday/02-simple-example/hello"
)

func main() {
	fmt.Println(hello.Say(os.Args[1:]))
}
