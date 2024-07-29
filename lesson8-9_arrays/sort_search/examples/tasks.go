package main

// Сортировка пузырьком (Bubble Sort)

// Отсортировать срез целых чисел от 1 до 5 в порядке убывания.
func bubbleSort(slice []int) {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if slice[j] < slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

// Отсортировать срез строк, содержащих имена животных, в порядке возрастания.
func bubbleSortStrings(slice []string) {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

// Отсортировать срез вещественных чисел от 0.1 до 0.5 в порядке убывания.
func bubbleSortFloats(slice []float64) {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if slice[j] < slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

// Реализовать пузырьковую сортировку для среза структур Person, отсортировать по возрасту в порядке убывания.

type Person struct {
	Name string
	Age  int
}

func bubbleSortPersons(slice []Person) {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if slice[j].Age < slice[j+1].Age {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

// Сортировать срез строк по длине строки в порядке возрастания с использованием пузырьковой сортировки.
func bubbleSortStringsByLength(slice []string) {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if len(slice[j]) > len(slice[j+1]) {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

// Отсортировать срез целых чисел с использованием пузырьковой сортировки и определить, если количество
// перестановок элементов превышает заданный порог.
func bubbleSortWithCount(slice []int, threshold int) (sorted []int, swaps int) {
	n := len(slice)
	swaps = 0
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
				swaps++
				if swaps > threshold {
					return slice, swaps
				}
			}
		}
	}
	return slice, swaps
}

// Сортировка вставками (Insertion Sort)

// Отсортировать срез целых чисел от 1 до 5 в порядке возрастания.
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

// Отсортировать срез строк в алфавитном порядке.
func insertionSortStrings(slice []string) {
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

// Отсортировать срез вещественных чисел от 0.5 до 0.1 в порядке возрастания.
func insertionSortFloats(slice []float64) {
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

// Реализовать сортировку вставками для среза структур Book, отсортировать по году издания в порядке возрастания.

type Book struct {
	Title string
	Year  int
}

func insertionSortBooks(slice []Book) {
	for i := 1; i < len(slice); i++ {
		key := slice[i]
		j := i - 1
		for j >= 0 && slice[j].Year > key.Year {
			slice[j+1] = slice[j]
			j--
		}
		slice[j+1] = key
	}
}

// Сортировать срез строк по длине строки в порядке убывания с использованием сортировки вставками.
func insertionSortStringsByLength(slice []string) {
	for i := 1; i < len(slice); i++ {
		key := slice[i]
		j := i - 1
		for j >= 0 && len(slice[j]) < len(key) {
			slice[j+1] = slice[j]
			j--
		}
		slice[j+1] = key
	}
}

// Сортировать срез целых чисел с использованием сортировки вставками и определить количество сравнений,
// сделанных при сортировке.
func insertionSortWithCount(slice []int) (sorted []int, comparisons int) {
	comparisons = 0
	for i := 1; i < len(slice); i++ {
		key := slice[i]
		j := i - 1
		for j >= 0 {
			comparisons++
			if slice[j] > key {
				slice[j+1] = slice[j]
				j--
			} else {
				break
			}
		}
		slice[j+1] = key
	}
	return slice, comparisons
}

// Сортировка выбором (Selection Sort)

// Отсортировать срез целых чисел от 1 до 4 в порядке возрастания.
func selectionSort(slice []int) {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if slice[j] < slice[minIdx] {
				minIdx = j
			}
		}
		slice[i], slice[minIdx] = slice[minIdx], slice[i]
	}
}

// Отсортировать срез строк в алфавитном порядке.
func selectionSortStrings(slice []string) {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if slice[j] < slice[minIdx] {
				minIdx = j
			}
		}
		slice[i], slice[minIdx] = slice[minIdx], slice[i]
	}
}

// Отсортировать срез вещественных чисел от 0.1 до 0.5 в порядке убывания.
func selectionSortFloats(slice []float64) {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if slice[j] > slice[maxIdx] {
				maxIdx = j
			}
		}
		slice[i], slice[maxIdx] = slice[maxIdx], slice[i]
	}
}

// Реализовать сортировку выбором для среза структур Employee, отсортировать по
// зарплате в порядке возрастания.

type Employee struct {
	Name   string
	Salary int
}

