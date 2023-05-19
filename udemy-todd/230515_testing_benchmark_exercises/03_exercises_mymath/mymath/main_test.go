package mymath

import (
	"fmt"
	"testing"
)

// test
// example
// benchmark

func TestCenteredAvg(t *testing.T) {
	xi := []int{1, 2, 3, 4, 5}
	result := CenteredAvg(xi)
	expected := 3.0
	if result != expected {
		t.Error("Expected", expected, "Got", result)
	}
}

func TestCenteredAvg_table(t *testing.T) {

	type test struct {
		data   []int
		answer float64
	}

	tests := []test{
		{[]int{0, 1, 1, 0}, 0.5},
		{[]int{2, 3, 4, 5, 0}, 3.0},
		{[]int{99, 1, 1, 1}, 1},
		{[]int{-1, 0, 1}, 0},
		{[]int{0, 0, 0, 0, 0, 0, 0}, 0},
	}

	for _, v := range tests {
		result := CenteredAvg(v.data)
		expected := v.answer
		if result != expected {
			t.Error("Expected", expected, "Got", result)
		}
	}
}

func ExampleCenteredAvg() {
	xi := []int{1, 1, 2, 3, 4, 5}
	fmt.Println(CenteredAvg(xi))
	// Output:
	// 2.5
}

func ExampleCenteredAvg_second() {
	xi := []int{-1, 5, 1, 2, 3, 4, 5}
	fmt.Println(CenteredAvg(xi))
	// Output:
	// 3
}

func BenchmarkCenteredAvg(b *testing.B) {
	xi := []int{-1, 5, 1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		CenteredAvg(xi)
	}
}
