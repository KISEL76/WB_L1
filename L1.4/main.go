package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
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
	ctx, cancel := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sig
		fmt.Printf("\nПолучен сигнал завершения, выходим...\n")
		cancel()
		close(ch)
	}()

	for i := range numWorkers {
		go worker(i, ch, ctx)
	}

	counter := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Главная горутина завершена.")
			return
		case ch <- counter:
			counter++
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func worker(id int, data chan int, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d завершает работу\n", id)
			return
		case val, ok := <-data:
			if !ok {
				fmt.Printf("Worker %d: канал закрыт\n", id)
				return
			}
			fmt.Printf("Worker %d получил: %d\n", id, val)
		}
	}
}
