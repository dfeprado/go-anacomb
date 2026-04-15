package anacomb

import "testing"

func TestCombination(t *testing.T) {
	count, _ := CountCombina(4, 2)
	AssertEquals(6, count, t)

	count, _ = CountCombina(5, 2)
	AssertEquals(10, count, t)
}
