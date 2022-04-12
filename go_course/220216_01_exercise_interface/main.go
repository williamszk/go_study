// https://www.youtube.com/watch?v=3pbmLfcgN2s&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=103&ab_channel=AprendaGo
package main

import (
	"fmt"
	"math"
)

type square struct {
	sideLen float64
}

type circle struct {
	radius float64
}

func (geomObj square) area() float64 {
	return math.Pow(geomObj.sideLen, 2)
}

func (geomObj circle) area() float64 {
	return 2 * geomObj.radius * math.Pi
}

type GeomFig interface {
	area()
}

func areaGeomObj(aGeomObj GeomFig) float64 {
	return areaGeomObj.(square).area()
	// switch aGeomObj.(type) {
	// case square:
	// 	return aGeomObj.(square).sideLen
	// case circle:
	// 	return aGeomObj.(circle).radius
	// }
}

func main() {
	asquare := square{
		sideLen: 10.0,
	}

	acircle := circle{
		radius: 3.0,
	}

	fmt.Println("Area of the square", asquare.area())
	fmt.Println("Area of the circle", acircle.area())

	fmt.Println("Area using polymorphism:", areaGeomObj(asquare))
}

// I don't know what is wrong
