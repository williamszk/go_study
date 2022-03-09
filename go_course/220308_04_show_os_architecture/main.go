package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("OS is:", runtime.GOOS)
	fmt.Println("The Arch is:", runtime.GOARCH)
}

// go run main.go
// go build main.go
// ./main.exe
