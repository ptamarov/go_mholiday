package makeindex

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func GetComicFromNumber(comicNumber int) ([]byte, int) {
	var url = fmt.Sprintf("https://xkcd.com/%d/info.0.json", comicNumber)
	var body []byte

	resp, err := http.Get(url)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}

	defer resp.Body.Close()

	switch resp.StatusCode {

	case http.StatusOK:
		body, err = io.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}
		return body, 200

	default:
		return nil, resp.StatusCode

	}
}

func CreateIndex(upperLimit int) int {
	var status int
	var newEntry []byte

	number := 1
	for status == 200 && number < upperLimit+1 {
		newEntry, status = GetComicFromNumber(number)

		switch status {
		case 200:
			err := ioutil.WriteFile(fmt.Sprintf("./comic-index/comic-%d.json", number), newEntry, 0644)

			if err != nil {
				fmt.Printf("(!) Error %s while storing comic number %d. Moving on...\n", err, number)
			}
		case 404:
			fmt.Println("Error 404. Moving on...")

		default:
			// stop if any other kind of error is found.
			fmt.Printf("Error while fetching information for comic #%d. Status %d\n", number, status)
			return number
		}
		number++
	}
	return number
}
