// https://medium.com/asecuritysite-when-bob-met-alice/golang-and-cryptography-914db9d7069f

package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	s := "hello"
	argCount := len(os.Args[1:])
	if argCount > 0 {
		s = os.Args[1]
	}
	hash := md5.Sum([]byte(s))
	hash2 := sha256.Sum256([]byte(s))
	//h1.Write([]byte(s))
	fmt.Printf("Input:\t\t\t%s\n", s)
	fmt.Printf("Input (hex):\t\t%x\n", []byte(s))
	fmt.Printf("Input (hex):\t\t%s\n", hex.EncodeToString([]byte(s)))
	fmt.Printf("Input (Base64):\t\t%s\n", base64.StdEncoding.EncodeToString([]byte(s))) // base64 is another package, we use it to transform a slice of bytes into a string of base64 numbers
	fmt.Printf("MD5 (Hex):\t\t%x\n", hash)
	fmt.Printf("MD5 (hex):\t\t%s\n", hex.EncodeToString(hash[:])) // here we use the hex package to transform a slice of bytes into a string of hexadecimal numbers
	fmt.Printf("MD5 (Base64):\t\t%s\n", base64.StdEncoding.EncodeToString(hash[:]))
	fmt.Printf("SHA-256 (Hex):\t\t%x\n", hash2)
	fmt.Printf("SHA-256 (hex):\t\t%s\n", hex.EncodeToString(hash2[:]))
	fmt.Printf("SHA-256 (Base64):\t%s\n", base64.StdEncoding.EncodeToString(hash2[:]))

	fmt.Printf("The type of hash : %T\n", hash)       // hash is an array
	fmt.Printf("The type of hash[:] : %T\n", hash[:]) // when we use [:] we convert the array into a slice
}

// to run the program
// go build .
// .\base64-converstion.exe
// we can pass an argument to the program
// .\base64-converstion.exe goisawesomewithcryptography
