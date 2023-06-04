package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// create a race condition

func main() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)

	fmt.Println("CPU\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	var counter int64 = 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {

			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter \t", atomic.LoadInt64(&counter))
			runtime.Gosched()
			wg.Done()
			fmt.Println("Goroutines\t", runtime.NumGoroutine())
		}()
	}

	fmt.Println("CPU\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("counter\t", counter)
}
