package main

import "fmt"

func main() {
	defer foo()
	bar()
}

func foo() {
	defer func() {
		fmt.Println("FOO INSIDE")
	}()
	fmt.Println("FOO")
}

func bar() {
	fmt.Println("BAR")
}
