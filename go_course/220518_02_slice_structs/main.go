// We want to test instead of building a type that is a slice of people
// we first define the type that is a struct
// and then we create slice based on this type
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	message := `[{"FirstName":"James","Age":32},{"FirstName":"Moneypenny","Age":27},{"FirstName":"M","Age":54}]`

	messageBytes := []byte(message)

	var SliceOfPerson []Person

	err := json.Unmarshal(messageBytes, &SliceOfPerson)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("SliceOfPerson = %v, %T\n", SliceOfPerson, SliceOfPerson)
	// SliceOfPerson = [{James 32} {Moneypenny 27} {M 54}], []main.Person

	fmt.Printf("SliceOfPerson[0] = %v, %T\n", SliceOfPerson[0], SliceOfPerson[0])
	// SliceOfPerson[0] = {James 32}, main.Person

	//
}

type Person struct {
	FirstName string `json:"FirstName"`
	Age       int    `json:"Age"`
}

// this is like the interface in typescript
// this is an alternative instead of building a type that is
// a slice of people
