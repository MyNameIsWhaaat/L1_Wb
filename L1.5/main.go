package main

import (
	"fmt"
	"time"
)

func main() {
	N := 5
	ch := make(chan int)

	go func() {
		counter := 1
		for {
			ch <- counter
			counter++
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for v := range ch {
			fmt.Println("Получено:", v)
		}
	}()

	<-time.After(time.Duration(N) * time.Second)

	fmt.Println("Завершение работы...")
	close(ch)
}
