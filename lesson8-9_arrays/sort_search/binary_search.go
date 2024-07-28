package main

func binarySearch(slice []int, target int) int {
	low, high := 0, len(slice)-1
	for low <= high {
		mid := low + (high-low)/2
		if slice[mid] == target {
			return mid
		}
		if slice[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1 // элемент не найден
}
