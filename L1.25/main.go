package main

import (
	"fmt"
	"time"
)

func Sleep(d time.Duration) {
	if d <= 0 {
		return
	}
	timer := time.NewTimer(d)
	<-timer.C
}

func main() {
	start := time.Now()
	fmt.Println("start:", start.Format(time.RFC3339Nano))

	Sleep(1500 * time.Millisecond)

	fmt.Println("done :", time.Now().Format(time.RFC3339Nano))
	fmt.Println("elapsed:", time.Since(start))
}
