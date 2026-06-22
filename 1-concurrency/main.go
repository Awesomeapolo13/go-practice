package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	intCount := 10
	resultChannel := make(chan int, intCount)
	randomIntegersChan := make(chan int, intCount)
	go createTenRandomIntegers(resultChannel)
	go pow2(randomIntegersChan, resultChannel)

	for i := 0; i < intCount; i++ {
		fmt.Println(<-resultChannel)
	}
}

func createTenRandomIntegers(randomIntegersChan chan int) {
	defer close(randomIntegersChan)
	elementsCount := 10
	randomIntegersSlice := make([]int, elementsCount)
	for i := 0; i < elementsCount; i++ {
		randomIntegersSlice[i] = rand.IntN(100)
	}

	for _, randomInteger := range randomIntegersSlice {
		randomIntegersChan <- randomInteger
	}
}

func pow2(randomIntegersChan, resultChan chan int) {
	defer close(resultChan)
	for number := range randomIntegersChan {
		resultChan <- number * number
	}
}
