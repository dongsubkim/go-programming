package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) speak() {
	fmt.Println(p.Name, "is", p.Age, "Years Old")
}

type Human interface {
	speak()
}

func saySomething(h Human) {
	h.speak()
}

func main() {
	p1 := Person{
		Name: "James Bond",
		Age:  32,
	}
	saySomething(&p1)
	// saySomething(p1) // Fail to run

	p1.speak()
}
