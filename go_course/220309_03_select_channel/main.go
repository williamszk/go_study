// https://www.youtube.com/watch?v=dp8s5jAc7h0&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=150&ab_channel=AprendaGo
package main

import "fmt"

func main() {

	chn := make(chan int)
	quit := make(chan int)

	go receiveQuit(chn, quit)

	sendToChannel(chn, quit)

}

func receiveQuit(chn chan int, quit chan int) {

	for i := 0; i < 50; i++ {
		fmt.Println("Rceived:", <-chn)
	}
	quit <- 0

}

func sendToChannel(chn chan int, quit chan int) {

	anything := 0

	for {
		select {
		case chn <- anything:
			anything++
		case <-quit:
			return
		}
	}

}

// still
// https://www.youtube.com/watch?v=dp8s5jAc7h0&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=150&ab_channel=AprendaGo
