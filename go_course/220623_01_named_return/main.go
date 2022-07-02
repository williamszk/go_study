// https://www.youtube.com/watch?v=rnq5A-yuhkA&ab_channel=HunCoding
package main

import "fmt"

func main() {

	fmt.Println("Call a1:", a1())
	fmt.Println("Call a2:", a2())

}

func a1() int {
	return 2
}

// In this case the return value of the function
// is already defined on top with the function definition
func a2() (myReturn int) {
	myReturn = 10
	return
}

// note that we don't have to tell again what the return
// variable is, it is declared on top
// this is called named returns
