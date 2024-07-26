package main

import "fmt"

// Анонимный тип продукта
// Создайте переменную анонимного типа, которая содержит Name и Price для продукта.
// Напишите функцию, которая принимает эту переменную и выводит информацию о продукте.

// Функция для вывода информации о продукте
func describeProduct2(p struct {
	Name  string
	Price float64
}) {
	fmt.Printf("Product: %s, Price: %.2f\n", p.Name, p.Price)
}

// Анонимный тип человека
// Создайте переменную анонимного типа, которая содержит FirstName, LastName и Age для человека.
// Напишите функцию, которая принимает эту переменную и выводит полное имя и возраст человека.

// Функция для вывода информации о человеке
func describePerson3(p struct {
	FirstName string
	LastName  string
	Age       int
}) {
	fmt.Printf("Name: %s %s, Age: %d\n", p.FirstName, p.LastName, p.Age)
}

// Анонимный тип автомобиля
// Создайте переменную анонимного типа, которая содержит Make, Model и Year для автомобиля.
// Напишите функцию, которая принимает эту переменную и выводит информацию о автомобиле.

// Функция для вывода информации о автомобиле
func describeCar(c struct {
	Make  string
	Model string
	Year  int
}) {
	fmt.Printf("Car: %s %s, Year: %d\n", c.Make, c.Model, c.Year)
}

// Анонимный тип книги
// Создайте переменную анонимного типа, которая содержит Title и Author для книги.
// Напишите функцию, которая принимает эту переменную и выводит информацию о книге.

// Функция для вывода информации о книге
func describeBook(b struct {
	Title  string
	Author string
}) {
	fmt.Printf("Book: %s by %s\n", b.Title, b.Author)
}

// Анонимный тип события
// Создайте переменную анонимного типа, которая содержит Title, Date и Location (анонимная структура
// с полями Street и City) для события. Напишите функцию, которая принимает эту переменную и выводит
// информацию о событии.

// Функция для вывода информации о событии
func describeEvent2(e struct {
	Title    string
	Date     string
	Location struct {
		Street string
		City   string
	}
}) {
	fmt.Printf("Event: %s, Date: %s, Location: %s, %s\n", e.Title, e.Date, e.Location.Street, e.Location.City)
}
