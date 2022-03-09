package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	numgoroutines := 100

	counter := 0

	for i := 0; i < numgoroutines; i++ {
		go func() {
			holder := counter
			time.Sleep(5)
			holder++
			counter = holder
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	fmt.Println("Counter:", counter)
}
