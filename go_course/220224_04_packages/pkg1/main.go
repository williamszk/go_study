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

// https://rakyll.org/style-packages/
// next
// https://www.youtube.com/watch?v=Y9QEvz4D_9E&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=140&ab_channel=AprendaGo
