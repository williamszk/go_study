// https://www.udemy.com/course/end-to-end-go-3-in-1/learn/lecture/11722432#overview
package main

import "fmt"

type Person struct {
	Name    string
	Age     int
	Address string
}

func main() {

	v0 := map[string]int{
		"first":  1,
		"second": 2,
	}

	fmt.Println(v0)

	// allocate memory for a map
	var v map[string]int = make(map[string]int)
	v["first"] = 1
	v["second"] = 2

	// comma ok syntax
	if item, ok := v["third"]; ok {

		fmt.Println(item) // if the item doesn't exist it will show 0
		fmt.Println(ok)   // and the ok will be false
	} else {
		fmt.Println("We couldn't find the value")
	}

	v2 := make(map[int]string)

	v2[0] = "zero"
	v2[1] = "first"
	v2[2] = "second"

	fmt.Println(v2)

	// how to delete an item from the map
	delete(v2, 0)
	fmt.Println(v2)

	// structs
	p1 := Person{
		Name:    "Bob",
		Age:     80,
		Address: "My street",
	}

	fmt.Println(p1)

	// Maps are more like dictionaries in Python
	// and structs are more like classes
	// the JavaScript Object is more like a dictionary
	// JS also has the concept of classes

}
