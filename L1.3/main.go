package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

func main() {
	
	def := runtime.NumCPU()
	n := flag.Int("workers", def, "number of workers")
	flag.Parse()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	jobs := make(chan int, 128)

	var wg sync.WaitGroup
	for i := 1; i <= *n; i++ {
		wg.Add(1)
		go worker(ctx, i, jobs, &wg)
	}

	go func() {
		defer close(jobs)
		counter := 0
		t := time.NewTicker(100 * time.Millisecond)
		defer t.Stop()

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
	fmt.Fprintln(os.Stdout, "done")
}