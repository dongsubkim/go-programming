// THIS IS NOT HOW TO CODE. CHECK Mutex.go TO SEE HOW TO FIX THE RACE CONDITION.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("GOROUTINES:", runtime.NumGoroutine())

	counter := 0

	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			// time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("GOROUTINES:", runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println("GOROUTINES:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
