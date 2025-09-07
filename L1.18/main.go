package main

import (
	"fmt"
	"sync"
)

const (
	go_count = 1000
)

type Counter struct {
	count int
	mu    sync.Mutex
}

func NewCounter() *Counter {
	return &Counter{
		count: 0,
	}
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {
	var c Counter
	var wg sync.WaitGroup

	wg.Add(go_count)
	for i := 0; i < go_count; i++ {
		go func() {
			defer wg.Done()
			c.Increment()
		}()
	}
	wg.Wait()
	fmt.Println(c.Value())
}
