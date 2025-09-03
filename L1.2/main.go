package main

import (
	"fmt"
	"sync"
)

func main() {
	s := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	wg.Add(len(s))
	for i := range s {
		i := i
		go func(idx int) {
			defer wg.Done()
			s[i] = s[i] * s[i]
		}(i)
	}
	wg.Wait()
	fmt.Println(s)
}
