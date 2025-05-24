package main

import "fmt"

const USD, EUR, RUB = "USD", "EUR", "RUB"
const usdToEur, usdToRub = 0.9, 81.07

// Вместо свича использовать мап из валюты
func main() {
	fmt.Println("__Welcome to the currency converter__")

	moneyAmount, initialCurrency, targetCurrency := getUserInput()

	msg := fmt.Sprintf("You want to convert %.0f %s to %s", moneyAmount, initialCurrency, targetCurrency)
	fmt.Println(msg)

	result := convertAmount(moneyAmount, initialCurrency, targetCurrency)

	msg = fmt.Sprintf("Result is %.2f %s", result, targetCurrency)
	fmt.Println(msg)
}

func getUserInput() (float64, string, string) {
	initialCurrency := getInitialCurrencyInput()
	moneyAmount := getAmountInout()
	targetCurrency := getTargetCurrencyInput(initialCurrency)

	return moneyAmount, initialCurrency, targetCurrency
}

func getInitialCurrencyInput() string {
	var initialCurrency, msg string
	for {
		fmt.Print("Set a currency you want to convert from (USD, EUR, RUB are allowed): ")
		fmt.Scan(&initialCurrency)

		if isValidInitialCurrency(initialCurrency) {
			return initialCurrency
		}

		msg = fmt.Sprintf("The currency %s is not allowed. Try again.", initialCurrency)
		fmt.Println(msg)
	}
}

func getAmountInout() float64 {
	var moneyAmount float64
	for {
		fmt.Print("Set an amount of money you want to convert (positive number more than 0): ")
		fmt.Scan(&moneyAmount)

		if isValidAmount(moneyAmount) {
			return moneyAmount
		}

		fmt.Println("The amount is not allowed. Try again.")
	}
}

func getTargetCurrencyInput(initialCurrency string) string {
	var targetCurrency, msg string
	for {
		fmt.Print("Set a currency you want to convert to (Use USD, EUR, RUB except already used on the first step): ")
		fmt.Scan(&targetCurrency)

		if isValidTargetCurrency(targetCurrency, initialCurrency) {
			return targetCurrency
		}

		msg = fmt.Sprintf("The currency %s is not allowed. Try again.", targetCurrency)
		fmt.Println(msg)
	}
}

func convertAmount(amount float64, initialCurrency string, targetCurrency string) float64 {
	return amount * resolveExchangeRate(initialCurrency, targetCurrency)
}

func resolveExchangeRate(initialCurrency string, targetCurrency string) float64 {
	eurToRub := usdToRub / usdToEur
	eurToUsd := 1 / usdToEur
	rubToUsd := 1 / usdToRub
	rubToEur := usdToEur / usdToRub
	exchanges := map[string]float64{
		EUR + ":" + USD: eurToUsd,
		USD + ":" + EUR: usdToEur,
		RUB + ":" + USD: rubToUsd,
		USD + ":" + RUB: usdToRub,
		RUB + ":" + EUR: rubToEur,
		EUR + ":" + RUB: eurToRub,
	}

	exchangeRate := exchanges[initialCurrency+":"+targetCurrency]

	if exchangeRate <= 0 {
		return 1
	}

	return exchangeRate
}

func isValidInitialCurrency(initialCurrency string) bool {
	return initialCurrency == USD || initialCurrency == EUR || initialCurrency == RUB
}

func isValidAmount(amount float64) bool {
	return amount > 0
}

func isValidTargetCurrency(targetCurrency string, initialCurrency string) bool {
	return isValidInitialCurrency(targetCurrency) && initialCurrency != targetCurrency
}
