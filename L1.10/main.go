package main

import (
	"fmt"
	"math"
)

func main() {
	mas := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	m := make(map[int][]float64)

	for _, temp := range mas {
		var num int
		if temp < 0 {
			num = int(math.Ceil(temp/10.0) * 10)
		} else {
			num = int(math.Floor(temp/10.0) * 10)
		}
		m[num] = append(m[num], temp)
	}

	for k, v := range m {
		fmt.Printf("%d: %v\n", k, v)
	}
}
