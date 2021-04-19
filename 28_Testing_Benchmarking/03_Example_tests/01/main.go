package main

import (
	"fmt"

	"github.com/dongsubkim/go-programming/28_Testing_Benchmarking/03_Example_tests/01/acdc"
)

func main() {
	fmt.Println(acdc.Sum(2, 3))
	fmt.Println(acdc.Sum(2, 3, 4, 5, 6, 7, 8, 9))
}
