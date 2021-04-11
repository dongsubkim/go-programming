package main

import "fmt"

// // var x int // package scope

// func main() {
// 	var x int // func main scope
// 	fmt.Println(x)
// 	x++

// 	{
// 		y := 42 // code block in code block scope
// 		fmt.Println(y)
// 	}
// 	// fmt.Println(y)
// 	fmt.Println(x)
// 	foo()
// }

// func foo() {
// 	fmt.Println("Hello")
// 	// x++
// }

func main() {
	a := incrementor()
	b := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())

}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
