package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}

	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   28,
	}

	pSlice := []person{p1, p2}

	// sliceByte is the slice of bytes that contain the json corresponding
	// to the value that we passed to the Marshall function
	sliceByte, err := json.Marshal(pSlice)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(sliceByte))
}

type person struct {
	First string
	Last  string
	Age   int
}
