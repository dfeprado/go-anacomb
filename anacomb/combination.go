package anacomb

type indexes struct {
	iArr        []int
	iArrLastIdx int
	sourceLen   int
	eol         bool
}

func (i *indexes) update() bool {
	if i.eol {
		return false
	}

	pivotIdx := i.iArrLastIdx
LL:
	i.iArr[pivotIdx] += 1
	if i.iArr[pivotIdx] < i.sourceLen {
		if pivotIdx == i.iArrLastIdx {
			return true
		}

		for idx := pivotIdx + 1; idx < len(i.iArr); idx++ {
			i.iArr[idx] = i.iArr[idx-1] + 1
		}

		if pivotIdx > 0 && i.iArr[i.iArrLastIdx] >= i.sourceLen {
			pivotIdx -= 1
			goto LL
		}

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

func initializeIndexes(sourceLen int, parts int) *indexes {
	i := &indexes{
		iArr:        make([]int, parts),
		iArrLastIdx: parts - 1,
		sourceLen:   sourceLen,
	}

	for idx := range parts {
		i.iArr[idx] = idx
	}

	return i
}

func Combine(A []int, parts int) ([][]int, error) {
	combinations, err := CountCombination(len(A), parts)
	if err != nil {
		return nil, err
	}

	combIndexes := initializeIndexes(len(A), parts)
	result := make([][]int, 0, combinations)
	for combAvail := true; combAvail; combAvail = combIndexes.update() {
		combination := make([]int, parts)
		for idx, idxPointer := range combIndexes.iArr {
			combination[idx] = A[idxPointer]
		}
		result = append(result, combination)
	}

	return result, nil
}

func CountCombination(n, parts int) (int, error) {
	nFac, err := Factorial(n)
	if err != nil {
		return 0, err
	}
	nParts, err := Factorial(parts)
	if err != nil {
		return 0, err
	}
	n_minus_parts_fac, _ := Factorial(n - parts)

	return nFac / (nParts * n_minus_parts_fac), nil
}
