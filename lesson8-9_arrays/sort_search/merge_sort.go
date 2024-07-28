package main

func merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

func mergeSort(slice []int) []int {
	if len(slice) <= 1 {
		return slice
	}
	mid := len(slice) / 2
	left := mergeSort(slice[:mid])
	right := mergeSort(slice[mid:])
	return merge(left, right)
}
