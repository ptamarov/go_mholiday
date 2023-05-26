package main

import "fmt"

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func walkTree(t *Tree) string {

	if t == nil {
		return "()"
	}

	s := ""

	if t.Left != nil {
		s += walkTree(t.Left) + " "
	}

	s += fmt.Sprint(t.Value)

	if t.Right != nil {
		s += " " + walkTree(t.Right)
	}

	return "(" + s + ")"
}
