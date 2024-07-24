package main

import "fmt"

// PrintMultiplicationTable
// Создайте функцию PrintMultiplicationTable, которая принимает целое число и печатает таблицу умножения для этого числа до 10.
func PrintMultiplicationTable(number int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d * %d = %d\n", number, i, number*i)
	}
}

// PrintFilteredStrings
// Создайте функцию PrintFilteredStrings, которая принимает срез строк и функцию-условие, печатающую только те строки, для которых условие истинно.
func PrintFilteredStrings(strings []string, condition func(string) bool) {
	for _, str := range strings {
		if condition(str) {
			fmt.Println(str)
		}
	}
}

// CalculateStatistics
// Создайте функцию CalculateStatistics, которая принимает срез чисел и печатает их среднее значение и сумму.
func CalculateStatistics(numbers []int) {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	average := float64(sum) / float64(len(numbers))
	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Average: %.2f\n", average)
}
