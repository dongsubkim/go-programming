package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := foo(xi...)
	s2 := bar(xi)
	fmt.Println(s1)
	fmt.Println(s2)
}

func foo(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func bar(xi []int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}
