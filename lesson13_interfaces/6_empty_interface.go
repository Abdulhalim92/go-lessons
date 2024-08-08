package main

import (
	"fmt"
)

// Хранение значений любого типа
func emptyInterface() {
	var i interface{}

	i = 42
	fmt.Println(i) // Output: 42

	i = "hello"
	fmt.Println(i) // Output: hello

	i = true
	fmt.Println(i) // Output: true
}

// Функции, принимающие параметры любого типа

func PrintAnything(v interface{}) {
	fmt.Println(v)
}

func funcEmptyInterface() {
	PrintAnything(42)      // Output: 42
	PrintAnything("hello") // Output: hello
	PrintAnything(true)    // Output: true
}

// Срезы и карты с пустыми интерфейсами
func emptyInterfaceSliceMap() {
	// Срез значений любого типа
	slice := []interface{}{42, "hello", true}
	for _, v := range slice {
		fmt.Println(v)
	}

	// Карта значений любого типа
	m := map[string]interface{}{
		"age":       30,
		"name":      "Alice",
		"isStudent": false,
	}
	for k, v := range m {
		fmt.Printf("%s: %v\n", k, v)
	}
}
