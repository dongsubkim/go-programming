package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	// first string
	ltk bool
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}
	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		age:   27,
	}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
	// Embedded structs
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		// first: "something collision",
		ltk: true,
	}
	fmt.Println(sa1)
	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk)
	// Anonymous Structs
	p3 := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   32,
	}
	fmt.Println(p3)
}
