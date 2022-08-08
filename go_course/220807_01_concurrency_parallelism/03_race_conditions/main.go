package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	NumGoRountines := 100

	fmt.Println("Num of CPUs: ", runtime.NumCPU())
	fmt.Println("Num of goroutines: ", runtime.NumGoroutine())

	counter := 0

	wg.Add(NumGoRountines)

	for i := 0; i < NumGoRountines; i++ {
		go func() {
			v := counter

			// runtime.Gosched()
			time.Sleep(30)
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("Num of goroutines: ", runtime.NumGoroutine())
	}
	fmt.Println("Num of goroutines: ", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println(counter)
}

// go run --race main.go
