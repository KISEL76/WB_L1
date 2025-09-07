package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func worker(stop *atomic.Bool) {
	for {
		if stop.Load() {
			fmt.Println("Останавливаемся по флагу атомика...")
			return
		}
		fmt.Println("do something")
		time.Sleep(30 * time.Millisecond)
	}
}

func main() {
	var stop atomic.Bool
	go worker(&stop)

	time.Sleep(100 * time.Millisecond)
	stop.Store(true)
	time.Sleep(50 * time.Millisecond)
}
