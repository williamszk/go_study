package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// --- set up log file ---
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	// --- create an error ---
	f2, err := os.Open("no-file.txt")
	if err != nil {
		// fmt.Println("Sorry, an error happened:", err)
		// log.Println("Sorry, an error happened:", err)
		// log.Fatalln("Sorry, an error happened:", err)
		panic(err)
	}
	defer f2.Close()
	fmt.Println("Check the log.txt")
}
