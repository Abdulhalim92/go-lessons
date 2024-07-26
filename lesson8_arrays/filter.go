package main

import "fmt"

// Фильтрация среза
func filter(slice []int, fn func(int) bool) []int {
	var result []int
	for _, v := range slice {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// Применение функции к каждому элементу среза
func mapSlice(slice []int, fn func(int) int) []int {
	result := make([]int, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

func op() {
	slice := []int{1, 2, 3, 4, 5}

	// Фильтрация: оставить только четные числа
	even := filter(slice, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println(even) // [2 4]

	// Маппинг: возвести каждый элемент в квадрат
	squared := mapSlice(slice, func(n int) int {
		return n * n
	})
	fmt.Println(squared) // [1 4 9 16 25]
}
