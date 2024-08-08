package main

import (
	"fmt"
)

// Animal
// Интерфейс Animal с методом Speak
type Animal interface {
	Speak() string
}

// Dog
// Тип Dog реализует метод Speak
type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

// Cat
// Тип Cat реализует метод Speak
type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

// Cow Тип Cow реализует метод Speak
type Cow struct{}

func (c Cow) Speak() string {
	return "Moo!"
}

// Describe
// Функция для работы с любым Animal
func Describe(a Animal) {
	fmt.Println(a.Speak())
}

func duck() {
	dog := Dog{}
	cat := Cat{}
	cow := Cow{}

	Describe(dog) // Output: Woof!
	Describe(cat) // Output: Meow!
	Describe(cow) // Output: Moo!
}
