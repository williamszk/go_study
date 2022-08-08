// This is executed in serial
package main

import "fmt"

func func1() {
	for i := 0; i < 10; i++ {
		fmt.Println("func1:", i)
	}
}

func func2() {
	for i := 0; i < 10; i++ {
		fmt.Println("func2:", i)
	}
}

func main() {

	func1()
	func2()

}
