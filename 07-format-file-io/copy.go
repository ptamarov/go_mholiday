package main

import (
	"fmt"
	"io"
	"os"
)

func copyExample() {

	for _, filename := range os.Args[1:] {
		file, err := os.Open(filename)

		if err != nil {
			fmt.Fprint(os.Stderr, err)
			continue // ignore this file if there was an error when opening it
		}

		if _, err := io.Copy(os.Stdout, file); err != nil {
			fmt.Fprint(os.Stderr, err)
			continue

		}

		file.Close()

	}
}
