package main

import "fmt"

func main() {
	nums := []int{10, 20, 30, 40, 50}
	fmt.Println("До:", nums)

	nums = DeleteItem(nums, 2)
	fmt.Println("После:", nums)
}

func DeleteItem(s []int, idx int) []int {
	if idx < 0 || idx > len(s) {
		panic("index out of range")
	}
	copy(s[idx:], s[idx+1:])
	s[len(s)-1] = 0

	return s[:len(s)-1]
}
