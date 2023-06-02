package main

import (
	"bytes"
	"fmt"
	"io"
)

// Interface has a pointer that represents the type, and a pointer to the concrete data.

var r io.Reader     // is nil
var b *bytes.Buffer // is also nil

// Interface is nil only if the pointer to the type is nil and the pointer to the data is nil.
func main() {
	if r == nil {
		fmt.Println("r is nil")
	}

	r = b // r is no longer nil, but it points to nil data.

	if r == nil {
		fmt.Println("r is nil")
	} else {
		fmt.Println("r is no longer nil.")
	}

	if b == nil {
		fmt.Println("...but b is still nil.")
	}

	store := []byte{45, 12, 23, 36}
	b = bytes.NewBuffer(store)

	if b == nil {
		fmt.Println("...but b is still nil.")
	} else {
		fmt.Println("... and now b is not nil.")
	}
}
