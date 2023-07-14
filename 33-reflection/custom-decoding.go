package main

import "encoding/json"

type response struct {
	Item   string `json:"item"`
	Album  string
	Title  string
	Artist string
}

type respWraper struct {
	response
}

// Custom JSON deconding
func (r *respWraper) UnmarshalJSON(b []byte) (err error) {
	var raw map[string]any

	json.Unmarshal(b, &r.response) // ignore errors
	err = json.Unmarshal(b, &raw)  // ignore errors

	switch r.Item {
	case "album":
		// get album data out of it
		innerDict, ok := raw["album"].(map[string]any)
		if ok {
			if title, ok := innerDict["title"].(string); ok {
				r.Album = title
			}
		}
	case "song":
		innerDict, ok := raw["song"].(map[string]any)
		if ok {
			if title, ok := innerDict["title"].(string); ok {
				r.Title = title
			}
			if artist, ok := innerDict["artist"].(string); ok {
				r.Artist = artist
			}
		}
	}
	return err
}
