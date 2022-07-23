// https://www.udemy.com/course/learn-how-to-code/learn/lecture/4037302#overview
// https://www.udemy.com/course/learn-how-to-code/learn/lecture/4077564#overview
package main

import (
	"fmt"
)

func main() {
	// the slice is a reference type
	// that is it is a pointer that points to an array
	// the slice stores the length and capacity of the array

	// maps are also reference types
	//

	// when dealing with reference types we don't take the address
	// of a slice or map, because they by themselves are pointers

	// what is the difference between:
	// var myGreeting map[string]string
	// var myGreeting = make(map[string]string)
	// myGreeting := make(map[string]string)
	// myGreeting := map[string]string{}

	m := make(map[string]int)

	// how to include values in a map
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// how to delete a key from a map
	delete(m, "k2")
	fmt.Println("after delete k2 map:", m)

	// comma ok idiom
	// get a key that does not exist
	v, ok := m["k2"]
	if ok {
		fmt.Println("v:", v)
	} else {
		fmt.Println("k2 key is not present")
	}

	// instead of using make to instantiate the map object
	// we can use a composite literal to declare an include
	// values in the map

	n := map[string]int{"foo": 1, "bar": 2}

	fmt.Println("map n:", n)
	// maps are not ordered

	// maps are reference types

	var myGreeting map[string]string
	// this will create a nil reference
	// if it is a nil reference we can't use append
	fmt.Println(myGreeting) // map[]

	// myGreeting["Tim"] = "Good Morning"
	// panic: assignment to entry in nil map
	//
	// goroutine 1 [running]:
	// main.main()
	//         /root/go_study/go_course/220717_01_maps/main.go:67 +0x3d3
	//
	// we'll get this runtime error

	fmt.Println("is the map a nil reference? ", myGreeting == nil) // is the map a nil reference?  true

	// we can declare a map that is not a nil reference using {}
	var myGreeting2 = map[string]string{}
	myGreeting2["Tim"] = "Good Morning"
	fmt.Println(myGreeting2) // map[Tim:Good Morning]
	// let's see if the myGreeting2 is a nil referece
	fmt.Println("is the map a nil reference? ", myGreeting2 == nil) // is the map a nil reference? false

	// var myGreeting2 = map[string]string{}
	// and
	// m := make(map[string]int)
	// have the same result

	// underneath a map there is a hashtable
	// underneath a slice there is an array
}
