package main

import (
	"fmt"
	"image/color"
)

func main() {
	// Example 1.
	side := Line{Point{1, 2}, Point{3, 4}}
	path := Path{{1, 1}, {5, 1}, {5, 4}, {1, 1}}
	PrintLength(side, path)

	// Example 2.
	f := 2.00
	side.ScaleBy(f)
	fmt.Println(side)

	// Example 3.
	pUsual := Point{1, 2}
	pColor := ColoredPoint{Point{5, 4}, color.RGBA{255, 0, 0, 255}}
	distance := pColor.DistanceTo(pUsual) // Works, even though pColor is not of type Point.
	fmt.Println(distance)
}
