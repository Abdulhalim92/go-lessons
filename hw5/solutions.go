package main

import (
	"fmt"
	"time"
)

// Напишите программу, которая используя конструкцию циклов выводит числа от 1 до 10.
func printNums() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

// Напишите программу, которая используя конструкцию циклов выводит квадраты чисел от 1 до 5.
func printSquares() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i * i)
	}
}

// Напишите программу, которая выводит таблицу умножения на 3 от 1 до 10.
func printTripeMultiplicationTable() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i * 3)
	}
}

// Напишите программу, которая выводит первые 10 чисел последовательности Фибоначчи.
func printFibonacci() {
	a, b := 0, 1
	for i := 0; i < 10; i++ {
		fmt.Println(b)
		a, b = b, a+b
	}
}

// Напишите программу, которая находит наибольший общий делитель (НОД) двух чисел, используя алгоритм Евклида.
func findGCD(a, b int) int {
	if b == 0 {
		return a
	}
	return findGCD(b, a%b)
}

// Напишите программу, которая выводит числа от 1 до 100, заменяя числа, кратные 3, на "Fizz", числа, кратные 5, на "Buzz", а числа, кратные 3 и 5, на "FizzBuzz".
func printFizzBuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

// Напишите программу, которая проверяет, является ли число простым.
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Напишите программу, которая выводит все делители числа.
func printDivisors(n int) {
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			fmt.Println(i)
		}
	}
}

// Напишите программу, которая находит сумму цифр числа.
func sumDigits(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

// Напишите программу, которая запрашивает у пользователя ввод положительного числа и повторяет запрос, пока не будет введено отрицательное число.
func readPositiveNumber() int {
	var n int
	for {
		fmt.Print("Enter a positive number: ")
		fmt.Scan(&n)
		if n < 0 {
			break
		}
	}
	return n
}

// Напишите программу, которая вычисляет произведение всех чисел от 1 до введённого числа n, но прекращает выполнение, если результат превышает 1000.
func calculateFactorial(n int) int {
	product := 1
	for i := 1; i <= n; i++ {
		product *= i
		if product > 1000 {
			break
		}
	}
	return product
}

// Напишите программу, которая запрашивает у пользователя ввод числа и проверяет, является ли оно палиндромом.
func isPalindrome(n int) bool {
	original, reverse := n, 0
	for {
		reverse = reverse*10 + n%10
		n /= 10
		if n == 0 {
			break
		}
	}
	return original == reverse
}

// Напишите программу, которая выводит пирамиду из символов '*' высотой n, введённого пользователем
func printPyramid(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// Напишите программу, которая выводит таблицу умножения от 1 до n, где n - введенное пользователем число.
func printMultiplicationTable() {
	n := readPositiveNumber()
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
		}
	}
}

// Напишите программу, которая выводит треугольник Паскаля высотой n, где n - введенное пользователем число.
func factorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	} // result = 1*1*2*3*...*n = n!
	return result
}

func binomialCoefficient(n, k int) int {
	return factorial(n) / (factorial(k) * factorial(n-k))
}

// Напишите программу, которая бесконечно запрашивает у пользователя ввод числа и выводит его факториал, пока не будет введено число меньше 0.
func printFactorial() {
	for {
		n := readPositiveNumber()
		fmt.Printf("Factorial of %d is: %d\n", n, calculateFactorial(n))
	}
}

// Напишите программу, которая бесконечно запрашивает у пользователя ввод двух чисел и выводит их сумму.
func printSum() {
	for {
		time.Sleep(1 * time.Second)
		n1 := readPositiveNumber()
		n2 := readPositiveNumber()
		fmt.Printf("Sum of %d and %d is: %d\n", n1, n2, n1+n2)
		if n1+n2 < 0 {
			break
		}
	}
}

// Напишите программу, которая выводит все числа от 1 до 100, но пропускает числа, которые являются квадратами.
func printSquaresWithCondition() {
	for i := 1; i <= 100; i++ {
		if isSquare(i) {
			continue
		}
		fmt.Println(i)
	}
}

func isSquare(n int) bool {
	for i := 1; i*i <= n; i++ {
		if i*i == n {
			return true
		}
	}
	return false
}

// Напишите программу, которая выводит только те числа от 1 до 50, которые являются простыми.
func printPrimeNumbers() {
	for i := 1; i <= 50; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}

// Напишите программу, которая выводит все числа от 1 до 30, кроме чисел, которые делятся на 2 или 3.
func printNumbersWithCondition() {
	for i := 1; i <= 30; i++ {
		if i%2 == 0 || i%3 == 0 {
			continue
		}
		fmt.Println(i)
	}
}

// Напишите программу, которая выводит числа от 1 до 50, но прекращает выполнение, если встречает число, которое является кубом.
func printCubeNumbers() {
	for i := 1; i <= 50; i++ {
		if isCube(i) {
			continue
		}
		fmt.Println(i)
	}
}

func isCube(n int) bool {
	for i := 1; i*i*i <= n; i++ {
		if i*i*i == n {
			return true
		}
	}
	return false
}

// Напишите программу, которая выводит числа от 1 до 100, но прекращает выполнение, если сумма всех выведенных чисел превышает 200.
func printNumbersWithSum() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
		if sum > 200 {
			break
		}
		fmt.Println(i)
	}
}

// Напишите программу, которая бесконечно запрашивает у пользователя ввод числа и прекращает выполнение, если введенное число делится на 7.
func printNumbersWithCallback(cond func(int) bool) {
	for {
		n := readPositiveNumber()
		if cond(n) {
			break
		}
		fmt.Println(n)
	}
}
