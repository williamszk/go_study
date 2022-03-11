package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	chnres := converges(work("1"), work("2"))

	for i := 0; i < 20; i++ {
		fmt.Println(<-chnres)
	}

}

func work(mymess string) chan string {
	channel := make(chan string)
	go func(mymess string, channel chan string) {
		for i := 0; ; i++ {
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(5e3)))
			channel <- fmt.Sprintf("Function %v say:\t %v", mymess, i)
		}
	}(mymess, channel)

	return channel
}

func converges(chn1, chn2 chan string) chan string {
	newchn := make(chan string)

	go func() {
		for {
			newchn <- <-chn1
		}
	}()

	go func() {
		for {
			newchn <- <-chn2
		}
	}()

	return newchn
}

// next
// https://www.youtube.com/watch?v=8X6eOnSJu5g&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=152&ab_channel=AprendaGo
//
