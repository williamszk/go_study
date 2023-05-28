package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	rawString := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":28}]`
	sliceBytes := []byte(rawString)

	var people []Person
	err := json.Unmarshal(sliceBytes, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(people)

}
