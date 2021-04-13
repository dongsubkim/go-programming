package main

import "fmt"

type vehicle struct {
	doors int
	color string
}
type truck struct {
	vehicle
	fourWheel bool
}
type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "white",
		},
		fourWheel: false,
	}
	s1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		luxury: true,
	}
	fmt.Println(t1)
	fmt.Println(t1.doors)
	fmt.Println(t1.color)
	fmt.Println(t1.fourWheel)
	fmt.Println(s1)
	fmt.Println(s1.doors)
	fmt.Println(s1.color)
	fmt.Println(s1.luxury)
}
