// main.go
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	demo := flag.String("demo", "list", "which demo to run: list|flag|stopchan|jobsclose|ctxcancel|ctxtimeout|goexit|osexit|panicrecover")
	flag.Parse()

	switch *demo {
	case "list":
		fmt.Println("available demos:")
		fmt.Println("  -demo=flag         # выход по условию (atomic.Bool)")
		fmt.Println("  -demo=stopchan     # остановка через отдельный канал уведомления")
		fmt.Println("  -demo=jobsclose    # завершение по закрытию рабочего канала (range)")
		fmt.Println("  -demo=ctxcancel    # context.WithCancel")
		fmt.Println("  -demo=ctxtimeout   # context.WithTimeout / WithDeadline")
		fmt.Println("  -demo=goexit       # runtime.Goexit() завершает ТЕКУЩУЮ горутину")
		fmt.Println("  -demo=osexit       # os.Exit — аварийный выход всего процесса (defers не выполняются!)")
		fmt.Println("  -demo=panicrecover # panic внутри горутины с recover")
	case "flag":
		demoFlag()
	case "stopchan":
		demoStopChan()
	case "jobsclose":
		demoJobsClose()
	case "ctxcancel":
		demoCtxCancel()
	case "ctxtimeout":
		demoCtxTimeout()
	case "goexit":
		demoGoexit()
	case "osexit":
		demoOsExit()
	case "panicrecover":
		demoPanicRecover()
	default:
		fmt.Println("unknown demo:", *demo)
		os.Exit(2)
	}
}

// 1) Выход по условию.
func demoFlag() {
	var stop atomic.Bool
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if stop.Load() {
				fmt.Println("[flag] stop flag set -> return")
				return
			}
			fmt.Println("[flag] working...")
			time.Sleep(200 * time.Millisecond)
		}
	}()

	time.Sleep(800 * time.Millisecond)
	stop.Store(true)
	wg.Wait()
}

// 2) Канал уведомления: закрытие/отправка в stop-канал.
func demoStopChan() {
	stop := make(chan struct{})
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-stop:
				fmt.Println("[stopchan] received stop -> return")
				return
			default:
				fmt.Println("[stopchan] working...")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	time.Sleep(800 * time.Millisecond)
	close(stop)
	wg.Wait()
}

// 3) Завершение по закрытию рабочего канала.
func demoJobsClose() {
	jobs := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for j := range jobs {
			fmt.Println("[jobsclose] got job", j)
		}
		fmt.Println("[jobsclose] jobs closed -> return")
	}()

	for i := 1; i <= 3; i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
}

// 4) Контекст: ручная отмена WithCancel.
func demoCtxCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("[ctxcancel]", ctx.Err(), "-> return")
				return
			default:
				fmt.Println("[ctxcancel] working...")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	time.Sleep(800 * time.Millisecond)
	cancel()
	wg.Wait()
}

// 5) Контекст: авто-остановка по таймауту.
func demoCtxTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 700*time.Millisecond)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("[ctxtimeout]", ctx.Err(), "-> return")
				return
			default:
				fmt.Println("[ctxtimeout] working...")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	wg.Wait()
}

// 6) Принудительное завершение текущей горутины: runtime.Goexit().
func demoGoexit() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		defer fmt.Println("[goexit] deferred cleanup before exit")
		fmt.Println("[goexit] doing work...")
		time.Sleep(200 * time.Millisecond)
		runtime.Goexit()
	}()

	wg.Wait()
	fmt.Println("[goexit] main still alive")
}

// 7) Жёсткая остановка процесса: os.Exit.
func demoOsExit() {
	go func() {
		fmt.Println("[osexit] fatal condition -> exit whole process")
		time.Sleep(200 * time.Millisecond)
		os.Exit(1)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("[osexit] you should not see this")
}

// 8) panic + recover ВНУТРИ горутины.
func demoPanicRecover() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("[panicrecover] recovered from panic:", r, "-> return")
			}
		}()
		fmt.Println("[panicrecover] simulating fatal error...")
		time.Sleep(200 * time.Millisecond)
		panic("something went wrong")
	}()

	wg.Wait()
	fmt.Println("[panicrecover] main still alive")
}
