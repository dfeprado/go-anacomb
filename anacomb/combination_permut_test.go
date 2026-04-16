package anacomb

import "testing"

func TestPermutInit(t *testing.T) {
	i := initPermut(3, 2)
	AssertEqualSlice([]int{0, 1}, i.iArr, t)
	AssertEquals(1, i.iArrLastIdx, t)
}

func TestPermut_2Parts_SourceLen3(t *testing.T) {
	i := initPermut(3, 2)
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

func TestPermut_3Parts_SourceLen4(t *testing.T) {
	i := initPermut(4, 3)
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

func TestPermut_4Parts_SourceLen6(t *testing.T) {
	i := initPermut(6, 4)
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
	i.update()           // 6th combination
	i.update()           // 7th combination
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
