package main

// Возведение в квадрат
// Определите тип функции UnaryOperation, которая принимает int и возвращает int. Напишите функцию
// для возведения числа в квадрат и используйте тип UnaryOperation для вызова этой функции.

type UnaryOperation func(int) int

// Функция для возведения в квадрат
func square(a int) int {
	return a * a
}

// Удвоение числа
// Определите тип функции UnaryOperation, которая принимает int и возвращает int. Напишите
// функцию для удвоения числа и используйте тип UnaryOperation для вызова этой функции.

// Функция для удвоения числа
func double(a int) int {
	return a * 2
}

// Проверка четности
// Определите тип функции Check, которая принимает int и возвращает bool. Напишите функцию
// для проверки четности числа и используйте тип Check для вызова этой функции.

type Check func(int) bool

// Функция для проверки четности числа
func isEven(num int) bool {
	return num%2 == 0
}

// Проверка положительности
// Определите тип функции Check, которая принимает int и возвращает bool. Напишите функцию
// для проверки, является ли число положительным, и используйте тип Check для вызова этой функции.

// Функция для проверки положительности числа
func isPositive(num int) bool {
	return num > 0
}

// Комбинирование функций
// Определите тип функции Operation, которая принимает два int и возвращает int. Напишите функции
// для сложения, вычитания и умножения. Напишите функцию, которая принимает два int и массив
// функций Operation, и возвращает массив результатов применения этих функций.

type Operation func(int, int) int

// Функции для арифметических операций
func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

func multiply(a int, b int) int {
	return a * b
}

// Функция для применения операций
func applyOperations(a int, b int, ops ...Operation) []int {
	results := make([]int, len(ops))
	for i, op := range ops {
		results[i] = op(a, b)
	}
	return results
}
