package main

import "fmt"

// PrintNumber
// Создайте функцию PrintNumber, которая принимает целое число и выводит его на экран.
func PrintNumber(num int) {
	fmt.Println(num)
}

// GreetUser
// Создайте функцию GreetUser, которая принимает строку с именем пользователя и выводит приветствие.
func GreetUser(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// ToggleLight
// Создайте функцию ToggleLight, которая принимает булевое значение и выводит "Light on" или "Light off" в зависимости от значения аргумента.
func ToggleLight(state bool) {
	if state {
		fmt.Println("Light on")
	} else {
		fmt.Println("Light off")
	}
}
