package main

import "fmt"

type customErr struct {
	name string
}

func (err customErr) Error() string {
	return "ERROR"
}

func foo(e error) {
	fmt.Println("in foo", e)
}

func main() {
	e1 := customErr{"James Bond"}
	foo(e1)
}
