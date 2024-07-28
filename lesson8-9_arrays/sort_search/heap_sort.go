package main

func heapify(slice []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2
	if left < n && slice[left] > slice[largest] {
		largest = left
	}
	if right < n && slice[right] > slice[largest] {
		largest = right
	}
	if largest != i {
		slice[i], slice[largest] = slice[largest], slice[i]
		heapify(slice, n, largest)
	}
}

func heapSort(slice []int) {
	n := len(slice)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(slice, n, i)
	}
	for i := n - 1; i >= 0; i-- {
		slice[0], slice[i] = slice[i], slice[0]
		heapify(slice, i, 0)
	}
}
