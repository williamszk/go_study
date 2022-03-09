package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var mutex sync.Mutex

func main() {
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	numgoroutines := 100

	counter := 0

	for i := 0; i < numgoroutines; i++ {
		go func() {
			mutex.Lock()
			holder := counter
			time.Sleep(5)
			holder++
			counter = holder
			mutex.Unlock()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	fmt.Println("Counter:", counter)
}

// this solution is wrong and idk why
// it should not have race condition
// go run --race main.go
// idk what wait groups do...
