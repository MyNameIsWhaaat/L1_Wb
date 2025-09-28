package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var m sync.Map

	m.Store("hits", 0)

	const workers = 16
	const itersPerWorker = 10000

	var wg sync.WaitGroup
	wg.Add(workers)

	for w := 0; w < workers; w++ {
		go func(id int) {
			defer wg.Done()
			for i := 0; i < itersPerWorker; i++ {
				for {
					v, _ := m.Load("hits")
					cur := v.(int)
					newV := cur + 1
					if m.CompareAndSwap("hits", cur, newV) {
						break
					}
				}

				if i%2500 == 0 {
					time.Sleep(time.Microsecond)
				}
			}
		}(w)
	}

	wg.Wait()
	v, _ := m.Load("hits")
	fmt.Printf("hits = %d (ожидалось = %d)\n", v.(int), workers*itersPerWorker)
}
