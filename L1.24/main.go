package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

func (p Point) Distance(other Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Hypot(dx, dy)
}

func main() {
	a := NewPoint(0, 0)
	b := NewPoint(3, 4)

	fmt.Printf("Расстояние между A и B: %.2f\n", a.Distance(b))
}