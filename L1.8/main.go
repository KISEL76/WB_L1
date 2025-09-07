package main

import (
	"fmt"
)

func setBit(n int64, i int, val int) int64 {
	if val == 1 {
		return n | (1 << i)
	}
	return n &^ (1 << i)
}

func main() {
	var n int64
	var i, val int
	fmt.Print("Введите число, номер бита и значение (0/1): ")
	if _, err := fmt.Scan(&n, &i, &val); err != nil || (val != 0 && val != 1) {
		fmt.Println("Ошибка ввода")
		return
	}

	fmt.Println("Результат:", setBit(n, i, val))
}
