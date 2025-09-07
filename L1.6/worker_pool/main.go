package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Job struct{ id int }

func worker(ctx context.Context, id int, jobs <-chan Job, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			results <- fmt.Sprintf("worker %d: ctx canceled", id)
			return
		case job, ok := <-jobs:
			if !ok {
				results <- fmt.Sprintf("worker %d: jobs closed", id)
				return
			}
			results <- fmt.Sprintf("worker %d: done job %d", id, job.id)
		}
	}
}

func main() {
	const workers = 3
	jobs := make(chan Job)
	results := make(chan string, 16)

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	for i := 1; i <= workers; i++ {
		wg.Add(1)
		go worker(ctx, i, jobs, results, &wg)
	}

	go func() {
		for i := 1; i <= 5; i++ {
			jobs <- Job{id: i}
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		cancel()
		close(results)
	}()

	for r := range results {
		fmt.Println(r)
	}
	_ = time.Now()
}
