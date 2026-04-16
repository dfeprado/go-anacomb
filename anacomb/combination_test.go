package anacomb

import "testing"

func TestCombination(t *testing.T) {
	count, _ := CountCombination(4, 2)
	AssertEquals(6, count, t)

	count, _ = CountCombination(5, 2)
	AssertEquals(10, count, t)
}

func TestCombine_ArrayOf3_TakenBy2(t *testing.T) {
	input := []int{1, 2, 3}
	expected := [][]int{
		{1, 2},
		{1, 3},
		{2, 3},
	}
	combinations, err := Combine(input, 2)
	AssertNil(err, t)
	AssertEquals(3, len(combinations), t)
	for idx := range combinations {
		AssertEqualSlice(expected[idx], combinations[idx], t)
	}
}

func TestCombine_ArrayOf4_TakenBy2(t *testing.T) {
	input := []int{1, 2, 3, 4}
	expected := [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{2, 3},
		{2, 4},
		{3, 4},
	}
	combinations, err := Combine(input, 2)
	AssertNil(err, t)
	AssertEquals(6, len(combinations), t)
	for idx := range combinations {
		AssertEqualSlice(expected[idx], combinations[idx], t)
	}
}

func TestCombine_ArrayOf4_TakenBy3(t *testing.T) {
	input := []int{1, 2, 3, 4}
	expected := [][]int{
		{1, 2, 3},
		{1, 2, 4},
		{1, 3, 4},
		{2, 3, 4},
	}
	combinations, err := Combine(input, 3)
	AssertNil(err, t)
	AssertEquals(4, len(combinations), t)
	for idx := range combinations {
		AssertEqualSlice(expected[idx], combinations[idx], t)
	}
}
