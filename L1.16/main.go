package main

import "fmt"

func partition(arr []int, left, right int) int {
	pivot := arr[(left+right)/2]
	for left <= right {
		for arr[left] < pivot {
			left++
		}
		for arr[right] > pivot {
			right--
		}
		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}
	return left
}

func qSort(arr []int, left, right int) {
	if left < right {
		pi := partition(arr, left, right)
		qSort(arr, left, pi-1)
		qSort(arr, pi, right)
	}
}

func main() {
	arr := []int{33, 10, 55, 71, 29, 3, 18, 42, 5}
	fmt.Println("До сортировки:", arr)

	qSort(arr, 0, len(arr)-1)

	fmt.Println("После сортировки:", arr)
}
