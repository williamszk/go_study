package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// Make a race condition for a variable being incremented.
func main() {
	ctr := 0
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			hold := ctr
			time.Sleep(2)
			hold++
			ctr = hold
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter: ", ctr)
}

// how to build checking for race condition?
// go build -race
// go run main.go -race
