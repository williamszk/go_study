package main

import "fmt"

func main() {

	chneven := make(chan int)
	chnodd := make(chan int)
	chnquit := make(chan bool)

	go sendNums(chneven, chnodd, chnquit)

	receiveNums(chneven, chnodd, chnquit)

}

func sendNums(chneven chan int, chnodd chan int, chnquit chan bool) {

	total := 100

	for i := 0; i < total; i++ {
		if i%2 == 0 {
			chneven <- i
		} else {
			chnodd <- i
		}
	}

	close(chneven)
	close(chnodd)
	chnquit <- true

}

func receiveNums(chneven chan int, chnodd chan int, chnquit chan bool) {
	for {
		select {
		case v := <-chneven:
			fmt.Println("Received number from even:", v)
		case v := <-chnodd:
			fmt.Println("Received number from odd:", v)
		case <-chnquit:
			return
		}
	}
}

// https://www.youtube.com/watch?v=wWQ0BbbQ-28&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=150&ab_channel=AprendaGo
// this class wasn't so clear...
