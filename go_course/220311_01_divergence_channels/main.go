package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	channel01 := make(chan int)
	channel02 := make(chan int)

	go send(20, channel01)

	go another(channel01, channel02)

	for v := range channel02 {
		fmt.Println(v)
	}

}

func send(n int, channel chan int) {
	for i := 0; i < n; i++ {
		channel <- i
	}
	close(channel)
}

func another(channel01, channel02 chan int) {
	var wg sync.WaitGroup

	for v := range channel01 {
		wg.Add(1)
		go func(x int) {
			channel02 <- work(x)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(channel02)
}

func work(x int) int {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(5e3)))
	return x
}
