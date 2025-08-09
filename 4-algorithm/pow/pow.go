package pow

func RaiseNumberToIntPositivePower(num, power int) int {
	if power == 0 {
		return 1
	}

	if num == 0 {
		return 0
	}

	isEven := power%2 == 0
	if !isEven {
		return num * RaiseNumberToIntPositivePower(num, power-1)
	}

	if power == 2 {
		return powIntTo2(num)
	}

	return powIntTo2(RaiseNumberToIntPositivePower(num, power/2))
}

func powIntTo2(num int) int {
	return num * num
}
