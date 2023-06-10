package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// send
	go send(even, odd, quit)

	// receive

	// while the go send is running the receive will block and will be
	// waiting for another goroutine to fill the channel
}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}

func receive(e, o, q <-chan int) {
	// an infinite loop
	for {
		select {
		case v := <-e:
			fmt.Println("From even channel", v)
		case v := <-o:
			fmt.Println("From odd channel", v)
		case v := <-q:
			fmt.Println("From quit channel", v)
			// close(q)
			return
		}
	}
}
