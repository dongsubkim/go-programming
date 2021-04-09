package main

import "fmt"

func main() {
	// composite literal
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])
	// for range
	for i, v := range x {
		fmt.Println(i, v)
	}
	// slicing a slice
	fmt.Println(x[1:])
	fmt.Println(x[1:3])

	for i := 0; i <= 4; i++ {
		fmt.Println(i, x[i])
	}
	// append to a slice
	x = append(x, 77, 88, 99, 1014)
	fmt.Println(x)

	y := []int{234, 456, 678, 987}
	x = append(x, y...)
	fmt.Println(x)

	// deleting from a slice
	x = append(x[:2], x[4:]...)
	fmt.Println(x)

	// make
	x = make([]int, 10, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x[0] = 42
	x[9] = 999
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x = append(x, 1234)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x = append(x, 12345)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x = append(x, 123456)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// multi-dimensional slice
	jb := []string{"James", "Bond", "Chocolate", "martini"}
	fmt.Println(jb)
	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}
	fmt.Println(mp)

	xp := [][]string{jb, mp}
	fmt.Println(xp)
}
