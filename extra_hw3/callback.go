package main

import "fmt"

// ProcessNumbers
// Создайте функцию ProcessNumbers, которая принимает срез чисел и функцию для их обработки. Примените переданную функцию к этим числам и выведите результат.
func ProcessNumbers(numbers []int, processor func(int) int) {
	for _, num := range numbers {
		result := processor(num)
		fmt.Println(result)
	}
}

// ApplyToStrings
// Создайте функцию ApplyToStrings, которая принимает срез строк и функцию, преобразующую строку, и возвращает новый срез с преобразованными строками.
func ApplyToStrings(strings []string, transformer func(string) string) []string {
	var transformed []string
	for _, str := range strings {
		transformed = append(transformed, transformer(str))
	}
	return transformed
}

// Создайте функцию FilterAndPrint, которая принимает срез чисел и функцию-условие, печатающую только те числа, для которых условие истинно.
type cond func(int) bool

func FilterAndPrint(numbers []int, cond cond) {
	for _, num := range numbers {
		if cond(num) {
			fmt.Println(num)
		}
	}
}
