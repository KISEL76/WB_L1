package main

import (
	"fmt"
)

func main() {
	a := int64(2_000_000)
	b := int64(1_500_000)

	sum := a + b
	diff := a - b
	mul := a * b
	div := a / b

	fmt.Println("a =", a, "b =", b)
	fmt.Println("Сумма =", sum)
	fmt.Println("Разность =", diff)
	fmt.Println("Произведение =", mul)
	fmt.Println("Частное =", div)
}
