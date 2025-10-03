package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}

	in := make(chan int)
	out := make(chan int)

	go func() {
		for _, x := range nums {
			in <- x
		}
		close(in)
	}()

	go func() {
		defer close(out)
		for x := range in {
			out <- x * 2
		}
	}()

	for v := range out {
		fmt.Println(v)
	}

}
