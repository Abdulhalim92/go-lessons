package main

// Calculate
// Создайте функцию Calculate, которая принимает два целых числа и функцию для их обработки. Примените переданную функцию к этим числам и верните результат.
func Calculate(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

// Execute
// Создайте функцию Execute, которая принимает булевое значение и функцию, которая принимает булевое значение и ничего не возвращает.
// Выполните переданную функцию с переданным булевым значением.
func Execute(state bool, action func(bool)) {
	action(state)
}

// Apply
// Создайте функцию Apply, которая принимает целое число и функцию, которая принимает целое число и возвращает целое число.
// Примените переданную функцию к переданному числу и верните результат.
func Apply(num int, transformer func(int) int) int {
	return transformer(num)
}
