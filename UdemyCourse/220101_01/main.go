// https://www.udemy.com/course/learn-how-to-code/learn/lecture/11922034?start=0#overview

package main

import "fmt"

var z = 90

func main() {
	x := 42 // short declaration operator
	// this is declaring the variable

	fmt.Println(x)

	// assign a new value to x
	x = 99 // this is assigning a new value to the variable
	// there is a difference between declaration and assignment
	fmt.Println(x)

	// there is a difference between DECLARE a variable
	// and ASSIGN a value to a variable

	var y = 43
	fmt.Println(y)
	// try to short the scope of variables
	// short declaration operator cannot be used outside a function body
	//

	foo()
	// we could say that INITIALIZE a variable is like
	// DELCARE and ASSIGN at the same time

	// PRIMITIVE DATATYPES
	// COMPOSITE DATATYPES

	// there is a difference between PRIMITIVE DATATYPES and
	// BUILT-IN DATATYPES
	// usually PRIMITIVE DATATYPES are build-in to the programming language

	// we can assign the print value to a variable using "Sprintf"
	myString := fmt.Sprintf("%v\t%T", y, y)
	fmt.Println(myString)

}

func foo() {
	fmt.Println("This is a global variable: ", z)
	// we can only use the variable z because it was declared in
	// the global scope
	// we can say that the scope of z is package level
	// each file is called a package
	// differently from other programming languages, in python we call script
	// in Java we call a file "a class", in C# I think it is a namespace.

}
