package main

import "fmt"

func main() {
	mas := []int{3, 6, 12, 15, 20, 43, 51, 65, 81, 100}
	fmt.Println(binarySearch(mas, 100))
}

func binarySearch(mas []int, val int) int {
	l, r := 0, len(mas)-1
	for l <= r {
		mid := l + (r-l)/2
		if mas[mid] == val {
			return mid
		} else if mas[mid] > val {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}
