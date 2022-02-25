package main

import (
	"fmt"

	"github.com/williamszk/go_study/go_course/220224_04_packages/pkg2"
)

func main() {
	fmt.Println("This is the main function.")
	AnotherFunc()
	pkg2.MyPkg2Func()
}

// to run with all files we need to do:
// go run *.go
// go run !(*).go
