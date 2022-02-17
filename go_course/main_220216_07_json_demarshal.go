package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	myslice := []byte(`{"Name":"Faramir","Surname":"Denethorson","Age":34,"Profession":"Knight of Gondor","Bankbalance":100000}`)

	// the `json:"Name"` are called tags when importing json the tag name will be translated into
	// the field in the struct, and if we export to json the field name will the transformed into
	// the corresponding tag
	type Person struct {
		Name        string `json:"Name"`
		Surname     string `json:"Surname"`
		Age         int    `json:"Age"`
		Profession  string `json:"Profession"`
		Bankbalance int    `json:"Bankbalance"`
	}

	var faramir Person

	err := json.Unmarshal(myslice, &faramir)

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Printf("%v, %T\n", faramir, faramir)
}
