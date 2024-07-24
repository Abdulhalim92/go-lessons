package main

import "fmt"

func main() {
	printNums()
	printSquares()
	printTripeMultiplicationTable()
	printFibonacci()

	gcd := findGCD(10, 15)
	fmt.Printf("GCD of %d and %d is: %d\n", 10, 15, gcd)

	printFizzBuzz()

	// Print all the prime numbers up to 100
	for i := 2; i <= 100; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}

	// Print all the divisors of 100
	printDivisors(100)

	// Calculate the sum of the digits of 100
	sum := sumDigits(100)
	fmt.Printf("Sum of digits of 100 is: %d\n", sum)

	readPositiveNumber()

	// Calculate the factorial of 7
	factorial := calculateFactorial(7)
	fmt.Printf("Factorial of %d is: %d\n", 7, factorial)

	// Check if 121 is a palindrome
	var n int
	fmt.Print("Введите число: ")
	fmt.Scanln(&n)
	if isPalindrome(n) {
		fmt.Println(n, "является палиндромом.")
	} else {
		fmt.Println(n, "не является палиндромом.")
	}

	// Print the pyramid
	fmt.Print("Введите высоту пирамиды: ")
	fmt.Scanln(&n)
	printPyramid(n)

	// Print the multiplication table
	printMultiplicationTable()

	// Print the binomial coefficient
	fmt.Print("Введите высоту треугольника Паскаля: ")
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%d ", binomialCoefficient(i, j))
		}
		fmt.Println()
	}

	// Print the factorial
	printFactorial()

	printSum()

	printSquaresWithCondition()

	printPrimeNumbers()

	printNumbersWithCondition()

	printCubeNumbers()

	printNumbersWithSum()

	printNumbersWithCallback(func(n int) bool {
		return n%7 == 0
	})
}
