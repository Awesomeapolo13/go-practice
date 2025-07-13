package binary

import (
	"errors"
	"strconv"
)

func BinarySearchInt(integers []int, target int) (int, error) {
	length := len(integers)
	msg := "The value " + strconv.Itoa(target) + " is not in the slice"
	if length == 0 {
		return -1, errors.New(msg)
	}

	firstIdx, lastIdx := 0, length-1

	for firstIdx <= lastIdx {
		midIdx := firstIdx + (lastIdx-firstIdx)/2

		if integers[midIdx] == target {
			return midIdx, nil
		}
		if integers[midIdx] > target {
			lastIdx = midIdx - 1
		} else {
			firstIdx = midIdx + 1
		}
	}

	return -1, errors.New(msg)
}
