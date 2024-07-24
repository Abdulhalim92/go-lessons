package main

import "fmt"

// CreateMultiplier
// Создайте функцию CreateMultiplier, которая принимает целое число factor и возвращает функцию, умножающую переданное ей число на factor.
func CreateMultiplier(factor int) func(int) int {
	return func(num int) int {
		return num * factor
	}
}

// CreateFormatter
// Создайте функцию CreateFormatter, которая принимает формат строки и возвращает функцию, форматирующую переданную строку в соответствии с заданным форматом.
func CreateFormatter(format string) func(string) string {
	return func(str string) string {
		return fmt.Sprintf(format, str)
	}
}

// CreateCalculator
// Создайте функцию CreateCalculator, которая возвращает функции для выполнения операций сложения и вычитания с переданными операндами.
func CreateCalculator() (func(int, int) int, func(int, int) int) {
	add := func(a, b int) int {
		return a + b
	}
	subtract := func(a, b int) int {
		return a - b
	}
	return add, subtract
}
