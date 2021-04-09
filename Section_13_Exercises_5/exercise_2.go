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

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {
		// fmt.Println(k)
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, v2 := range v.favIcecream {
			fmt.Println(i, v2)
		}
		fmt.Println("---------------")
	}
}
