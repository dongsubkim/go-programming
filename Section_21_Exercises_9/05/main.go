package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64

	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
			fmt.Println(atomic.LoadInt64(&counter))
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}
