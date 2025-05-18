package main

import "fmt"

func main() {
	const usdToEur = 0.9
	const usdToRub = 81.07

	eurToRub := (1 / usdToEur) * usdToRub

	fmt.Println(eurToRub)
}
