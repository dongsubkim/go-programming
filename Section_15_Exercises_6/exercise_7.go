package main

import "fmt"

var x int

// var x int = 7
// var g func()
var g func() = func() {
	fmt.Println("g from outside")
}

func main() {
	f := func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}
	}
	f()
	fmt.Printf("%T\n", f)
	fmt.Println(x)

	g()
	g = f
	g()
	fmt.Printf("%T\n", g)

}
