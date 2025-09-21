package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	const workers = 4
	jobs := make(chan int, 128)

	var wg sync.WaitGroup
	for i := 1; i <= workers; i++ {
		wg.Add(1)
		go worker(ctx, i, jobs, &wg)
	}

	go func() {
		defer close(jobs)
		t := time.NewTicker(100 * time.Millisecond)
		defer t.Stop()
		counter := 0
		for {
			select {
			case <-ctx.Done():
				return
			case <-t.C:
				jobs <- counter
				counter++
			}
		}
	}()

	wg.Wait()
	fmt.Fprintln(os.Stdout, "\ngraceful shutdown complete")
}
