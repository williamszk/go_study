// https://www.youtube.com/watch?v=egd4WHJMwC0&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=130&ab_channel=AprendaGo
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	NumGoRountines := 50
	wg.Add(NumGoRountines)

	fmt.Println("Num of CPUs: ", runtime.NumCPU())
	fmt.Println("Num of goroutines: ", runtime.NumGoroutine())

	counter := 0

	var mu sync.Mutex

	for i := 0; i < NumGoRountines; i++ {
		go func() {
			mu.Lock()
			v := counter

			// runtime.Gosched()
			time.Sleep(500)
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Num of goroutines: ", runtime.NumGoroutine())
	}
	fmt.Println("Num of goroutines: ", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println(counter)
}

// go run --race main.go
