package main

import "fmt"

func main() {
	bubble()
	insert()
	selection()
	merge()
}

func bubble() {
	slice := []int{1, 2, 3, 4, 5}
	bubbleSort(slice)
	fmt.Println(slice) // [5 4 3 2 1]

	slice2 := []string{"dog", "cat", "elephant"}
	bubbleSortStrings(slice2)
	fmt.Println(slice2) // [cat dog elephant]

	slice3 := []float64{0.1, 0.4, 0.3, 0.5}
	bubbleSortFloats(slice3)
	fmt.Println(slice3) // [0.5 0.4 0.3 0.1]

	slice4 := []Person{{"Alice", 30}, {"Bob", 25}, {"Charlie", 35}}
	bubbleSortPersons(slice4)
	fmt.Println(slice4) // [{Charlie 35} {Alice 30} {Bob 25}]

	slice5 := []string{"short", "longer", "longest"}
	bubbleSortStringsByLength(slice5)
	fmt.Println(slice5) // [short longer longest]

	slice6 := []int{5, 4, 3, 2, 1}
	sortedSlice, count := bubbleSortWithCount(slice6, 5)
	fmt.Println(sortedSlice, count) // [1 2 3 4 5] 10
}

func insert() {
	slice := []int{5, 4, 3, 2, 1}
	insertionSort(slice)
	fmt.Println(slice) // [1 2 3 4 5]

	slice2 := []string{"banana", "apple", "cherry"}
	insertionSortStrings(slice2)
	fmt.Println(slice2) // [apple banana cherry]

	slice3 := []float64{0.5, 0.4, 0.3, 0.2, 0.1}
	insertionSortFloats(slice3)
	fmt.Println(slice3) // [0.1 0.2 0.3 0.4 0.5]

	slice4 := []Book{{"Book A", 2020}, {"Book B", 2015}, {"Book C", 2018}}
	insertionSortBooks(slice4)
	fmt.Println(slice4) // [{Book B 2015} {Book C 2018} {Book A 2020}]

	slice5 := []string{"short", "longer", "longest"}
	insertionSortStringsByLength(slice5)
	fmt.Println(slice5) // [longest longer short]

	slice6 := []int{4, 3, 2, 1}
	sortedSlice, count := insertionSortWithCount(slice6)
	fmt.Println(sortedSlice, count) // [1 2 3 4] 6
}

func selection() {
	slice := []int{4, 3, 2, 1}
	selectionSort(slice)
	fmt.Println(slice) // [1 2 3 4]

	slice2 := []string{"banana", "apple", "cherry"}
	selectionSortStrings(slice2)
	fmt.Println(slice2) // [apple banana cherry]

	slice3 := []float64{0.1, 0.3, 0.2, 0.5}
	selectionSortFloats(slice3)
	fmt.Println(slice3) // [0.5 0.3 0.2 0.1]

	slice4 := []Employee{{"Alice", 50000}, {"Bob", 40000}, {"Charlie", 60000}}
	selectionSortEmployees(slice4)
	fmt.Println(slice4) // [{Bob 40000} {Alice 50000} {Charlie 60000}]

	slice5 := []string{"longest", "longer", "short"}
	selectionSortStringsByLength(slice5)
	fmt.Println(slice5) // [short longer longest]

	slice6 := []int{3, 1, 4, 1, 5}
	sortedSlice, count := selectionSortWithCount(slice6)
	fmt.Println(sortedSlice, count) // [1 1 3 4 5] 4
}

func merge() {
	slice := []int{4, 2, 5, 1, 3}
	sortedSlice := mergeSort(slice)
	fmt.Println(sortedSlice) // [1 2 3 4 5]

	slice2 := []string{"banana", "apple", "cherry"}
	sortedSlice2 := mergeSortStrings(slice2)
	fmt.Println(sortedSlice2) // [apple banana cherry]

	slice3 := []float64{0.5, 0.2, 0.4, 0.1, 0.3}
	sortedSlice3 := mergeSortFloats(slice3)
	fmt.Println(sortedSlice3) // [0.5 0.4 0.3 0.2 0.1]

	slice4 := []Student{{"Alice", 90.0}, {"Bob", 85.0}, {"Charlie", 95.0}}
	sortedSlice4 := mergeSortStudents(slice4)
	fmt.Println(sortedSlice4) // [{Bob 85} {Alice 90} {Charlie 95}]

	slice5 := []string{"short", "longer", "longest"}
	sortedSlice5 := mergeSortStringsByLength(slice5)
	fmt.Println(sortedSlice5) // [longest longer short]

	slice6 := []int{4, 2, 5, 1, 3}
	sortedSlice6, count := mergeSortWithMergeCount(slice6)
	fmt.Println(sortedSlice6, count) // [1 2 3 4 5] 2
}
