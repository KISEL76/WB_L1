package main

import (
	"fmt"
	"strings"
)

func createHugeString(size int) string {
	return strings.Repeat("A", size)
}

var justString string

func someFunc() {
	v := createHugeString(1 << 10)

	n := 100
	if len(v) < n {
		n = len(v)
	}

	justString = string([]byte(v[:n]))
}

func main() {
	someFunc()
	fmt.Println("justString =", justString)
	fmt.Println("len(justString) =", len(justString))
}
