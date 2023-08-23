package sort

func MergeSort(unsortedArray []int) []int {
	if len(unsortedArray) < 2 {
		return unsortedArray
	}
	middle := len(unsortedArray) / 2
	return merge(
		MergeSort(unsortedArray[:middle]),
		MergeSort(unsortedArray[middle:]),
	)
}

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	leftIndex, rightIndex := 0, 0
	lenLeft := len(left)
	lenRight := len(right)

	for leftIndex < lenLeft && rightIndex < lenRight {
		if left[leftIndex] <= right[rightIndex] {
			result[leftIndex+rightIndex] = left[leftIndex]
			leftIndex++
		} else {
			result[leftIndex+rightIndex] = right[rightIndex]
			rightIndex++
		}
	}
	for leftIndex < lenLeft {
		result[leftIndex+rightIndex] = left[leftIndex]
		leftIndex++
	}
	for rightIndex < lenRight {
		result[leftIndex+rightIndex] = right[rightIndex]
		rightIndex++
	}
	return result
}
