package main

import (
	"encoding/json"
	"os"
)

type Person struct {
	Name        string `json:"Name"`
	Surname     string `json:"Surname"`
	Age         int    `json:"Age"`
	Profession  string `json:"Profession"`
	Bankbalance int    `json:"Bankbalance"`
}

func main() {

	faramir := Person{
		Name:        "Faramir",
		Surname:     "Denethorson",
		Age:         34,
		Profession:  "Knight of Gondor",
		Bankbalance: 100000.0,
	}

	// the difference between using the Marshal and the enconder
	// is that with the encoder we can send the information directly
	// to the interface of os.Stdout without the need to save into a variable
	// we can substitu the os.Stdout with a file e.g.
	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(faramir)

}
