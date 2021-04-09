package main

import "fmt"

func main() {
	a := []string{"James", "Bonds", "Shaken, not stirred"}
	b := []string{"Miss", "Moneypenny", "Helloooooo, James"}
	x := [][]string{}
	x = append(x, a, b)
	fmt.Println(x)
	for i, v := range x {
		for j, y := range v {
			fmt.Println(i, v, j, y)
		}
	}
}
