package main

import (
	"fmt"
	"strings"
)

// Создайте переменную calculator, которая является функцией, принимающей два числа и возвращающей их сумму, разность, произведение и частное в виде строки.
var calculator func(int, int) string = func(a, b int) string {
	sum := a + b
	diff := a - b
	product := a * b
	var quotient string
	if b != 0 {
		quotient = fmt.Sprintf("%.2f", float64(a)/float64(b))
	} else {
		quotient = "Division by zero"
	}
	return fmt.Sprintf("Sum: %d, Difference: %d, Product: %d, Quotient: %s", sum, diff, product, quotient)
}

// Создайте переменную stringManipulator, которая является функцией, принимающей строку и возвращающей эту строку, в которой каждое слово начинается с заглавной буквы.
var stringManipulator func(string) string = func(input string) string {
	words := strings.Fields(input)
	for i, word := range words {
		words[i] = strings.Title(word)
	}
	return strings.Join(words, " ")
}

// Создайте переменную filter, которая является функцией, принимающей срез чисел и функцию-условие, возвращающую новый срез, содержащий только те числа, для которых условие истинно.
type condition func(int) bool

var filter func([]int, condition) []int = func(numbers []int, cond condition) []int {
	var filtered []int
	for _, num := range numbers {
		if cond(num) {
			filtered = append(filtered, num)
		}
	}
	return filtered
}
