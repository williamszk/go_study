// https://www.coderedcorp.com/blog/golang-for-python-devs/

package main

import "fmt"

func main() {
	fruits := []string{"Apple", "Banana", "Cherry"}
	// int myArr[] = {1,2,3}; // in C

	// Print each fruit with its index.
	for index, fruit := range fruits {
		fmt.Printf("Fruit %d is %s\n", index, fruit)
	}

	// Print each fruit. Here we are throwing away the index by assigning it to _
	// which is a special variable in Go used for ignoring the output.
	for _, fruit := range fruits {
		fmt.Println(fruit)
	}

	// Slices in Golang are equivalent to Lists in Python. Golang also
	// has more traditional arrays, like you’d see in C, but when
	// starting out just focus on slices.

	// Define a slice.
	// The trailing comma after "Cherry" is required or it won't compile!

	// Zero-indexed random access.
	apple := fruits[0]
	cherry := fruits[2]
	fmt.Println(apple, cherry)

	// Golang cannot do negative indexes. So you'll have to count backwards instead.
	cherry = fruits[len(fruits)-1]
	fmt.Println(cherry)

	// As the name indicates, you can take slices of a slice.
	apple_banana := fruits[0:2]
	fmt.Println(apple_banana)

	apple_banana = fruits[:2]
	fmt.Println(apple_banana) // we can print slices in Go

	// Append to the slice. The append function in Go actually returns a new slice.
	fruits = append(fruits, "Durian")
	// maybe it is expensive to append to a slice given that everything needs to be copied again
	fmt.Println(fruits)

	// Combine two lists.
	// Here we are using the special variadic syntax `...` which has a similar
	// behavior to the commonly used `*args` in Python.

	veggies := []string{"Tomato", "Lettuce"}
	foods := append(fruits, veggies...)
	fmt.Println(foods)

	// This is equivalent to fully exploding the veggies slice and appending each
	// element separately.
	foods = append(fruits, "Tomato", "Lettuce")
	fmt.Println(foods)

	// First we must allocate a new map with using ``make()``,
	// declaring key & value types. The syntax here is:
	//
	//     map[key]value
	//
	fruits2 := make(map[string]string)
	fmt.Println(fruits2)
	// As a shortcut, you could instead allocate it with empty contents.
	fruits2 = map[string]string{}
	fmt.Println(fruits2)

	// Or you can allocate one with pre-defined contents.
	fruits2 = map[string]string{
		"a": "Apple",
		"b": "Banana",
		"c": "Cherry",
	}

	fmt.Println(fruits2)

	// Access
	word := fruits2["a"]
	fmt.Println(word) // Apple

	// Assignment
	fruits2["d"] = "Dragonfruit"

	// Looping/iterating over a map.
	for key, value := range fruits2 {
		fmt.Printf("%s: %s\n", key, value)
	}

	// Updating or merging two maps together.
	// There is no direct equivalent to Python's `update()` in Go,
	// but you can accomplish it with a simple loop.
	updateFruits := map[string]string{
		"c": "Cantaloupe",
		"d": "Durian",
	}

	for key, value := range updateFruits {
		fruits2[key] = value
	}
	fmt.Println("After update:", fruits2)
	// note that "Cantaloupe" overwrote "Cherry"

}
