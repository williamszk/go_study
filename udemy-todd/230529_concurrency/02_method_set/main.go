package main

import (
	"fmt"
	"math"
)

// func (t T) myFunc
// T *T

// pointer receiver
// func (t *T) myFunc
// we can actually use T and *T here
// but it does not count for the case of interface

func main() {
	c := circle{5}
	// info(c) // this will break
	info(&c) // this doesn't break
	fmt.Println(c.area())
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

// area has a pointer receiver
func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
// this methods means that only *T implements the shape interface
// but T does not implement the interface

func info(s shape) {
	fmt.Println("area", s.area())
}
