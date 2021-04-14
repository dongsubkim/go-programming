package main

import "fmt"

func main() {
	c := make(chan<- int, 2)

	c <- 42
	c <- 43

	// BELOW CODE DOES NOT WORK
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	fmt.Println("------")
	fmt.Printf("%T\n", c)
}
