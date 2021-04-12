package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	s := square{
		side: 10.5,
	}
	c := circle{
		radius: 10.1,
	}
	info(s)
	info(c)
}
