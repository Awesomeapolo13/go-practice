package main

import "fmt"

func main() {
	moneyAmount, initialCurrency, targetCurrency := getUserInput()

	msg := fmt.Sprintf("You want to convert %.0f %s to %s", moneyAmount, initialCurrency, targetCurrency)
	fmt.Println(msg)

	const usdToEur = 0.9
	const usdToRub = 81.07
	eurToRub := (1 / usdToEur) * usdToRub

	fmt.Println(eurToRub)
}

func getUserInput() (float64, string, string) {
	var moneyAmount float64
	var initialCurrency, targetCurrency string

	fmt.Println("__Welcome to the currency converter__")
	fmt.Print("Set an amount of money you want to convert: ")
	fmt.Scan(&moneyAmount)
	fmt.Print("Set a currency you want to convert from: ")
	fmt.Scan(&initialCurrency)
	fmt.Print("Set a currency you want to convert to: ")
	fmt.Scan(&targetCurrency)

	return moneyAmount, initialCurrency, targetCurrency
}

func convertAmount(amount float64, initialCurrency string, targetCurrency string) float64 {
}
