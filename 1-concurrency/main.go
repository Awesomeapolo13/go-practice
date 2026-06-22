package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	intCount := 10
	resultChannel := make(chan int, intCount)
	for i := 0; i < intCount; i++ {
		go createTenRandomIntegers(resultChannel)
	}

	for i := 0; i < intCount; i++ {
		fmt.Println(<-resultChannel)
	}
}

func createTenRandomIntegers(resultChan chan int) {
	randomNum := rand.IntN(100)
	go pow2(randomNum, resultChan)
}

func pow2(number int, resultChan chan int) {
	resultChan <- number * number
}
