package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("my fisrt func expression")
	}
	f()
	g := func(x int) {
		fmt.Println("the year big brother started watching:", x)
	}
	g(1984)
}
