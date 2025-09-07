package main

import (
	"fmt"
	"math"
)

type SquareAreaCalculator interface {
	CalcSquareArea(side float64) float64
}

type CircleAreaCalculator struct{}

func (c *CircleAreaCalculator) CalcCircleArea(radius float64) float64 {
	return math.Pi * radius * radius
}

type CircleToSquareAdapter struct {
	circle *CircleAreaCalculator
}

func NewCircleToSquareAdapter(c *CircleAreaCalculator) *CircleToSquareAdapter {
	return &CircleToSquareAdapter{circle: c}
}

func (a *CircleToSquareAdapter) CalcSquareArea(side float64) float64 {
	radius := side / 2
	return a.circle.CalcCircleArea(radius)
}

func printArea(calculator SquareAreaCalculator, side float64) {
	fmt.Printf("Площадь через интерфейс (side=%.2f) = %.2f\n", side, calculator.CalcSquareArea(side))
}

func main() {
	circleCalc := &CircleAreaCalculator{}
	adapter := NewCircleToSquareAdapter(circleCalc)

	printArea(adapter, 10)
	printArea(adapter, 20)
}
