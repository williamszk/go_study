package math

import "testing"

type AddData struct {
	x, y   int
	result int
}

// this is the function signature for testing functions
func TestAdd(t *testing.T) {

	// result := Add(1, 3)

	// if result != 4 {
	// 	t.Errorf("Add(1, 3) FAILED. Expected %d, got %d\n", 4, result)
	// } else {
	// 	t.Logf("Add(1, 3) PASSED. Expected %d, got %d\n", 4, result)
	// }

	testData := []AddData{
		{1, 2, 3},
		{3, 5, 8},
		{7, -4, 3},
	}

	for _, d := range testData {
		result := Add(d.x, d.y)

		if result != d.result {
			t.Errorf("Add(%d, %d) FAILED. Expected %d, got %d\n",
				d.x, d.y, d.result, result)
		} else {
			t.Logf("Add(%d, %d) PASSED. Expected %d, got %d\n",
				d.x, d.y, d.result, result)
		}
	}
}

func TestDivide(t *testing.T) {

	result := Divide(5, 0)

	if result != 0 {
		t.Errorf("Divide(5, 0) FAILED. Expected %f, got %f\n", 0.0, result)
	} else {
		t.Logf("Divide(5, 0) PASSED. Expected %f, got %f\n", 0.0, result)
	}

}

// to run the tests inside the package
// we go to the directory and run
// go test
// the file needs to end with _test
// and the functions need to start with Test<rest of the name>

// we can also use verbose
// go test -v

// go test -v --cover
