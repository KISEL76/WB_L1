package main

import (
	"fmt"
	"runtime"
)

func worker() {
	defer fmt.Println("worker: defer still runs before Goexit returns")
	fmt.Println("worker: about to Goexit()")
	runtime.Goexit()
	fmt.Println("Этого никто не увидит")
}

func main() {
	go worker()
	fmt.Scanln()
}
