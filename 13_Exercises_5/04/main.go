package main

import "fmt"

func main() {
	p1 := struct {
		first   string
		flavors []string
	}{
		first:   "James",
		flavors: []string{"mint", "chocolate"},
	}
	fmt.Println(p1)
	s := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "James",
		friends: map[string]int{
			"Moneypenny": 555,
			"Q":          777,
			"M":          888,
		},
		favDrinks: []string{
			"Martini",
			"Water",
		},
	}
	fmt.Println(s.first)
	fmt.Println(s.friends)
	fmt.Println(s.favDrinks)

	for k, v := range s.friends {
		fmt.Println(k, v)
	}

	for i, val := range s.favDrinks {
		fmt.Println(i, val)
	}

}
