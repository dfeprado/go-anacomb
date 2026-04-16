package anacomb

func Combine[K any](A []K, parts int) ([][]K, error) {
	combinations, err := CountCombination(len(A), parts)
	if err != nil {
		return nil, err
	}

	permut := initPermut(len(A), parts)
	result := make([][]K, 0, combinations)
	for combAvail := true; combAvail; combAvail = permut.update() {
		combination := make([]K, parts)
		for idx, idxPointer := range permut.iArr {
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
