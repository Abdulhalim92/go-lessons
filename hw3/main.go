package main

import "fmt"

func main() {
	PrintGreeting()
	DisplaySeparator()
	NotifyStart()

	message := GetWelcomeMessage()
	fmt.Println(message) // Output: Welcome!

	pi := GetPiValue()
	fmt.Println(pi) // Output: 3.14

	active := IsServerActive()
	fmt.Println(active) // Output: true

	PrintNumber(42) // Output: 42

	GreetUser("Alice") // Output: Hello, Alice!

	ToggleLight(true)  // Output: Light on
	ToggleLight(false) // Output: Light off

	sum := Add(3, 4)
	fmt.Println(sum) // Output: 7

	result := Concat("Hello, ", "World!")
	fmt.Println(result) // Output: Hello, World!

	even := IsEven(4)
	odd := IsEven(5)
	fmt.Println(even) // Output: true
	fmt.Println(odd)  // Output: false

	sum2 := adder(5, 7)
	fmt.Println(sum2) // Output: 12

	concatResult := concatenator("Go ", "Language")
	fmt.Println(concatResult) // Output: Go Language

	positive := isPositive(10)
	negative := isPositive(-5)
	fmt.Println(positive) // Output: true
	fmt.Println(negative) // Output: false

	sumFunc := func(x, y int) int {
		return x + y
	}
	calcResult := Calculate(3, 4, sumFunc)
	fmt.Println(calcResult) // Output: 7

	printState := func(state bool) {
		if state {
			fmt.Println("State is true")
		} else {
			fmt.Println("State is false")
		}
	}
	Execute(true, printState) // Output: State is true

	doubleFunc := func(x int) int {
		return x * 2
	}
	doubleFuncResult := Apply(5, doubleFunc)
	fmt.Println(doubleFuncResult) // Output: 10

	multiplyBy2 := Multiplier(2)
	multipleResult := multiplyBy2(5)
	fmt.Println(multipleResult) // Output: 10

	repeat3Times := StringRepeater(3)
	repeaterResult := repeat3Times("Hi!")
	fmt.Println(repeaterResult) // Output: Hi!Hi!Hi!

	printIfTrue := ConditionalPrinter(true)
	printIfTrue("This will be printed") // Output: This will be printed

	printIfFalse := ConditionalPrinter(false)
	printIfFalse("This will not be printed")
}
