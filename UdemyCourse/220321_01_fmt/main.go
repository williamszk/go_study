// https://www.udemy.com/course/learn-how-to-code/learn/lecture/11922042?start=0#overview
package main

import "fmt"

var y = 42

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)

	y = 911
	fmt.Printf("%#x\n", y)
	fmt.Printf("%#X\n", y)
	fmt.Printf("%#X\t%b\t%X\n", y, y, y)

	mystring := fmt.Sprintf("%#X\t%b\t%X\n", y, y, y)

	fmt.Println(mystring)

}
