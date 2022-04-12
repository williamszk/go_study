// https://www.youtube.com/watch?v=EaOGcmXo4F8&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=80&ab_channel=AprendaGo
// https://www.youtube.com/watch?v=KBFprVi_haM&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=80&ab_channel=AprendaGo
// https://www.youtube.com/watch?v=dKaElWlGKo0&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=81&ab_channel=AprendaGo
// https://www.youtube.com/watch?v=Y4MKS3gJQ9Q&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=82&ab_channel=AprendaGo

package main

import "fmt"

type client struct {
	name    string
	surname string
	smoker  bool
}

type person struct {
	name string
	age  int
}

type professional struct {
	person
	title string
	wage  float32
}

func main() {
	client01 := client{
		name:    "Bob",
		surname: "Marley",
		smoker:  true,
	}

	client02 := client{
		"Joan",
		"Robinson",
		true,
	}

	fmt.Println(client01)
	fmt.Println(client02)

	person01 := person{
		name: "Alfred",
		age:  80,
	}

	person02 := professional{
		person: person{
			name: "Bruce",
			age:  35,
		},
		title: "Batman",
		wage:  99999999.0,
	}

	fmt.Println(person01)
	fmt.Println(person02)

	// how to access field values from structs
	fmt.Println("Alfred's age: ", person01.age)
	fmt.Println("Bruce's wage: ", person02.wage)
	fmt.Println("Bruce's age: ", person02.person.age)

	person03 := professional{
		person{
			"Joker",
			60,
		},
		"The Joker of Gothan",
		-1000000.0,
	}

	fmt.Println(person03)
	fmt.Println("Access Joker's title: ", person03.title)

	// Anonimous Struct
	// it is like a dictionary in Python or an Object in JavaScript?
	// The map in Go is more like dict and object actuatlly
	// actually in a map things are more restricted...
	// we need to have uniformity in the datatypes
	// The usual Struct in Go acts like a Class, but only with attributes

	anonStruct := struct {
		name string
		age  int
	}{
		name: "Obi-Wan",
		age:  40,
	}

	fmt.Println(anonStruct)
}
