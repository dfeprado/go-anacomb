
package anacomb

import "testing"

func TestFactorial(t *testing.T) {
	n := 4
	facN, _ := Factorial(n)
	if facN != 24 {
		t.Errorf("Expected 24, got %v", facN)
	}

	n = 5
	facN, _ = Factorial(n)
	if facN != 120 {
		t.Errorf("Expected 120, got %v", facN)
	}
}

func TestFactorialOfZero(t *testing.T) {
	n := 0
	facN, _ := Factorial(n)
	if facN != 1 {
		t.Errorf("Expected 1, got %v", facN)
	}
}

func TestFactorialOfOneAndTwo(t *testing.T) {
	n := 1
	facN, _ := Factorial(n)
	if facN != 1 {
		t.Errorf("Expected 1, got %v", facN)
	}

	n = 2
	facN, _ = Factorial(n)
	if facN != n {
		t.Errorf("Expected 2, got %v", facN)
	}
}

func TestFactorialError(t *testing.T) {
	fac, err := Factorial(-1)
	if fac != 0 {
		t.Errorf("fac should be 0, but it was %v", fac)
	}
	if err == nil {
		t.Errorf("Error is nil")
	}

	if err.Error() != "x must be >= 0" {
		t.Errorf("Unexpected error \"%v\"", err)
	}
}
