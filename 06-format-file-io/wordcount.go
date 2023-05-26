package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func wordCount() {
	var tcc, twc, tlc int

	for _, filename := range os.Args[1:] {
		var cc, wc, lc int

		file, err := os.Open(filename)

		if err != nil {
			fmt.Fprint(os.Stderr, err)
			continue // ignore this file if there was an error when opening it
		}

		scan := bufio.NewScanner(file)

		for scan.Scan() {
			s := scan.Text()

			wc += len(strings.Fields(s))
			cc += len(s)
			lc++ // Scan() goes through each line, so just count each iteration
		}

		tcc += cc
		tlc += lc
		twc += wc

		fmt.Printf(" %7d %7d %7d \t %s\n", lc, wc, cc, filename)
		file.Close()

	}
	fmt.Printf(" %7d %7d %7d \t %s\n", tlc, twc, tcc, "total")

}
