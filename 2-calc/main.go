package main

import (
	"fmt"
	"strconv"
	"strings"
)

const AVG, SUM, MED = "AVG", "SUM", "MED"

func main() {
	operation, integersStr := getUserInput()
	integers := splitInput(integersStr)

	fmt.Println(operation, integersStr, integers)
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

func splitInput(str string) []string {
	return strings.Split(str, ", ")
}
