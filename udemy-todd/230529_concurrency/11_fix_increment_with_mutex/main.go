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

	var mu sync.Mutex

	for i := 0; i < numGoroutines; i++ {
		go func() {
			mu.Lock()
			hold := counter
			hold++
			runtime.Gosched()
			counter = hold
			mu.Unlock()
			wg.Done()
			fmt.Println("Goroutines\t", runtime.NumGoroutine())
		}()
	}
	wg.Wait()
	fmt.Println("Counter:", counter)
}
