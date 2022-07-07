package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) MakeYoung() {
	p.Age = 10
	fmt.Println(p.Name, "is", p.Age, "years old.")
}

// This is just to show the difference when we pass the pointer reference
// this will change implicitly the properties of the object
func (p *Person) MakeYoungForever() {
	p.Age = 10
	fmt.Println(p.Name, "is", p.Age, "years old.")
}

func Try03() {

	person01 := Person{
		Name: "Bob",
		Age:  90,
	}

	person01.MakeYoung()
	fmt.Println("After MakeYoung():", person01)

	person01.MakeYoungForever()
	fmt.Println("After MakeYoungForever():", person01)
}
