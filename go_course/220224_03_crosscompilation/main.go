// https://www.youtube.com/watch?v=k0pV_JoiZbI&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=137&ab_channel=AprendaGo

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Hello World.")
	fmt.Println("This program was compiled in a amd64 processor architecture with Windows OS and is now running in a", runtime.GOARCH, runtime.GOOS)
}

// compile the program to linux amd64
// GOOS=linux GOARCH=amd64 go build main.go
