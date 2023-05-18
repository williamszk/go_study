package dog

import (
	"fmt"
	"testing"
)

func TestYear(t *testing.T) {
	result := Years(10)
	expected := 70
	if result != expected {
		t.Error("Expected", expected, "Got", result)
	}
}

func TestYearsTwo(t *testing.T) {
	result := YearsTwo(9)
	expected := 63
	if result != expected {
		t.Error("Expected", expected, "Got", result)
	}
}

func ExampleYears() {
	fmt.Println(Years(8))
	// Output:
	// 56
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(13))
	// Output:
	// 91
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(8)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(8)
	}
}
