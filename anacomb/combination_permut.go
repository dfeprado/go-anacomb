package anacomb

type permut struct {
	iArr        []int
	iArrLastIdx int
	sourceLen   int
	eol         bool
}

func (i *permut) update() bool {
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
