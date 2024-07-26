package main

import "fmt"

// Описание продукта
// Создайте структуру Product с полями Name и Price. Напишите функцию, которая принимает
// Product и выводит его описание.

type Product struct {
	Name  string
	Price float64
}

// Функция для вывода описания продукта
func describeProduct(p Product) {
	fmt.Printf("Product: %s, Price: %.2f\n", p.Name, p.Price)
}

// Ввод информации о человеке
// Создайте структуру Person с полями FirstName, LastName и Age. Напишите функцию, которая
// принимает Person и выводит его полное имя и возраст.

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// Функция для вывода информации о человеке
func describePerson(p Person) {
	fmt.Printf("Name: %s %s, Age: %d\n", p.FirstName, p.LastName, p.Age)
}

//  Сравнение продуктов
// Создайте структуру Product с полями Name и Price. Напишите функцию, которая принимает два
// Product и возвращает более дорогой продукт.

// Функция для сравнения продуктов
func compareProducts(p1, p2 Product) Product {
	if p1.Price > p2.Price {
		return p1
	}
	return p2
}

// Управление инвентарем
// Создайте структуру Item с полями Name и Quantity. Напишите функцию, которая принимает
// массив Item и возвращает суммарное количество всех предметов.

type Item struct {
	Name     string
	Quantity int
}

// Функция для подсчета суммарного количества предметов
func totalQuantity(items []Item) int {
	total := 0
	for _, item := range items {
		total += item.Quantity
	}
	return total
}

// Управление библиотекой
// Создайте структуру Library с полями Name и Books (массив структур Book, где Book имеет поля Title и Author).
// Напишите функцию, которая принимает Library и выводит информацию о библиотеке и всех книгах в ней.

type Book struct {
	Title  string
	Author string
}

type Library struct {
	Name  string
	Books []Book
}

// Функция для вывода информации о библиотеке и книгах
func describeLibrary(lib Library) {
	fmt.Printf("Library: %s\n", lib.Name)
	for _, book := range lib.Books {
		fmt.Printf("Book: %s by %s\n", book.Title, book.Author)
	}
}
