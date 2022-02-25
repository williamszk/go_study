// https://www.youtube.com/watch?v=0qGILXmLfMM&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=127&ab_channel=AprendaGo
// https://www.youtube.com/watch?v=XxG7qqJzDKk&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=129&ab_channel=AprendaGo

// race condition error

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	totalgoroutines := 500

	counter := 0

	wg.Add(totalgoroutines)

	for i := 0; i < totalgoroutines; i++ {
		go func() {
			holder := counter
			// runtime.Gosched()
			holder++
			counter = holder
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("counter:", counter)
}

// mutex : multual exclusion
// atomic
// go channels

// next:
// https://www.youtube.com/watch?v=egd4WHJMwC0&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=129&ab_channel=AprendaGo
