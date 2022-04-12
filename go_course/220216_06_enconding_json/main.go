package main

import (
	"encoding/json"
	"fmt"
)

// For certain object or fields to be visible they need to start with a
// Upper case, lower case variables or symbols are treated as private
type Person struct {
	Name        string
	Surname     string
	Age         int
	Profession  string
	Bankbalance float64
}

func main() {

	faramir := Person{
		Name:        "Faramir",
		Surname:     "Denethorson",
		Age:         34,
		Profession:  "Knight of Gondor",
		Bankbalance: 100000.0,
	}

	// fmt.Println(faramir)

	saruman := Person{
		"Saruman",
		"The White",
		2000,
		"Astarii",
		800000.0,
	}

	// fmt.Println(saruman)

	faramirJson, err := json.Marshal(faramir)
	if err != nil {
		fmt.Println(err)
	}
	sarumanJson, err := json.Marshal(saruman)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(faramirJson))
	fmt.Println(string(sarumanJson))

}
