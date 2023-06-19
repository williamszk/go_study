package main

import (
	"fmt"
	"log"
)

// var ErrNorgateMath = errors.New("norgate math: square root of negative number")

func main() {
	// fmt.Printf("%T\n", ErrNorgateMath)
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}

}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, invokeErrNorgateMath(f)
	}
	return 42, nil
}

func invokeErrNorgateMath(v float64) error {
	var ErrNorgateMath = fmt.Errorf("norgate math: square root of negative number: %v", v)
	return ErrNorgateMath
}
