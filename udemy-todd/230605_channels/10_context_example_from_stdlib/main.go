package main

import (
	"context"
	"fmt"
)

// the context is used so that we can control the canceling of goroutines
// by sending a message through the context
// by the parent goroutine
// the only other way was to have the goroutine itself cancel itself
// but this can lead to memory leaks
// this happens because the program is still running while the
// other goroutines are running still, but they don't have a way to cancel
// themselves
// so we use this pattern which is by using 'context' on the parent goroutine
// and 'select' inside the child goroutine

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished

	for v := range gen(ctx) {
		fmt.Println(v)
		if v == 10 {
			break
		}
	}

}

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return // return no to leak go routine
			case dst <- n:
				n++
			}
		}
	}()

	return dst
}
