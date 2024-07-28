package main

func insertionSort(slice []int) {
	for i := 1; i < len(slice); i++ {
		key := slice[i]
		j := i - 1
		for j >= 0 && slice[j] > key {
			slice[j+1] = slice[j]
			j--
		}
		slice[j+1] = key
	}
}
