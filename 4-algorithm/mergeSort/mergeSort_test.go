package mergeSort_test

import (
	"algorithm/mergeSort"
	"testing"
)

var testCases = []struct {
	name string
	arr  []int
}{
	{"Positive integers", []int{10, 8, 12, 3, 1, 5}},
	{"Negative integers", []int{-10, -8, -12, -3, -1, -5}},
	{"Positive and negative integers", []int{10, -8, 12, -3, 1, -5}},
}

func TestMergeSortIntSliceASC(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			lastIdx := len(tc.arr) - 1
			result := mergeSort.MergeSortIntSliceASC(tc.arr)
			for id, elem := range result {
				if id == lastIdx {
					break
				}
				if elem > result[id+1] {
					t.Errorf("Wrong sort element %d with id %d is more than the next one (%d)", elem, id, result[id+1])
				}
			}
		})
	}
}
