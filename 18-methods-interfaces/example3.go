package main

import (
	"image/color"
	"math"
)

func (p Point) DistanceTo(q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}
