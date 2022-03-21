package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel() // cancel when we are finished

	dst := gen(ctx)
	for n := range dst {
		fmt.Println(n)
		if n == 5 {
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
				return // returning not to leak the goroutine
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}

// return some classes to the past:
// https://www.youtube.com/watch?v=jnnIgvV0_yA&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=113&ab_channel=AprendaGo
