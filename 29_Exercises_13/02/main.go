package main

import (
	"fmt"

	"github.com/dongsubkim/go-programming/29_Exercises_13/02/quote"
	"github.com/dongsubkim/go-programming/29_Exercises_13/02/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))
	fmt.Println(word.CountTwo(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
