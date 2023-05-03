package main

import (
	"fmt"
	"time"
	"williamszk/concurrency_bank/bank"
)

func main() {
	// Alice:
	go func() {
		bank.Deposit(200)                  // A1
		fmt.Println("2 =", bank.Balance()) // A2
	}()

	// Bob:
	go bank.Deposit(100) // B
	fmt.Println("1 =", bank.Balance())

	time.Sleep(2 * time.Second)
}
