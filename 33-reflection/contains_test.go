package main

import (
	"encoding/json"
	"testing"
)

var unknown = `{
    "id": 1, 
    "name": "bob", 
    "addr": {
            "street": "Lazy Lane",
            "city": "Exit",
            "zip" : "99999"
    },
    "extra": 21.1
}`

func CheckData(want, got []byte) error {
	var w, g Data

	if err := json.Unmarshal(want, &w); err != nil {
		return err
	}

	if err := json.Unmarshal(got, &g); err != nil {
		return err
	}

	return Contains(w, g)
}

func TestContains(t *testing.T) {
	var known = []string{
		`{"id": 1}`,
		`{"extra": 21.1 }`,
		`{"name": "bob"}`,
		`{"names": "bob"}`,
		`{"addr": {"street": "Lazy Lane", "city": "Exit"}}`,
	}

	for _, k := range known {
		if err := CheckData([]byte(k), []byte(unknown)); err != nil {
			t.Errorf("invalid: %s (%s)\n", k, err)
		}
	}
}
