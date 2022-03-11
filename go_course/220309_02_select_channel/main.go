// https://www.youtube.com/watch?v=dp8s5jAc7h0&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=150&ab_channel=AprendaGo
package main

import "fmt"

func main() {
	a := make(chan int)
	b := make(chan int)
	x := 500

	go func(x int) {
		for i := 0; i < x; i++ {
			a <- i
		}
	}(x / 2)

	go func(x int) {
		for i := 0; i < x; i++ {
			b <- i
		}
	}(x / 2)

	for i := 0; i < x; i++ {
		select {
		case v := <-a:
			fmt.Println("Received from channel A:", v)
		case v := <-b:
			fmt.Println("Received from channel B:", v)
		}
	}

}
