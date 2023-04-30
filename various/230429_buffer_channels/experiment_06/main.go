package main

import (
	"fmt"
	"time"
)

// What if then we have lots of go routines being create indefinitely?
// What the read of myChan will do?
// I believe it will wait, if no value could be read from the channel
// and there are no other goroutines running
// then it knows that this channel will not be filled
// and the program will crash
func main() {

	myChan := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			time.Sleep(1 * time.Second)
			fmt.Println("Create a goroutine!", i)
			go func(idx int) {
				time.Sleep(5 * time.Second)
				fmt.Println("End a goroutine!", idx)
			}(i)
		}

	}()

	myVal := <-myChan

	fmt.Println(myVal)
}

// Create a goroutine! 1
// Create a goroutine! 2
// Create a goroutine! 3
// Create a goroutine! 4
// Create a goroutine! 5
// End a goroutine! 1
// Create a goroutine! 6
// End a goroutine! 2
// Create a goroutine! 7
// End a goroutine! 3
// Create a goroutine! 8
// End a goroutine! 4
// Create a goroutine! 9
// End a goroutine! 5
// Create a goroutine! 10
// End a goroutine! 6
// End a goroutine! 7
// End a goroutine! 8
// End a goroutine! 9
// End a goroutine! 10
// fatal error: all goroutines are asleep - deadlock!

// goroutine 1 [chan receive]:
// main.main()
//         /root/go_study/various/230429_buffer_channels/experiment_06/main.go:29 +0x4f
// exit status 2
