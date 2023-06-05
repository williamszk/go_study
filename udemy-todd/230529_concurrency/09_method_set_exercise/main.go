package main

import "fmt"

func main() {

	bob := person{
		name: "Bob",
	}

	saySomething(&bob)

	bob.speak()

}

type person struct {
	name string
}

func (p *person) speak() {
	fmt.Println("Hi my name is", p.name)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}
