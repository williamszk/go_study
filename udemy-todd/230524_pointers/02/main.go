// https://www.udemy.com/course/learn-how-to-code/learn/lecture/11922286#overview
package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func changeMe(p *person) {
	(*p).first = "changeMe"
	// we could use p.first here, as a shortcut, syntax sugar
	// p.first = "changeMeSugar"
}

func main() {
	// pay attention how to dereference a struct
	bob := person{
		first: "Bob",
		last:  "Ferguson",
		age:   10,
	}
	fmt.Println("bob before", bob)

	changeMe(&bob)
	fmt.Println("bob after", bob)
}
