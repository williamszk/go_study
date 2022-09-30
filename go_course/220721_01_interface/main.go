// https://www.udemy.com/course/learn-how-to-code/learn/lecture/4077396#overview

package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float64
}

// an interface defines a set of functionalities
// any type that attends to this set of functionalities
// will attend to the interface
// in this way we can specify that an object should
// respect a certain interface to be a valid argument e.g.
// instead of specifying a struct type
type Shape interface {
	area() float64
}

func (s Square) area() float64 {
	return math.Pow(s.side, 2)
}

func info(z Shape) {
	fmt.Println(z.area())
}

type Point struct {
	lat int64
	lon int64
}

// Polygon is a representation of a geographic polygon.
// Polygon is composed by a slice of points. The order of the points matter, they should be counter clockwise.
// If the points are ordered in the clockwise they will generatea a negative area, that could be interpreted
// as a lake e.g.
type Polygon struct {
	points []Point
}

func (p Polygon) area() float64 {

	// ideally in here we should implemenet the correct formula for polygon area
	return 10.0
}

func main() {
	s := Square{
		side: 10,
	}
	// Square is a Shape
	// Interfaces work almost like parent classes in this sense
	// and also Squre is a polygon
	info(s)
	// we can say that Square is the "Concrete Type"
	// and Shape is the "Abstract Type"
	// Interface define behavior

	p := Polygon{
		points: []Point{{10, 10}, {20, 20}},
	}

	info(p)
}

// We could say that using interfaces implement polymorphism
// We could also say that interfaces implement substitutability
