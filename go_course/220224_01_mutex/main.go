// https://www.youtube.com/watch?v=egd4WHJMwC0&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=129&ab_channel=AprendaGo

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	totalgoroutines := 100

	counter := 0

	wg.Add(totalgoroutines)

	for i := 0; i < totalgoroutines; i++ {
		go func() {
			mu.Lock()
			holder := counter
			runtime.Gosched()
			holder++
			counter = holder
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("counter:", counter)
}

// go run --race main.go
