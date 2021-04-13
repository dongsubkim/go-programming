package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	counter := 0

	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			defer wg.Done()

			v := counter
			runtime.Gosched()
			v++
			counter = v
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}