func selectionSortEmployees(slice []Employee) {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if slice[j].Salary < slice[minIdx].Salary {
				minIdx = j
			}
		}
		slice[i], slice[minIdx] = slice[minIdx], slice[i]
	}
}

// Сортировать срез строк по длине строки в порядке возрастания с использованием сортировки выбором.
func selectionSortStringsByLength(slice []string) {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if len(slice[j]) < len(slice[minIdx]) {
				minIdx = j
			}
		}
		slice[i], slice[minIdx] = slice[minIdx], slice[i]
	}
}

// Сортировать срез целых чисел с использованием сортировки выбором и определить количество обменов,
// сделанных при сортировке.
func selectionSortWithCount(slice []int) (sorted []int, swaps int) {
	n := len(slice)
	swaps = 0
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if slice[j] < slice[minIdx] {
				minIdx = j
			}
		}
		if minIdx != i {
			slice[i], slice[minIdx] = slice[minIdx], slice[i]
			swaps++
		}
	}
	return slice, swaps
}

// Сортировка слиянием (Merge Sort)

// Отсортировать срез целых чисел от 1 до 5 в порядке возрастания с использованием сортировки слиянием.
func mergeSort(slice []int) []int {
	if len(slice) <= 1 {
		return slice
	}
	mid := len(slice) / 2
	left := mergeSort(slice[:mid])
	right := mergeSort(slice[mid:])
	return mergeSlices(left, right)
}

func mergeSlices(left, right []int) []int {
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

// Отсортировать срез строк в алфавитном порядке с использованием сортировки слиянием.
func mergeSortStrings(slice []string) []string {
	if len(slice) <= 1 {
		return slice
	}
	mid := len(slice) / 2
	left := mergeSortStrings(slice[:mid])
	right := mergeSortStrings(slice[mid:])
	return mergeStrings(left, right)
}

func mergeStrings(left, right []string) []string {
	result := []string{}
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

// Отсортировать срез вещественных чисел от 0.5 до 0.1 в порядке убывания с использованием
// сортировки слиянием.
func mergeSortFloats(slice []float64) []float64 {
	if len(slice) <= 1 {
		return slice
	}
	mid := len(slice) / 2
	left := mergeSortFloats(slice[:mid])
	right := mergeSortFloats(slice[mid:])
	return mergeFloats(left, right)
}

func mergeFloats(left, right []float64) []float64 {
	result := []float64{}
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] > right[j] {
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

// Реализовать сортировку слиянием для среза структур Student, отсортировать по среднему баллу
// в порядке возрастания.

type Student struct {
	Name  string
	Grade float64
}

func mergeSortStudents(slice []Student) []Student {
	if len(slice) <= 1 {
		return slice
	}
	mid := len(slice) / 2
	left := mergeSortStudents(slice[:mid])
	right := mergeSortStudents(slice[mid:])
	return mergeStudents(left, right)
}

func mergeStudents(left, right []Student) []Student {
	result := []Student{}
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i].Grade < right[j].Grade {
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

// Сортировать срез строк по длине строки в порядке убывания с использованием сортировки слиянием.
func mergeSortStringsByLength(slice []string) []string {
	if len(slice) <= 1 {
		return slice
	}
	mid := len(slice) / 2
	left := mergeSortStringsByLength(slice[:mid])
	right := mergeSortStringsByLength(slice[mid:])
	return mergeStringsByLength(left, right)
}

func mergeStringsByLength(left, right []string) []string {
	result := []string{}
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if len(left[i]) > len(right[j]) {
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

// Сортировать срез целых чисел с использованием сортировки слиянием и определить количество слияний,
// сделанных при сортировке.
func mergeSortWithMergeCount(slice []int) (sorted []int, mergeCount int) {
	mergeCount = 0
	sorted = mergeSortCount(slice, &mergeCount)
	return sorted, mergeCount
}

func mergeSortCount(slice []int, mergeCount *int) []int {
	if len(slice) <= 1 {
		return slice
	}
	mid := len(slice) / 2
	left := mergeSortCount(slice[:mid], mergeCount)
	right := mergeSortCount(slice[mid:], mergeCount)
	*mergeCount++
	return mergeCountSlices(left, right)
}

func mergeCountSlices(left, right []int) []int {
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
