package main

import (
	"encoding/json"
	"fmt"
	"log"
)

var j1 = `{
    "item": "album",
    "album" : {"title": "Dark Side of the Moon"}
}`

var j2 = `{
    "item": "song",
    "song" : {"title": "Bella Donna", "artist": "Stevie Nicks"}
}`

func main() {
	var resp1, resp2 respWraper
	var err error

	if err = json.Unmarshal([]byte(j1), &resp1); err != nil { // uses custom method UnmarshalJSON
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", resp1.response)

	if err = json.Unmarshal([]byte(j2), &resp2); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", resp2.response)

}
