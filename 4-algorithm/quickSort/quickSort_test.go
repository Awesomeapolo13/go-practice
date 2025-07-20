package quickSort_test

import (
	"algorithm/quickSort"
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

func TestQuickSortIntSliceASC(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			lastIdx := len(tc.arr) - 1
			arr := quickSort.QuickSortIntSliceASC(tc.arr, 0, lastIdx)
			for id, elem := range arr {
				if id == lastIdx {
					break
				}
				if elem > arr[id+1] {
					t.Errorf("Wrong sort element %d with id %d is more than the next one (%d)", elem, id, arr[id+1])
				}
			}
		})
	}
}
