package main

import "fmt"

func main() {

	c := incrementor()
	cSum := puller(c)

	for n := range cSum {
		fmt.Println(n)
	}
}

// incrementor will create a channel with some important data
func incrementor() chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

// puller uses a channel to gather ints and sum them, the summed ints are returned
// as another channel
func puller(c chan int) chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}