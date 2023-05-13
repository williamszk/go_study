// https://www.udemy.com/course/learn-how-to-code/learn/lecture/11922058#overview
package main

import "fmt"

var (
	x int
	y string
	z bool
)

func main() {
	// the default values are called zero value
	// they depend on the type, it could be 0, "", 0.0, false
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	// x: 0
	// y:
	// z: false
}
