// https://www.udemy.com/course/learn-how-to-code/learn/lecture/11922390#overview

package main

import "fmt"

func main() {

	myVal := 10

	_, err := foo(myVal)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Finishing the program...")
}

type customErr struct {
	message string
}

func (e customErr) Error() string {
	return fmt.Sprintf("Sorry, something went wrong. (%v)\n", e.message)
}

func RaiseCustomError() error {
	ce := customErr{
		message: "Hi! I'm a custom error message.",
	}
	return ce
}

func foo(val int) (int, error) {
	if val < 99 {
		return 0, RaiseCustomError()
	}
	return 42, nil
}
