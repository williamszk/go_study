package main

import "fmt"

func main() {
	// create an anonymous struct
	person := struct {
		first string
		last  string
	}{
		first: "James",
		last:  "Bond",
	}
	fmt.Println(person)
	// with anonymous structs we need to declare the types and the values
	// the keyword "type" will simply create an alias actually
	// in all cases, even with structs, the word that comes right after the
	// keyword "type" will be the alias to something, this something will come
	// after the alias.
	// `type <alias(1)> <something(2)>`
	// the `<something(2)>` usually is something that is too big to and cumbersome
	// to write every time that we need to use it.
	// Like: `struct{first string; last string}` or `map[string]int`
	// so the `<alias(1)>` is just that an alias, and we can simply substitute the
	// alias with the the literal `<something(2)>`

	// we can create anonymous composite types of:
	// struct, map, slices

}

// composite type
// composite literal
// "composite types" refer to structs, slices, maps, which are types that are made
// with primitive types, like string, int, bool
