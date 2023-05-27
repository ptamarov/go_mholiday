package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var rawHTML = ``

func main() {
	doc, err := html.Parse(bytes.NewReader([]byte(rawHTML)))

	if err != nil {
		fmt.Fprintf(os.Stderr, "parse failed: %s\n", err)
		os.Exit(-1)
	}

	word, pics := countWordsAndImages(doc)

	fmt.Printf("%d words and %d images.\n", word, pics)

}

func countWordsAndImages(doc *html.Node) (int, int) {
	var words, pics int

	visit(doc, &words, &pics)

	return words, pics
}

func visit(node *html.Node, toWords, toPics *int) {

	switch t := node.Type; t {

	case html.ElementNode:
		switch d := node.Data; d {
		case "script", "style": // stop at scripts or styles
			return
		case "img":
			*toPics++
		}

	case html.TextNode:
		*toWords += len(strings.Fields(node.Data))

	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		visit(c, toWords, toPics)
	}
}
