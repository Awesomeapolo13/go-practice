package main

import (
	"fmt"
	"strconv"
)

const AVG, SUM, MED = "AVG", "SUM", "MED"

func main() {
	// Принимаем строку из чисел через запятую ("2,3,10,-6"), их нужно разбить по запятой, удалить лищние пробелы
	// В итоге получаем слайс с числами, которые были переданы через запятую
	operation, integers := getUserInput()
	fmt.Println(operation, integers)
}

func getUserInput() (string, string) {
	operation := getOperationInput()
	integers := getIntegersInput()
	msg := fmt.Sprintf("You want to make %s operation with integers %s", operation, integers)
	fmt.Println(msg)

	return operation, integers
}

func getOperationInput() string {
	var operation string

	for {
		fmt.Print("Enter operation (AVG|SUM|MED): ")
		fmt.Scan(&operation)

		if operation == AVG || operation == SUM || operation == MED {
			break
		}
		fmt.Println("Wrong operation input. Try again.")
	}

	return operation
}

func getIntegersInput() string {
	var integers string
	var integer string

	for {
		fmt.Print("Enter integers (to finish input, enter not integer symbol, empty input is not allowed): ")
		fmt.Scan(&integer)
		_, err := strconv.Atoi(integer)

		if err != nil && integers != "" {
			break
		} else if err != nil && integers == "" {
			fmt.Println("Empty input is not allowed. Try again.")
			continue
		}

		if integers == "" {
			integers = integer
		} else {
			integers += ", " + integer
		}
	}

	return integers
}
