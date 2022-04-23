// https://www.youtube.com/watch?v=mcbj-wy8Ro8&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=115&ab_channel=AprendaGo

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name        string  `json:"Name"`
	Lastname    string  `json:"Lastname"`
	Age         int     `json:"Age"`
	Profession  string  `json:"Profession"`
	BankAccount float64 `json:"BankAccount"`
}

// the string literal that comes after the name of the field is called
// tags, we can change them

func main() {

	// slice of bytes
	sliceBytes := []byte(`{"Name":"Obi-Wan","Lastname":"Kenobi","Age":65,"Profession":"Jedi","BankAccount":1000000.0}`)
	fmt.Println(sliceBytes)

	// note that we declare a variable that is a struct
	var APerson Person

	// Unmarshal asks to input an address to a struct where we want to include the byte string
	// (which contains the json) into struct
	err := json.Unmarshal(sliceBytes, &APerson)

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Printf("%+v\n", APerson)

	// what %+v does?
	// fmt.Printf("%v %+v %#v\n", person, person, person)
	// // Result: {Kim 22} {Name:Kim Age:22} struct { Name string; Age int }{Name:"Kim", Age:22}

	// We can use the function Encode to send the information to some place
	// in this example we'll send it to the terminal console

	MyEncoder := json.NewEncoder(os.Stdout)
	fmt.Println("An example of using Encode to print a struct:")
	MyEncoder.Encode(APerson)
	fmt.Println("An example of using Encode to print a blob:")
	MyEncoder.Encode(sliceBytes)
	MyEncoder.Encode(string(sliceBytes))

	// [123 34 78 97 109 101 34 58 34 79 98 105 45 87 97 110 34 44 34 76 97 115 116 110 97 109 101 34 58 34 75 101 110 111 98 105 34 44 34 65 103 101 34 58 54 53 44 34 80 114 111 102 101 115 115 105 111 110 34 58 34 74 101 100 105 34 44 34 66 97 110 107 65 99 99 111 117 110 116 34 58 49 48 48 48 48 48 48 46 48 125]
	// {Name:Obi-Wan Lastname:Kenobi Age:65 Profession:Jedi BankAccount:1e+06}
	// An example of using Encode to print a struct:
	// {"Name":"Obi-Wan","Lastname":"Kenobi","Age":65,"Profession":"Jedi","BankAccount":1000000}
	// An example of using Encode to print a blob:
	// "eyJOYW1lIjoiT2JpLVdhbiIsIkxhc3RuYW1lIjoiS2Vub2JpIiwiQWdlIjo2NSwiUHJvZmVzc2lvbiI6IkplZGkiLCJCYW5rQWNjb3VudCI6MTAwMDAwMC4wfQ=="
	// "{\"Name\":\"Obi-Wan\",\"Lastname\":\"Kenobi\",\"Age\":65,\"Profession\":\"Jedi\",\"BankAccount\":1000000.0}"

}

// next:
// https://www.youtube.com/watch?v=S4hEdA0RPVI&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=116&ab_channel=AprendaGo
