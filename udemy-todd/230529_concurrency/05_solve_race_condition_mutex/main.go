package main

import (
	"fmt"
	"runtime"
	"sync"
)

// create a race condition

func main() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)

	fmt.Println("CPU\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			// we need a way for the goruntime to do somehting else
			// time.Sleep(time.Second*2) // this is one way
			runtime.Gosched() // this is another way
			v++
			counter = v
			mu.Unlock()
			wg.Done()
			fmt.Println("Goroutines\t", runtime.NumGoroutine())
		}()
	}

	fmt.Println("CPU\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("counter\t", counter)
}
