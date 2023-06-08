package main

import "fmt"

// this program will use the values from ch but it will not modify it
// then this program will output the values of ch multiplied by 2
func modifyChannel(ch chan int) chan int {
	newCh := make(chan int)
	go func() {
		for value := range ch {
			newCh <- value * 2
		}
		close(newCh)
	}()
	return newCh
}

// what is the purpose of this program?
func main() {
	ch := make(chan int)

	go func() {
		for value := range modifyChannel(ch) {
			fmt.Println(value)
		}
	}()

	ch <- 21
	close(ch)
}
