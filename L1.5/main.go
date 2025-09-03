package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Укажите количество секунд: go run main.go <число>")
		return
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Ошибка чтения числа количества секунд: %v", err)
		return
	}
	ch := make(chan int)

	go func() {
		i := 0
		for {
			ch <- i
			i++
			time.Sleep(time.Millisecond * 500)
		}
	}()

	loopWithTimeout(n, ch)
}

func loopWithTimeout(seconds int, ch chan int) {
	timeout := time.After(time.Duration(seconds) * time.Second)
	for {
		select {
		case val := <-ch:
			fmt.Println(val)
		case <-timeout:
			fmt.Printf("Прошло 10 секунд, выходим...\n")
			return
		}
	}
}
