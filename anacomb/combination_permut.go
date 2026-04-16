package anacomb

type permut struct {
	iArr        []int
	iArrLastIdx int
	sourceLen   int
	eol         bool
}

/*
	This function allows us to go over all possible

combinations by permuting the pointer indexes.

It returns true if there's still permutations to go, or false
otherwise.
*/
func (i *permut) update() bool {
	if i.eol {
		return false
	}

	// Start by the last index in permut
	pivotIdx := i.iArrLastIdx
LL:
	/* Sum one to it. If the new index is greater
	equal or grater than sourceLen, it means we
	reached source end. So we need to back propagate an
	update for all pointerIndexes. */
	i.iArr[pivotIdx] += 1
	if i.iArr[pivotIdx] < i.sourceLen {
		/* If current pivotIdx is not the last one,
		then we need to backpropagate the updates */
		if pivotIdx == i.iArrLastIdx {
			return true
		}

		/* Backpropagating the updates is a process where
		a pointer index value becomes its prior one + 1.

		After all this, if the last pointer index is still
		equal or greater than sourceLen, we need to repeat the process */
		for idx := pivotIdx + 1; idx < len(i.iArr); idx++ {
			i.iArr[idx] = i.iArr[idx-1] + 1
		}

		if pivotIdx > 0 && i.iArr[i.iArrLastIdx] >= i.sourceLen {
			pivotIdx -= 1
			goto LL
		}

		/* If last pointer index is less than sourceLen,
		it means we still have permutations to go. Otherwise,
		no more permutations is available and we've reached
		the end of it (eol = true) */
		if i.iArr[i.iArrLastIdx] < i.sourceLen {
			return true
		} else {
			i.eol = true
			return false
		}
	}
	pivotIdx -= 1
	goto LL
}

func initPermut(sourceLen int, parts int) *permut {
	i := &permut{
		iArr:        make([]int, parts),
		iArrLastIdx: parts - 1,
		sourceLen:   sourceLen,
	}

	for idx := range parts {
		i.iArr[idx] = idx
	}

	return i
}
