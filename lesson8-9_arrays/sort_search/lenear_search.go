package main

func linearSearch(slice []int, target int) int {
	for i, v := range slice {
		if v == target {
			return i
		}
	}
	return -1 // элемент не найден
}
