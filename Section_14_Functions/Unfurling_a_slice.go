package main

import "fmt"

func main() {
	// xi := []int{2, 3, 4, 5, 6, 7, 8, 9}
	// x := sum(xi...)
	// x := sum()
	x := sum("James")
	fmt.Println("The total is", x)
}

// func sum(x ...int) int {
func sum(s string, x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("For item in index position,", i, "we are now adding,", v, "to the total which is now,", sum)
	}
	return sum
}
