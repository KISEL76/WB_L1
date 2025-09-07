package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

func (p Point) Distance(other Point) float64 {
	return math.Sqrt((other.x-p.x)*(other.x-p.x) + (other.y-p.y)*(other.y-p.y))
}

func main() {
	p1 := NewPoint(10, -2)
	p2 := NewPoint(-3, 4)

	fmt.Println(p1.Distance(p2))
}
