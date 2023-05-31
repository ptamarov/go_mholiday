package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ptamarov/go_mholiday/16-gopl-mh-solution/mhsearch"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "no file given")
		os.Exit(-1)
	}

	fn := os.Args[1]

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "no search terms")
		os.Exit(-1)
	}

	var (
		items []mhsearch.ComicInfo
		terms []string
		input io.ReadCloser
		cnt   int
		err   error
	)

	if input, err = os.Open(fn); err != nil {
		fmt.Fprintf(os.Stderr, "bad file %s\n", err)
		os.Exit(-1)
	}

	// Decode the file.
	if err = json.NewDecoder(input).Decode(&items); err != nil {
		fmt.Fprintf(os.Stderr, "bad decoding %s\n", err)
		os.Exit(-1)
	}

	fmt.Fprintf(os.Stderr, "read %d comics\n", len(items))

	// Get search terms.
	for _, t := range os.Args[2:] {
		terms = append(terms, strings.ToLower(t))
	}

	fmt.Printf("Looking up: %s", terms)

	// Search.
	cnt = mhsearch.Search(terms, items)

	fmt.Fprintf(os.Stderr, "found %d matches\n", cnt)
}
