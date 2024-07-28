package main

import "fmt"

func main() {
	slice := []int{64, 34, 25, 12, 22, 11, 90}
	bubbleSort(slice)
	fmt.Println("Bubble Sort:", slice) // [11 12 22 25 34 64 90]

	slice = []int{64, 34, 25, 12, 22, 11, 90}
	insertionSort(slice)
	fmt.Println("Insertion Sort:", slice) // [11 12 22 25 34 64 90]

	slice = []int{64, 34, 25, 12, 22, 11, 90}
	selectionSort(slice)
	fmt.Println("Selection Sort:", slice) // [11 12 22 25 34 64 90]

	slice = []int{64, 34, 25, 12, 22, 11, 90}
	sortedSlice := mergeSort(slice)
	fmt.Println("Merge Sort:", sortedSlice) // [11 12 22 25 34 64 90]

	slice = []int{64, 34, 25, 12, 22, 11, 90}
	sortedSlice = quickSort(slice)
	fmt.Println("Quick Sort:", sortedSlice) // [11 12 22 25 34 64 90]

	slice = []int{64, 34, 25, 12, 22, 11, 90}
	heapSort(slice)
	fmt.Println("Heap Sort:", slice) // [11 12 22 25 34 64 90]

	slice = []int{10, 20, 30, 40, 50}
	index := linearSearch(slice, 30)
	fmt.Println("Index of 30:", index) // Index of 30: 2

	slice = []int{10, 20, 30, 40, 50}
	index = binarySearch(slice, 30)
	fmt.Println("Index of 30:", index) // Index of 30: 2
}
