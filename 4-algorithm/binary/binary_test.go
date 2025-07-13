package binary_test

import (
	"algorithm/binary"
	"errors"
	"strconv"
	"testing"
)

var evenIntegersCount = []int{1, 2, 3, 4, 5, 6}
var oddIntegersCount = []int{1, 2, 3, 4, 5, 6, 7}
var onlyOneElement = []int{2}

var nonEmptySearchCases = []struct {
	name       string
	integers   []int
	target     int
	expectedId int
}{
	{name: "Even count, find the first value", integers: evenIntegersCount, target: 1, expectedId: 0},
	{name: "Even count, find in first half part of values", integers: evenIntegersCount, target: 2, expectedId: 1},
	{name: "Even count, find in last half part of values", integers: evenIntegersCount, target: 5, expectedId: 4},
	{name: "Even count, find the last value", integers: evenIntegersCount, target: 6, expectedId: 5},
	{name: "Odd count, find the first value", integers: oddIntegersCount, target: 1, expectedId: 0},
	{name: "Odd count, find in first half part of values", integers: oddIntegersCount, target: 2, expectedId: 1},
	{name: "Odd count, find in the middle part of values", integers: oddIntegersCount, target: 4, expectedId: 3},
	{name: "Odd count, find in last half part of values", integers: oddIntegersCount, target: 5, expectedId: 4},
	{name: "Odd count, find the last value", integers: oddIntegersCount, target: 7, expectedId: 6},
	{name: "Only one element", integers: onlyOneElement, target: 2, expectedId: 0},
}

var emptySearchCases = []struct {
	name     string
	integers []int
	target   int
}{
	{name: "Even count, the target is unexists in slice more than zero", integers: evenIntegersCount, target: 20},
	{name: "Even count, the target is unexists in slice less than zero", integers: evenIntegersCount, target: -20},
	{name: "Odd count, the target is unexists in slice more than zero", integers: oddIntegersCount, target: 20},
	{name: "Odd count, the target is unexists in slice less than zero", integers: oddIntegersCount, target: -20},
	{name: "Only one element", integers: onlyOneElement, target: 5},
	{name: "Empty slice", integers: []int{}, target: 2},
}

func TestBinarySearchIntWIthNonEmptyResult(t *testing.T) {
	for _, tc := range nonEmptySearchCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := binary.BinarySearchInt(tc.integers, tc.target)

			if err != nil {
				t.Errorf("Expected %d integer in %s test case, got an error %v", tc.target, tc.name, err)
			}

			if result != tc.expectedId {
				t.Errorf("Expected %d integer as a result in %s test case, got %d", tc.target, tc.name, result)
			}
		})
	}
}

func TestBinarySearchIntWithEmptyResult(t *testing.T) {
	for _, tc := range emptySearchCases {
		t.Run(tc.name, func(t *testing.T) {
			expectedErr := errors.New("The value " + strconv.Itoa(tc.target) + " is not in the slice")
			_, err := binary.BinarySearchInt(tc.integers, tc.target)

			if errors.Is(err, expectedErr) {
				t.Errorf("Expected an error %v, got %v", expectedErr, err)
			}
		})
	}
}
