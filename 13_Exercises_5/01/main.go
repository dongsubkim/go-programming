package main

import "fmt"

type person struct {
	first       string
	last        string
	favIcecream []string
}

func main() {
	p1 := person{
		first:       "James",
		last:        "Bond",
		favIcecream: []string{"Chocolate", "mint"},
	}
	p2 := person{
		first:       "Miss",
		last:        "Moneypenny",
		favIcecream: []string{"Strawberry", "vanilla"},
	}
	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.favIcecream {
		fmt.Println(i, v)
	}
	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, v := range p2.favIcecream {
		fmt.Println(i, v)
	}

}
