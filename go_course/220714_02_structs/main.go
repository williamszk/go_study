// https://www.udemy.com/course/learn-how-to-code/learn/lecture/11922206#overview

package main

import "fmt"

type Person struct {
	firstname string `json:"firstname"`
	lastname  string `json:"lastname"`
}

//
type SecretAgent struct {
	// person        Person `json:"Person"`
	Person
	licenseToKill bool `json:"LicenseToKill"`
}

func main() {

	p1 := Person{
		firstname: "James",
		lastname:  "Bond",
	}

	p2 := Person{
		firstname: "Miss",
		lastname:  "Moneypenny",
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.firstname, p1.lastname)

	// https://www.udemy.com/course/learn-how-to-code/learn/lecture/11922208#overview

	// secret agent
	sa1 := SecretAgent{
		Person: Person{
			firstname: "James",
			lastname:  "Bond",
		},
		licenseToKill: true,
	}

	fmt.Println(sa1)
	// note that we can call the field firstname that belongs to the Person struct
	// we can do this because when defining SecretAgent we do not include a field name
	// just its type
	// this is possible when we have a struct inside another struct
	fmt.Println(sa1.firstname)
	fmt.Println(sa1.lastname)
	// we say that the firstname got promoted from the inner type to the outertype
	// this will happen when there is no field firstname in the outertype

}

// next video:
// https://www.udemy.com/course/learn-how-to-code/learn/lecture/11922210#overview
