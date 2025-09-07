package main

import (
	"fmt"
	"time"
)

func worker(stopAt int) {
	for i := 0; i < 1000; i++ {
		if i == stopAt {
			fmt.Println("Остановка по условию, выход...")
			return
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	go worker(20)
	time.Sleep(500 * time.Millisecond)
}
