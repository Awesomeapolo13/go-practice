package mergeSort

func MergeSortIntSliceASC(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}

	midIdx := length / 2
	leftArr, rightArr := arr[:midIdx], arr[midIdx:]

	return merge(
		MergeSortIntSliceASC(leftArr),
		MergeSortIntSliceASC(rightArr),
	)
}

func merge(leftArr []int, rightArr []int) []int {
	leftLen, rightLen := len(leftArr), len(rightArr)
	result := make([]int, 0, leftLen+rightLen)
	i, j := 0, 0

	for i < leftLen && j < rightLen {
		if leftArr[i] <= rightArr[j] {
			result = append(result, leftArr[i])
			i++
		} else {
			result = append(result, rightArr[j])
			j++
		}
	}
	// Если левый массив был больше, то в нем могли остаться ещё элементы. Добавляем их все в конец
	for i < leftLen {
		result = append(result, leftArr[i])
		i++
	}
	// Аналогично для правого
	for j < rightLen {
		result = append(result, rightArr[j])
		j++
	}

	return result
}
