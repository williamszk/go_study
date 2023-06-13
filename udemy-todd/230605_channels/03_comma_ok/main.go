// the difference here is that we'll use the comma ok syntax to check if we want to
// end the receiving function, which is the function that is holding the maing
// go routine.
// the close(quit) and comma ok idiom is an alternative to use a semaphore and
// send some value like false to the 'quit' channel
// on a side note: we could use 'range' and 'close' to check if we want to end some
// goroutine that is holding the program

package main

import "fmt"

func main() {

	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(even, odd, quit)
	receive(even, odd, quit)
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("Receiving from even channel\t", v)
		case v := <-o:
			fmt.Println("Receiving from odd channel\t", v)
		case v, ok := <-q:
			// the receive channel will give ok=true when everything is ok
			// if ok=false then !ok is true, and we should end
			if !ok {
				fmt.Println("from comma ok", v, ok)
				return
			}
		}
	}
}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(q) // add this line, this will signal that we should end
}
