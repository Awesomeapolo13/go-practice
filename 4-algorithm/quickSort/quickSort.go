package quickSort

func QuickSortIntSliceASC(arr []int, firstIdx, lastIdx int) []int {
	if firstIdx >= lastIdx {
		return arr
	}

	pivotIdx := makePartition(arr, firstIdx, lastIdx)
	QuickSortIntSliceASC(arr, firstIdx, pivotIdx-1)
	QuickSortIntSliceASC(arr, pivotIdx+1, lastIdx)

	return arr
}

func makePartition(arr []int, firstIdx, lastIdx int) int {
	pivot := arr[lastIdx]
	i := firstIdx - 1

	for j := firstIdx; j < lastIdx; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[lastIdx] = arr[lastIdx], arr[i+1]

	return i + 1
}
