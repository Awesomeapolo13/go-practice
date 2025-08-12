package pow

func RaiseIntToIntPositivePower(num, power int) int {

	if num < 0 {
		panic("num cannot be negative")
	}

	if power == 0 {
		return 1
	}

	if num == 0 {
		return 0
	}

	isEven := power%2 == 0
	if !isEven {
		return num * RaiseIntToIntPositivePower(num, power-1)
	}

	if power == 2 {
		return powTo2(num)
	}

	return powTo2(RaiseIntToIntPositivePower(num, power/2))
}

func powTo2[T int | float64](num T) T {
	return num * num
}
