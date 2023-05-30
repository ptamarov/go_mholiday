package main

import (
	"fmt"
	"os"

	"github.com/ptamarov/go_mholiday/16-gopl-exercise/search"
)

const folderPath = "index/"

func main() {
	keywords := os.Args[1:]

	if len(keywords) == 0 {
		fmt.Println("No keywords given.")
		os.Exit(0)
	}

	search.Search(folderPath, keywords)
}
