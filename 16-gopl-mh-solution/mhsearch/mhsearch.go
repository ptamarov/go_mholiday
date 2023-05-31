package mhsearch

import (
	"fmt"
	"strings"
)

type ComicInfo struct {
	Num        int    `json:"num"`
	Day        string `json:"day"`
	Month      string `json:"month"`
	Year       string `json:"year"`
	Title      string `json:"title"`
	Transcript string `json:"transcript"`
}

func Search(terms []string, items []ComicInfo) int {
	var cnt int

	for _, item := range items {
		title := strings.ToLower(item.Title)
		transcript := strings.ToLower(item.Transcript)

		if checkMatches(title, terms) || checkMatches(transcript, terms) {
			fmt.Printf("https://xkcd.com/%d/ %s/%s/%s %s\n",
				item.Num,
				item.Month,
				item.Day,
				item.Year,
				item.Title)
			cnt++
		}

	}

	return cnt
}

func checkMatches(text string, terms []string) bool {
	for _, term := range terms {
		if !strings.Contains(text, term) {
			return false
		}
	}
	return true
}
