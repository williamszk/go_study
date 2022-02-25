// https://www.youtube.com/watch?v=iFlQ2yAYcp4&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=130&ab_channel=AprendaGo

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	totalgoroutines := 100

	var counter int64
	counter = 0

	wg.Add(totalgoroutines)

	for i := 0; i < totalgoroutines; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			holder := counter
			runtime.Gosched()
			counter = holder
			fmt.Println("Counter: \t", atomic.LoadInt64(&counter))
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("counter:", counter)
}

// go run --race main.go

// next
// https://www.youtube.com/watch?v=1E3KMnYnqCk&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=133&ab_channel=AprendaGo
