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
	totalFuncs := 3

	go send(20, channel01)

	go another(totalFuncs, channel01, channel02)

	for v := range channel02 {
		fmt.Println(v)
	}
	// if there is not signal that the channel is close then
	// we'll get error: fatal error: all goroutines are asleep - deadlock!
}

func send(n int, channel01 chan int) {
	for i := 0; i < n; i++ {
		channel01 <- i
	}
	close(channel01)
}

func another(totalFuncs int, channel01, channel02 chan int) {

	var wg sync.WaitGroup

	for i := 0; i < totalFuncs; i++ {
		wg.Add(1)
		go func() {
			for v := range channel01 {
				channel02 <- work(v)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	close(channel02)
}

func work(x int) int {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(5e3)))
	return x
}
