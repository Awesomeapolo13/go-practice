package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const AVG, SUM, MED = "AVG", "SUM", "MED"

func main() {
	var integers *[]float64
	var result float64
	operation, integersStr := getUserInput()
	splitted := splitInput(integersStr)
	integers = &splitted

	result = calculate(operation, integers)

	msg := fmt.Sprintf("The result of %s operation with integers %.2f is %.2f", operation, *integers, result)
	fmt.Println(msg)
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

func splitInput(str string) []float64 {
	split := strings.Split(str, ", ")
	integers := make([]float64, len(split))
	for idx, val := range split {
		converted, err := strconv.ParseFloat(val, 64)
		if err != nil {
			fmt.Println("Error while converting ", val, " to float. Try again.")
			panic(err)
		}

		integers[idx] = converted
	}

	return integers
}

func calculate(operation string, integers *[]float64) float64 {
	switch operation {
	case AVG:
		return calcAVG(integers)
	case SUM:
		return calcSUM(integers)
	case MED:
		return calcMED(integers)
	default:
		panic("Unsupported operation")
	}
}

func calcSUM(nums *[]float64) float64 {
	sum := 0.0
	for _, num := range *nums {
		sum += num
	}

	return sum
}

func calcAVG(nums *[]float64) float64 {
	elemCount := len(*nums)
	sum := calcSUM(nums)

	return sum / float64(elemCount)
}

func calcMED(nums *[]float64) float64 {
	sort.Float64s(*nums)
	fmt.Println(*nums)
	length := len(*nums)
	halfIdx := length / 2

	if length%2 != 0 {
		return (*nums)[halfIdx]
	}

	return ((*nums)[halfIdx] + (*nums)[halfIdx-1]) / 2
}
