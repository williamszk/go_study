// https://www.youtube.com/watch?v=4vCb7jmwkzM&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=119&ab_channel=AprendaGo
package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "20julho1980"

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(hashedPassword))
	fmt.Println(string([]byte(password)))

	// wrongPass := "21julho1980"

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))

	if err != nil {
		fmt.Println(err)
		// crypto/bcrypt: hashedPassword is not the hash of the given password
	} else {
		fmt.Println("Access granted")
	}
}

// next:
// https://www.youtube.com/watch?v=-wHOZpi697M&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=120&ab_channel=AprendaGo
