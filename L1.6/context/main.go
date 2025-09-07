package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	n := 5
	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx)
	time.Sleep(time.Duration(n) * time.Second)
	cancel()
	time.Sleep(50 * time.Millisecond)
}

func worker(ctx context.Context) {
	for {
		select {
		default:
			time.Sleep(500 * time.Millisecond)
			fmt.Println("do something...")
		case <-ctx.Done():
			fmt.Println("Выходим по контексту...")
			return
		}
	}
}
