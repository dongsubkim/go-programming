package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Hello World")
	wg.Add(1)
	go func() {
		fmt.Println("gorountine 2")
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		fmt.Println("goroutine 3")
		wg.Done()
	}()

	fmt.Println("goroutine 1")
	wg.Wait()
}
