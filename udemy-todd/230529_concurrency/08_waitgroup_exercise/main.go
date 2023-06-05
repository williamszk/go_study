package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	const numGoroutines = 100
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func(i int) {
			// fmt.Println("Message from \t", i)
			fmt.Println("Goroutines\t", runtime.NumGoroutine())
			wg.Done()
		}(i)
	}

	wg.Wait()
}
