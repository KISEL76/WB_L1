package main

import "fmt"

func main() {
	mas := []int{1, 3, 2, 6, 1, 7}
	in := make(chan int)
	out := make(chan int)

	go func() {
		for _, num := range mas {
			in <- num
		}
		close(in)
	}()

	go func() {
		for val := range in {
			out <- val * 2
		}
		close(out)
	}()

	for y := range out {
		fmt.Println(y)
	}
}
