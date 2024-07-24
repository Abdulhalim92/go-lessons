package main

import "fmt"

// FindMax
// Создайте функцию FindMax, которая принимает срез чисел и возвращает наибольшее из них.
func FindMax(numbers []int) int {
	if len(numbers) == 0 {
		return 0 // or handle error case
	}
	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	return max
}

// FormatName
// Создайте функцию FormatName, которая принимает имя и фамилию и возвращает их в формате "Фамилия, Имя".
func FormatName(firstName, lastName string) string {
	return fmt.Sprintf("%s, %s", lastName, firstName)
}

// FilterEvenNumbers
// Создайте функцию FilterEvenNumbers, которая принимает срез чисел и возвращает новый срез, содержащий только четные числа.
func FilterEvenNumbers(numbers []int) []int {
	var evens []int
	for _, num := range numbers {
		if num%2 == 0 {
			evens = append(evens, num)
		}
	}
	return evens
}
