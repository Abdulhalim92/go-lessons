package main

import (
	"fmt"
)

func interfaceValue() {
	var a Animal

	// Присваиваем значение типа Dog переменной интерфейса Animal
	a = Dog{}
	fmt.Printf("Value: %v, Type: %T\n", a, a) // Output: Value: {}, Type: main.Dog
	fmt.Println(a.Speak())                    // Output: Woof!

	// Присваиваем значение типа Cat переменной интерфейса Animal
	a = Cat{}
	fmt.Printf("Value: %v, Type: %T\n", a, a) // Output: Value: {}, Type: main.Cat
	fmt.Println(a.Speak())                    // Output: Meow!
}
