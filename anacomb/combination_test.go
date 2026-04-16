package anacomb

import "testing"

func TestIndexesInitialization(t *testing.T) {
	i := initializeIndexes(3, 2)
	AssertEqualSlice([]int{0, 1}, i.iArr, t)
	AssertEquals(1, i.iArrLastIdx, t)
}


func TestIndexesUpdating_2Indexes_SourceLen3(t *testing.T) {
	i := initializeIndexes(3, 2)
	AssertEqualSlice([]int{0, 1}, i.iArr, t)

	updated := i.update()
	AssertTrue(updated, t)
	AssertEqualSlice([]int{0, 2}, i.iArr, t)

	updated = i.update()
	AssertTrue(updated, t)
	AssertEqualSlice([]int{1, 2}, i.iArr, t)

	updated = i.update()
	AssertFalse(updated, t)
}

func TestIndexesUpdating_3Indexes_SourceLen4(t *testing.T) {
	i := initializeIndexes(4, 3)
	AssertEqualSlice([]int{0, 1, 2}, i.iArr, t)

	updated := i.update()
	AssertTrue(updated, t)
	AssertEqualSlice([]int{0, 1, 3}, i.iArr, t)

	updated = i.update()
	AssertTrue(updated, t)
	AssertEqualSlice([]int{0, 2, 3}, i.iArr, t)

	updated = i.update()
	AssertTrue(updated, t)
	AssertEqualSlice([]int{1, 2, 3}, i.iArr, t)

	updated = i.update()
	AssertFalse(updated, t)
}

func TestIndexesUpdating_4Indexes_SourceLen6(t *testing.T) {
	i := initializeIndexes(6, 4)
	AssertEqualSlice([]int{0, 1, 2, 3}, i.iArr, t) // 1st combination

	updated := i.update() // 2nd combination
	AssertTrue(updated, t)
	AssertEqualSlice([]int{0, 1, 2, 4}, i.iArr, t)

	updated = i.update() // 3rd combination
	AssertTrue(updated, t)
	AssertEqualSlice([]int{0, 1, 2, 5}, i.iArr, t)

	updated = i.update() // 4th combination
	AssertTrue(updated, t)
	AssertEqualSlice([]int{0, 1, 3, 4}, i.iArr, t)

	updated = i.update() // 5th combination
	AssertTrue(updated, t)
	AssertEqualSlice([]int{0, 1, 3, 5}, i.iArr, t)

	// I'll check by math induction
	i.update() // 6th combination
	i.update() // 7th combination
	updated = i.update() // 8th combination
	AssertTrue(updated, t)
	AssertEqualSlice([]int{0, 2, 3, 5}, i.iArr, t)

	// 7 combinations left
	for range 7 {
		updated = i.update()
		AssertTrue(updated, t)
	}

	// finished
	updated = i.update()
	AssertFalse(updated, t)
}

func TestCombination(t *testing.T) {
	t.Skip()
	count, _ := CountCombination(4, 2)
	AssertEquals(6, count, t)

	count, _ = CountCombination(5, 2)
	AssertEquals(10, count, t)
}

func TestCombine_ArrayOf3_TakenBy2(t *testing.T) {
	t.Skip()
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
	t.Skip()
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
	t.Skip()
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
