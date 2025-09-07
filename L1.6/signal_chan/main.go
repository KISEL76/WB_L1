package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})

	go worker(done)
	time.Sleep(5 * time.Second)
	done <- struct{}{}
	time.Sleep(2 * time.Microsecond)
}

func worker(done chan struct{}) {
	for {
		select {
		case <-done:
			fmt.Println("Выходим по сигналу...")
			return
		default:
			fmt.Println("do something...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
