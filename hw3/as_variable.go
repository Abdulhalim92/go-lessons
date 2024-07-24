package main

// Создайте переменную adder, которая является функцией, принимающей два целых числа и возвращающей их сумму.
var adder func(int, int) int = func(a, b int) int {
	return a + b
}

// Создайте переменную concatenator, которая является функцией, принимающей две строки и возвращающей их конкатенацию.
var concatenator func(string, string) string = func(str1, str2 string) string {
	return str1 + str2
}

// Создайте переменную isPositive, которая является функцией, принимающей целое число и возвращающей true, если число положительное.
var isPositive func(int) bool = func(num int) bool {
	return num > 0
}
