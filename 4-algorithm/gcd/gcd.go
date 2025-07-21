package gcd

func GCDByEuclid(a, b int) int {
	var greater int
	remainder := 1

	if a == 0 && b == 0 {
		return 0
	}

	if a != 0 && b == 0 {
		return a
	}

	if a == 0 && b != 0 {
		return b
	}

	if a > b {
		greater = a
		remainder = b
	} else {
		greater = b
		remainder = a
	}

	for remainder != 0 {
		oldRemainder := remainder
		remainder = greater % remainder
		greater = oldRemainder
	}

	return greater
}
