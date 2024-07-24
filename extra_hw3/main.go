package main

import (
	"fmt"
	"strings"
)

func main() {
	GenerateRandomPassword()
	GenerateInvoiceNumber()
	PrintFormattedDate()

	uuid := GenerateUUID()
	fmt.Println("Generated UUID:", uuid)

	original := "Hello, World!"
	reversed := ReverseString(original)
	fmt.Println("Original:", original)
	fmt.Println("Reversed:", reversed)

	text := "Go is a programming language."
	wordCount := CountWords(text)
	fmt.Printf("Number of words in \"%s\": %d\n", text, wordCount)

	PrintMultiplicationTable(7)

	words := []string{"apple", "banana", "grape", "orange"}
	printIfLongerThan5 := func(s string) bool {
		return len(s) > 5
	}
	PrintFilteredStrings(words, printIfLongerThan5)

	data := []int{10, 20, 30, 40, 50}
	CalculateStatistics(data)

	dataMax := []int{10, 5, 8, 12, 3}
	max := FindMax(dataMax)
	fmt.Println("Max number:", max) // Output: Max number: 12

	fullName := FormatName("John", "Doe")
	fmt.Println("Formatted name:", fullName) // Output: Formatted name: Doe, John

	numData := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evens := FilterEvenNumbers(numData)
	fmt.Println("Even numbers:", evens) // Output: Even numbers: [2 4 6 8 10]

	calcResult := calculator(10, 5)
	fmt.Println(calcResult)

	result := stringManipulator("hello world, how are you?")
	fmt.Println(result) // Output: Hello World, How Are You?

	condData := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	isEven := func(num int) bool {
		return num%2 == 0
	}
	evens2 := filter(condData, isEven)
	fmt.Println("Even numbers:", evens2) // Output: Even numbers: [2 4 6 8 10]

	data2 := []int{1, 2, 3, 4, 5}
	double := func(num int) int {
		return num * 2
	}
	ProcessNumbers(data2, double)

	data3 := []string{"apple", "banana", "cherry"}
	uppercase := func(str string) string {
		return strings.ToUpper(str)
	}
	transformed := ApplyToStrings(data3, uppercase)
	fmt.Println(transformed) // Output: [APPLE BANANA CHERRY]

	data4 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	isEven2 := func(num int) bool {
		return num%2 == 0
	}
	FilterAndPrint(data4, isEven2)

	multiplyBy2 := CreateMultiplier(2)
	multiResult := multiplyBy2(5)
	fmt.Println(multiResult) // Output: 10

	greet := CreateFormatter("Hello, %s!")
	greetResult := greet("Alice")
	fmt.Println(greetResult) // Output: Hello, Alice!

	add, subtract := CreateCalculator()
	sum := add(10, 5)
	difference := subtract(10, 5)
	fmt.Println("Sum:", sum)               // Output: Sum: 15
	fmt.Println("Difference:", difference) // Output: Difference: 5
}
