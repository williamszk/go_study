// https://www.youtube.com/watch?v=oUsTxBwHaMM&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=122&ab_channel=AprendaGo

package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	message := `[{"FirstName":"James","Age":32},{"FirstName":"Moneypenny","Age":27},{"FirstName":"M","Age":54}]`
	// we need to use json to go to find in a easy way the interface of this json
	// https://mholt.github.io/json-to-go/

	fmt.Println(message)

	// the unmarshall function needs to receive a slice of bytes
	messageBytes := []byte(message)

	var SlicePerson People
	// the second argument of the Unmarshall function is an object
	// where we can store the bytes into native Go objects
	err := json.Unmarshal(messageBytes, &SlicePerson)
	// bie that we need to pass the address of the object
	// this is called dereference

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("SlicePerson", SlicePerson)

	fmt.Printf("SlicePerson = %v, %T\n", SlicePerson, SlicePerson)

	fmt.Printf("SlicePerson[0] = %v, %T\n", SlicePerson[0], SlicePerson[0])
	// the type returned as:
	// struct { FirstName string "json:\"FirstName\""; Age int "json:\"Age\"" }
}

// this is a slice of structs
// not a struct itself
type People []struct {
	FirstName string `json:"FirstName"`
	Age       int    `json:"Age"`
}

type Person struct {
	FirstName string `json:"FirstName"`
	Age       int    `json:"Age"`
}
