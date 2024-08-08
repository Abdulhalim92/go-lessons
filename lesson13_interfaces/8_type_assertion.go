package main

import (
	"fmt"
)

func typeAssertion() {
	var i interface{} = "hello"

	// Утверждение типа
	s, ok := i.(string)
	if ok {
		fmt.Println("String value:", s) // Output: String value: hello
	} else {
		fmt.Println("Not a string")
	}

	// Попытка утверждения неверного типа
	n, ok := i.(int)
	if ok {
		fmt.Println("Int value:", n)
	} else {
		fmt.Println("Not an int") // Output: Not an int
	}
}

// Утверждение типа в функции
func describe(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Println("This is a string:", v)
	case int:
		fmt.Println("This is an int:", v)
	case bool:
		fmt.Println("This is a bool:", v)
	default:
		fmt.Println("Unknown type")
	}
}

func assertionInFunction() {
	describe("hello") // Output: This is a string: hello
	describe(42)      // Output: This is an int: 42
	describe(true)    // Output: This is a bool: true
}
