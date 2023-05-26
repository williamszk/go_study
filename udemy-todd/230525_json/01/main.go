package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		age:   28,
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
	first string
	last  string
	age   int
}
