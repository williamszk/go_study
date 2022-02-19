// https://www.youtube.com/watch?v=4vCb7jmwkzM&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=120&ab_channel=AprendaGo
package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	password := "20julho2980"

	// how to convert string to slice of bytes in go?
	seqBytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(seqBytes))

	fmt.Println(bcrypt.CompareHashAndPassword(seqBytes, []byte("20julho2980")))
	fmt.Println(bcrypt.CompareHashAndPassword(seqBytes, []byte("20julho2981")))

}

// go mod init example.com/greetings
// go get -u "golang.org/x/crypto/bcrypt"
// go get is like pip install

// next
// https://www.youtube.com/watch?v=-wHOZpi697M&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=121&ab_channel=AprendaGo
