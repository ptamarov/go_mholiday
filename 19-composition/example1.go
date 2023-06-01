package main

import "fmt"

type Pair struct {
	Path string
	Hash string
}

type PairWithLength struct {
	Pair
	Length int
}

func (p Pair) String() string {
	return fmt.Sprintf("Has of %s is %s", p.Path, p.Hash)
}

func (p PairWithLength) String() string {
	return fmt.Sprintf("Has of %s is %s; length %d", p.Path, p.Hash, p.Length)
}
