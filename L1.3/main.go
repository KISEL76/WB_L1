package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Укажите количество воркеров: go run main.go <число_воркеров>")
		return
	}

	numWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Ошибка чтения числа количества воркеров: %v\n", err)
		return
	}

	ch := make(chan int)
	for i := range numWorkers {
		go worker(i, ch)
	}

	counter := 0
	for {
		ch <- counter
		counter++
		time.Sleep(500 * time.Millisecond)
	}
}

func worker(id int, data chan int) {
	for val := range data {
		fmt.Printf("%d worker got %d\n", id, val)
	}
}
