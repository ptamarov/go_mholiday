package main

import (
	"fmt"
	"math"
	"strings"
)

type Point struct {
	X, Y float64
}

type Line struct {
	Start, End Point
}

type Path []Point

func (l Line) Length() float64 {
	return math.Hypot(l.Start.X-l.End.X, l.Start.Y-l.End.Y)
}

func (p Path) Length() (sum float64) {
	for i := 1; i < len(p); i++ {
		sum += Line{p[i-1], p[i]}.Length()
	}
	return
}

type Lengther interface {
	Length() float64
}

func PrintLength(l ...Lengther) {

	s := []string{}

	for _, elt := range l {
		s = append(s, fmt.Sprint(elt.Length()))
	}

	fmt.Println(strings.Join(s, ", "))
}
