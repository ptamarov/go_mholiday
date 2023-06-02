package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

//!+main

func main() {
	hostAddress := "localhost:8000"
	db := &database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/fetch", db.fetch)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/delete", db.delete)

	err := http.ListenAndServe(hostAddress, nil)
	if err != nil {
		log.Fatal(err)
	}
}

//!-main

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: \t %s\n", item, price)
	}
}

func (db database) fetch(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	if price, ok := db[item]; ok {
		msg := fmt.Sprintf("%s costs %s per unit\n", item, price)
		http.Error(w, msg, http.StatusOK)
		return
	}

	msg := fmt.Sprintf("no such item: %q\n", item)
	http.Error(w, msg, http.StatusNotFound) // 404

}

func (todb *database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	if _, ok := (*todb)[item]; !ok {
		msg := fmt.Sprintf("ERROR: no such item %s.\n", item) // 400
		http.Error(w, msg, http.StatusNotFound)
		return
	}

	newprice, err := strconv.ParseFloat(price, 32)

	if err != nil {
		msg := fmt.Sprintf("ERROR: %s is an invalid price.\n", price) // 400
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	(*todb)[item] = dollars(newprice)
	fmt.Fprintf(w, "New price for item %s is %s.\n", item, (*todb)[item])
}
func (todb *database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	if _, ok := (*todb)[item]; ok {
		msg := fmt.Sprintf("ERROR: %s already exists in the database.\n", item) // 400
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	newprice, err := strconv.ParseFloat(price, 32)

	if err != nil {
		msg := fmt.Sprintf("ERROR: %s is an invalid price.\n", price) // 400
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	(*todb)[item] = dollars(newprice)
	fmt.Fprintf(w, "Created entry %s with price %s.\n", item, (*todb)[item])
}

func (todb *database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	if _, ok := (*todb)[item]; ok {
		delete((*todb), item)
		msg := fmt.Sprintf("%s was deleted from the database.\n", item)
		http.Error(w, msg, http.StatusOK)
		return

	} else {
		msg := fmt.Sprintf("%s does not exist in the database.\n", item)
		http.Error(w, msg, http.StatusNotFound)
	}
}
