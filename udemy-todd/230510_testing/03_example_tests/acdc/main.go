// Package acdc asks if you are "ready to rock!"
package acdc

// Sum function adds an unlimited number of values of type int.
func Sum(xi ...int) int {
	sum := 0

	for _, y := range xi {
		sum += y
	}

	return sum
}
