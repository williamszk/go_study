package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	// go routines can be executed in parallel or not
	// it depends on the architecture of the CPU
	// but we know that this will be executed in different threads
	fmt.Println("Number of CPU cores: ", runtime.NumCPU())
	fmt.Println("Number of goroutines: ", runtime.NumGoroutine()) // should be 1

	wg.Add(2)

	// main is one thread
	go func1() // this is a go routine, it is like a bullet, after the fire we can't control it
	go func2()
	fmt.Println("Number of goroutines: ", runtime.NumGoroutine()) // should be 3

	// we need a way to wait for the go routine
	wg.Wait()
}

func func1() {
	for i := 0; i < 10; i++ {
		fmt.Println("func1:", i)
		time.Sleep(100)
	}
	// finished!
	wg.Done() // this will tell the main thread that this go routine is finished
}

func func2() {
	for i := 0; i < 10; i++ {
		fmt.Println("func2:", i)
		time.Sleep(40)
	}
	wg.Done()
}
