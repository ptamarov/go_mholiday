package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

const portNum = ":8080"

type Todo struct {
	UserID    int    `json:"userID"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var form = `
<h1>Todo #{{ .ID }}</h1>
<div>{{ printf "User %d" .UserID }}</div>
<div>{{ printf "%s (completed: %t)" .Title .Completed }}</div>
`

func main() {
	Server()
	// ClientProgram()
}

func handler(w http.ResponseWriter, r *http.Request) {
	const base = "https://jsonplaceholder.typicode.com/"

	target := base + r.URL.Path[1:]
	log.Println("Trying to read from: ", target)

	resp, err := http.Get(target)

	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	defer resp.Body.Close()

	// body, err := __ioutil.ReadAll(resp.Body) // Read the body of the response.
	// Superseded by line 48.

	var item Todo

	if err := json.NewDecoder(resp.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.New("MyTemplate")
	tmpl.Parse(form)
	tmpl.Execute(w, item)

}

func Server() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(portNum, nil)

	if err != nil {
		log.Fatal(err)
	}
}

func ClientProgram() {
	const url = "https://jsonplaceholder.typicode.com"

	resp, err := http.Get(url + "/todos/1")

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}

		var item Todo

		err = json.Unmarshal(body, &item)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}

		fmt.Printf("%#v\n", item) // %#v is a Go-syntax representation of the value

	}
}
