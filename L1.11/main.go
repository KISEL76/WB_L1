package main

import "fmt"

func main() {
	mas1 := []int{1, 2, 3}
	mas2 := []int{2, 3, 4}

	// Two pointers
	/*
		result := make([]int, 0, len(mas1)+len(mas2))

		sort.Ints(mas1)
		sort.Ints(mas2)

		l, r := 0, 0
		for l < len(mas1) && r < len(mas2) {
			if mas1[l] == mas2[r] {
				result = append(result, mas1[l])
				l++
				r++
			} else if mas1[l] > mas2[r] {
				r++
			} else if mas1[l] < mas2[r] {
				l++
			}
		}
		fmt.Println(result) */

	m := make(map[int]bool)
	for _, val := range mas1 {
		m[val] = true
	}

	intersection := make([]int, 0, len(mas1)+len(mas2))
	for _, val := range mas2 {
		if m[val] {
			intersection = append(intersection, val)
		}
	}

	fmt.Println(intersection)
}
