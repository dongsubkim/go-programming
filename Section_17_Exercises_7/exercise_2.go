package main

import "fmt"

type person struct {
	first string
	last  string
}

func changeMe(p *person) {
	(*p).first = "John"
	// p.first = "John"
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)

}
