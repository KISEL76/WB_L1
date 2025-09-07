package main

import "fmt"

func main() {
	sentence := "snow dog sun"
	res := ReverseSentence(sentence)
	fmt.Println(res)
}

func ReverseRunes(runes []rune, left, right int) {
	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}
}

func ReverseSentence(sentence string) string {
	sen := []rune(sentence)
	ReverseRunes(sen, 0, len(sen)-1)

	start := 0
	for i := 0; i <= len(sen); i++ {
		if i == len(sen) || sen[i] == ' ' {
			ReverseRunes(sen, start, i-1)
			start = i + 1
		}
	}
	return string(sen)
}
