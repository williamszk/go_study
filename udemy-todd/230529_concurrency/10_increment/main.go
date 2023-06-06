package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	counter := 0
	var wg sync.WaitGroup
	const numGoroutines = 100
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			hold := counter
			hold++
			runtime.Gosched()
			counter = hold
			wg.Done()
			fmt.Println("Goroutines\t", runtime.NumGoroutine())
		}()
	}
	wg.Wait()
	fmt.Println("Counter:", counter)
}
