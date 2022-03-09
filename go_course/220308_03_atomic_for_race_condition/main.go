package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func main() {
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	numgoroutines := 100

	var counter int64
	counter = 0

	for i := 0; i < numgoroutines; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			holder := counter
			time.Sleep(5)
			counter = holder
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	fmt.Println("Counter:", counter)
}

// the solution is wrong has something to do with wait groups
// what wait groups do?
