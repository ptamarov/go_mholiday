package search

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

type Comic struct {
	Day        string `json:"day"`
	Month      string `json:"month"`
	Year       string `json:"year"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Image      string `json:"img"`
	SafeTitle  string `json:"safe_title"`
	Title      string `json:"title"`
	News       string `json:"news"`
}

const wait = 0 // milliseconds

func Search(pathToFolder string, keywords []string) {

	// Make search case insensitive.
	for i := range keywords {
		keywords[i] = strings.ToLower(keywords[i])
	}

	// Get a list of all files in pathToFolder (index)
	items, err := os.ReadDir(pathToFolder)
	if err != nil {
		log.Fatalf("Error %s while accessing index.\n", err)
		os.Exit(-1)
	}

	for _, item := range items {
		var comic Comic
		time.Sleep(wait * time.Millisecond)

		// Open .json file for comic.
		pathToFile := item.Name()
		comicInfo, err := os.ReadFile(pathToFolder + pathToFile)
		if err != nil {
			fmt.Printf("Error while opening %s: %s.\n", pathToFile, err)
		}

		// Dump into comic variable.
		err = json.Unmarshal(comicInfo, &comic)
		if err != nil {
			fmt.Printf("Error while unmarshalling %s: %s.\n", pathToFile, err)
		}

		// Make search case insensitive.
		title := strings.ToLower(comic.Title)
		transcript := strings.ToLower(comic.Transcript)
		number := comic.Num

		// Check title for keywords.
		titleHit := checkHit(title, keywords)
		if titleHit {
			fmt.Printf("Hit! Comic #%d contains all keywords in title: %s.\n", number, strings.Join(keywords, ", "))
			fmt.Printf("URL is https://xkcd.com/%d\n", number)
		}

		// Check transcript for keywords.
		transcriptHit := checkHit(transcript, keywords)
		if transcriptHit {
			fmt.Printf("Hit! Comic #%d contains all keywords in transcript: %s.\n", number, strings.Join(keywords, ", "))
			fmt.Printf("URL is https://xkcd.com/%d\n", number)
		}
	}
}

func checkHit(text string, keywords []string) bool {
	for _, keyword := range keywords {
		if !strings.Contains(text, keyword) {
			return false
		}
	}
	return true
}

func ProcessTranscript(transcript string) string {

	transcriptPlusAlt := strings.Split(transcript, "\n")

	fmt.Println(len(transcriptPlusAlt))
	if len(transcriptPlusAlt) != 2 {
		fmt.Println("Transcript with unexpected format.")
		return ""
	}

	t := transcriptPlusAlt[0] // "[[texto]]"
	a := transcriptPlusAlt[1] // ""

	captureTranscript := regexp.MustCompile(`^\[\[.*\]\]$`)
	captureAlt := regexp.MustCompile(`^\{\{Alt(-title)?(.*)\}\}$`)

	t = captureTranscript.FindAllString(t, -1)[1]
	a = captureAlt.FindAllString(a, -1)[2]

	return t + "\n" + a
}
