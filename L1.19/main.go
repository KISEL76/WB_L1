package main

import "fmt"

func Reverse(in string) {
	s := []rune(in)
	l, r := 0, len(s)-1

	for l <= r {
		s[r], s[l] = s[l], s[r]
		l++
		r--
	}
	fmt.Println(string(s))
}

func main() {
	in := "главрыба"
	Reverse(in)
}
