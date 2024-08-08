package main

import "fmt"

// Describer
// Интерфейс Describer с методом Describe
type Describer interface {
	Describe() string
}

func notInitialized() {
	var d Describer

	// Проверка на nil
	if d == nil {
		fmt.Println("d is nil") // Output: d is nil
	} else {
		fmt.Println("d is not nil")
	}

	// Попытка вызова метода на nil-интерфейсе вызовет панику
	// fmt.Println(d.Describe()) // Panic: runtime error: invalid memory address or nil pointer dereference
}

func (p *Person) Describe() string {
	return "Person named " + p.Name
}

func nilAssigment() {
	var d Describer
	var p *Person

	// Присваиваем nil значение интерфейсу
	d = p

	// Проверка на nil
	if d == nil {
		fmt.Println("d is nil") // Output: d is nil
	} else {
		fmt.Println("d is not nil")
	}

	// Попытка вызова метода на nil-конкретном значении
	// fmt.Println(d.Describe()) // Panic: runtime error: invalid memory address or nil pointer dereference
}

func nilChecking() {
	var d Describer
	var p *Person

	// Присваиваем nil значение интерфейсу
	d = p

	// Проверка на nil перед вызовом метода
	if d != nil {
		fmt.Println(d.Describe())
	} else {
		fmt.Println("d is nil") // Output: d is nil
	}
}
