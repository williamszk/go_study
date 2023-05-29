package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password123`
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(hashedPassword)

	// pretend that we have a login password now
	loginPassword := `password123`
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(loginPassword))
	if err != nil {
		fmt.Println(err)
	}
}
