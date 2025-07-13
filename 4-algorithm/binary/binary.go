package binary

import (
	"errors"
	"strconv"
)

func BinarySearchInt(integers []int, target int) (int, error) {
	var midIdx int
	length := len(integers)
	msg := "The value " + strconv.Itoa(target) + " is not in the slice"
	if length == 0 {
		return -1, errors.New(msg)
	}

	firstIdx, lastIdx := 0, length-1
	leftBorder := firstIdx
	rightBorder := lastIdx

	for {
		if lastIdx < leftBorder || firstIdx > rightBorder {
			break
		}
		midIdx = (firstIdx + lastIdx) / 2

		value := integers[midIdx]

		if firstIdx == lastIdx {
			value = integers[firstIdx]
			if value == target {
				return midIdx, nil
			} else {
				break
			}
		}

		if value == target {
			return midIdx, nil
		}
		if value > target {
			lastIdx = midIdx - 1
		}
		if value < target {
			firstIdx = midIdx + 1
		}
	}

	return -1, errors.New(msg)
}
