package anacomb

func CountCombina(n, parts int) (int, error) {
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
