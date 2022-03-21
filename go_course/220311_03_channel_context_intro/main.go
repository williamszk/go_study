package main

import (
	"context"
	"fmt"
)

// https://www.youtube.com/watch?v=PhTtrrsUH8c&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=153&ab_channel=AprendaGo
func main() {
	ctx := context.Background()

	fmt.Println("context:\t", ctx)
	fmt.Println("contextt err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("-------------")
}

// context allow us to exchange messages between go routines
