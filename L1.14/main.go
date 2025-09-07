package main

import (
	"fmt"
)

func detectType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("Тип: int")
	case string:
		fmt.Println("Тип: string")
	case bool:
		fmt.Println("Тип: bool")
	case chan int:
		fmt.Println("Тип: chan int")
	case chan string:
		fmt.Println("Тип: chan string")
	default:
		fmt.Println("Неизвестный тип")
	}
}

func main() {
	detectType(42)
	detectType("hello")
	detectType(true)
	detectType(make(chan int))
	detectType(make(chan string))
	detectType(3.14)
}
