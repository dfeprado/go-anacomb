
package anacomb

import "testing"

func TestFactorial(t *testing.T) {
	facN, _ := Factorial(4)
	AssertEquals(24, facN, t)

	facN, _ = Factorial(5)
	AssertEquals(120, facN, t)
}

func TestFactorialOfZero(t *testing.T) {
	facN, _ := Factorial(0)
	AssertEquals(1, facN, t)
}

func TestFactorialOfOneAndTwo(t *testing.T) {
	facN, _ := Factorial(1)
	AssertEquals(1, facN, t)

	facN, _ = Factorial(2)
	AssertEquals(2, facN, t)
}

func TestFactorialError(t *testing.T) {
	fac, err := Factorial(-1)
	AssertEquals(0, fac, t)
	AssertNotNil(err, t)
	AssertEquals("x must be >= 0", err.Error(), t)
}
