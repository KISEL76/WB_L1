package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var s string
	fmt.Print("Введите строку: ")
	_, _ = fmt.Fscanln(in, &s)

	fmt.Println(isUniq(s))
}

func isUniq(in string) bool {
	m := make(map[rune]struct{}, len(in))
	for _, r := range in {
		if _, ok := m[unicode.ToLower(r)]; ok {
			return false
		}
		m[unicode.ToLower(r)] = struct{}{}
	}
	return true
}
