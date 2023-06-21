package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	_, err := sqrt(-10.22)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		e := errors.New("Sorry, we don't allow negative values.")
		return 0, sqrtError{lat: "10 N", lon: "90 S", err: e}
	}
	return 42, nil
}

type sqrtError struct {
	lat string
	lon string
	err error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v, %v, %v", se.lat, se.lon, se.err)
}
