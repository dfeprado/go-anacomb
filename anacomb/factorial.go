package anacomb

import "errors"

// TODO implements memoization in a way the dev can choose to enable it or disable it

func Factorial(x int) (int, error) {
	if x < 0 {
		return 0, errors.New("x must be >= 0")
	} else if x == 0 {
		return 1, nil
	} else if x <= 2 {
		return x, nil
	}

	result := 1
	for x > 1 {
		result *= x
		x -= 1
	}
	return result, nil
}
