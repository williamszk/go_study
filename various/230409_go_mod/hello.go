package hello

import (
	"rsc.io/quote/v3"
)

func Hello() string {
	// return "Hello, world."
	return quote.HelloV3()
}

func Proverb() string {
	return quote.Concurrency()
}
