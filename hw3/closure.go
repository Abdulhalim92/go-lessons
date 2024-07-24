package main

import "fmt"

// Multiplier
// Создайте функцию Multiplier, которая принимает целое число factor и возвращает функцию, умножающую переданное ей число на factor.
func Multiplier(factor int) func(int) int {
	return func(num int) int {
		return num * factor
	}
}

// StringRepeater
// Создайте функцию StringRepeater, которая принимает целое число n и возвращает функцию, повторяющую переданную ей строку n раз.
func StringRepeater(n int) func(string) string {
	return func(str string) string {
		result := ""
		for i := 0; i < n; i++ {
			result += str
		}
		return result
	}
}

// ConditionalPrinter
// Создайте функцию ConditionalPrinter, которая принимает булево значение condition и возвращает функцию, печатающую строку, если condition истинно.
func ConditionalPrinter(condition bool) func(string) {
	return func(message string) {
		if condition {
			fmt.Println(message)
		}
	}
}
