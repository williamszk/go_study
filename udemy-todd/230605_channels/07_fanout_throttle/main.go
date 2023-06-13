// The only difference in this case is that instead of 100 goroutines we want
// to use only 10 goroutines
// we say that this is the throttling case because we control the number of
// goroutines that we can use
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)

	go fanOutIn(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}
	fmt.Println("About to finish")
}

func populate(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup
	const numGoroutines = 10
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			// this range down here will block and be listening for any
			// data comes in, each of the 10 goroutines will consume
			// the incoming data from the c1 channel
			for v := range c1 {
				func(v2 int) {
					c2 <- timeConsumingWork(v2)
				}(v)
			}
			// the advantage of using channels like queues is that we don't need
			// to worry about the amount of time that each task could take because
			// it can take different amount of time
			wg.Done()
		}()
	}
	wg.Wait()
	close(c2)
}

func timeConsumingWork(v int) int {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
	return v + rand.Intn(1000)
}
