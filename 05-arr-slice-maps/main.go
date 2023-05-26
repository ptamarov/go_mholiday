package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	words := make(map[string]int)

	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		words[scan.Text()]++
	}

	fmt.Println(len(words), "unique words.")

	type keyValue struct {
		key string
		val int
	}

	var ss []keyValue

	for key, value := range words {
		ss = append(ss, keyValue{key, value})
	}

	sort.Slice(ss, func(i int, j int) bool {
		return ss[i].val > ss[j].val
	}) // how to sort the slice?

	fmt.Println("COUNT\t|\tWORD")
	fmt.Println("----------------------------")
	for _, s := range ss {
		fmt.Println(s.val, "\t|\t", s.key)
	}
}
