package main

func quickSort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}
	pivot := slice[len(slice)/2]
	left := []int{}
	right := []int{}
	for _, v := range slice {
		if v < pivot {
			left = append(left, v)
		} else if v > pivot {
			right = append(right, v)
		}
	}
	return append(append(quickSort(left), pivot), quickSort(right)...)
}
