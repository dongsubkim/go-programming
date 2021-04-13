package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password1234`

	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(bs)

	loginPassword := `password123`
	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPassword))
	if err != nil {
		fmt.Println("YOU CAN'T LOGIN")
		return
	}
	fmt.Println("You're logged in")

}
