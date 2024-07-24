package main

import "fmt"

// Определение возраста совершеннолетия
// Определите тип Age на основе int. Напишите функцию, которая принимает возраст и возвращает сообщение о том,
// является ли человек совершеннолетним (возраст 18 лет и старше) или нет.
func checkAdult(age Age) string {
	if age >= 18 {
		return "Adult"
	}
	return "Not an Adult"
}

// Проверка на четность
// Определите тип Number на основе int. Напишите функцию, которая принимает число и возвращает сообщение о том,
// является ли оно четным или нечетным.

type Number int

// Функция для проверки четности числа
func isEven(num Number) string {
	if num%2 == 0 {
		return "Even"
	}
	return "Odd"
}

// Проверка допустимого диапазона
// Определите тип Score на основе int. Напишите функцию, которая принимает оценку и возвращает сообщение,
// находится ли она в допустимом диапазоне от 0 до 100

type Score int

// Функция для проверки допустимого диапазона оценки
func isValidScore(score Score) string {
	if score >= 0 && score <= 100 {
		return "Valid Score"
	}
	return "Invalid Score"
}

// Арифметические операции
// Определите тип функции Operation, которая принимает два int и возвращает int. Напишите функции для сложения,
// вычитания и умножения. Используйте тип Operation для вызова этих функций.

type Operation func(int, int) int

// Функция для сложения
func adding(a int, b int) int {
	return a + b
}

// Функция для вычитания
func subtract(a int, b int) int {
	return a - b
}

// Функция для умножения
func multiply(a int, b int) int {
	return a * b
}

// Сравнение чисел
// Определите тип функции Comparator, которая принимает два int и возвращает bool. Напишите функции для проверки
// равенства и сравнения больше/меньше. Используйте тип Comparator для вызова этих функций.

type Comparator func(int, int) bool

// Функция для проверки равенства
func equal(a int, b int) bool {
	return a == b
}

// Функция для проверки, больше ли первое число второго
func greater(a int, b int) bool {
	return a > b
}

// Функция для проверки, меньше ли первое число второго
func less(a int, b int) bool {
	return a < b
}

// Применение функции к числу
// Определите тип функции UnaryOperation, которая принимает int и возвращает int. Напишите функции для удвоения
// и утроения числа. Используйте тип UnaryOperation для вызова этих функций.

type UnaryOperation func(int) int

// Функция для удвоения числа
func double(a int) int {
	return a * 2
}

// Функция для утроения числа
func triple(a int) int {
	return a * 3
}

// Обратный отсчет
// Создайте псевдоним для типа int под названием Count. Напишите функцию, которая принимает Count и выводит
// обратный отсчет до нуля.

type Count = int

// Функция для обратного отсчета
func countdown(c Count) {
	for i := c; i >= 0; i-- {
		fmt.Println(i)
	}
}

// Проверка температуры
// Создайте псевдоним для типа float64 под названием Temperature. Напишите функцию, которая принимает
// Temperature и выводит сообщение о состоянии (ниже нуля, выше нуля или ноль).

type Temperature = float64

// Функция для проверки температуры
func checkTemperature(t Temperature) string {
	if t < 0 {
		return "Below zero"
	} else if t > 0 {
		return "Above zero"
	}
	return "Zero"
}

// Применение скидки
// Создайте псевдоним для типа float64 под названием Price. Напишите функцию, которая принимает Price
// и возвращает новую цену с 10% скидкой.

type Price = float64

// Функция для применения скидки
func applyDiscount(p Price) Price {
	return p * 0.9
}

// Информация о пользователе
// Создайте структуру User с полями Name (строка) и Age (целое число). Напишите функцию, которая
// принимает User и выводит информацию о нем.

type User struct {
	Name string
	Age  int
}

// Функция для вывода информации о пользователе
func printUser(u User) {
	fmt.Printf("Name: %s, Age: %d\n", u.Name, u.Age)
}

// Продукт и его стоимость
// Создайте структуру Product с полями Name (строка) и Price (тип Price). Напишите функцию, которая
// принимает Product и возвращает сообщение о его стоимости.

type Price2 = float64

// Product
// Структура Product
type Product struct {
	Name  string
	Price Price
}

// Функция для вывода информации о продукте
func printProduct(p Product) {
	fmt.Printf("Product: %s, Price: %.2f\n", p.Name, p.Price)
}

// Информация о книге
// Создайте структуру Book с полями Title (строка), Author (строка) и Pages (целое число). Напишите функцию,
// которая принимает Book и выводит информацию о книге.

type Book struct {
	Title  string
	Author string
	Pages  int
}

// Функция для вывода информации о книге
func printBook(b Book) {
	fmt.Printf("Title: %s, Author: %s, Pages: %d\n", b.Title, b.Author, b.Pages)
}

// Изменение данных через указатель
// Создайте структуру Person с полями Name и Age. Напишите функцию, которая принимает указатель на Person
// и обновляет его возраст. Выведите информацию до и после обновления.

type Person3 struct {
	Name string
	Age  int
}

// Функция для обновления возраста через указатель
func updateAge(p *Person3, newAge int) {
	p.Age = newAge
}

// Создание и изменение структуры через указатель
// Создайте структуру Rectangle с полями Width и Height. Напишите функцию, которая принимает указатель
// на Rectangle, вычисляет и обновляет его площадь, а затем выводит обновленную площадь.

type Rectangle struct {
	Width  int
	Height int
}

// Функция для вычисления и обновления площади
func updateArea(r *Rectangle) {
	area := r.Width * r.Height
	fmt.Printf("Area: %d\n", area)
}

// Сравнение двух структур через указатели
// Создайте структуру Coordinate с полями X и Y. Напишите функцию, которая принимает указатели на две
// Coordinate и возвращает сообщение о том, равны ли они или нет.

type Coordinate struct {
	X int
	Y int
}

// Функция для сравнения двух координат
func compareCoordinates(c1, c2 *Coordinate) string {
	if c1.X == c2.X && c1.Y == c2.Y {
		return "Coordinates are equal"
	}
	return "Coordinates are not equal"
}

// Программа для управления студентом
// Создайте структуры Student и Address. Student должен содержать поля Name, Age и Address
// (вложенная структура Address). Напишите функцию, которая выводит информацию о студенте, включая адрес.

type Address2 struct {
	Street string
	City   string
}

type Student struct {
	Name    string
	Age     int
	Address Address2
}

// Функция для вывода информации о студенте
func printStudent(s Student) {
	fmt.Printf("Name: %s, Age: %d\n", s.Name, s.Age)
	fmt.Printf("Address: %s, %s\n", s.Address.Street, s.Address.City)
}

// Работа с сотрудниками
// Создайте структуры Employee и Department. Employee должен содержать поля Name, ID и Department
// (вложенная структура Department). Напишите функцию, которая выводит информацию о сотруднике и его отделе.

type Department struct {
	Name string
}

type Employee struct {
	Name       string
	ID         int
	Department Department
}

// Функция для вывода информации о сотруднике
func printEmployee(e Employee) {
	fmt.Printf("Name: %s, ID: %d\n", e.Name, e.ID)
	fmt.Printf("Department: %s\n", e.Department.Name)
}

// Дерево сотрудников
// Создайте структуру Employee с полем Manager, которое указывает на другого сотрудника (вложенная структура).
// Напишите функцию, которая рекурсивно выводит информацию о сотруднике и его менеджере.

type Employee2 struct {
	Name     string
	Position string
	Manager  *Employee2
}

// Функция для рекурсивного вывода информации о сотруднике и его менеджере
func printEmployeeHierarchy(e *Employee2) {
	if e == nil {
		return
	}
	fmt.Printf("Name: %s, Position: %s\n", e.Name, e.Position)
	if e.Manager != nil {
		fmt.Printf("Manager: %s\n", e.Manager.Name)
		printEmployeeHierarchy(e.Manager)
	}
}

// Временная структура для хранения информации о книге
// Создайте анонимную структуру для хранения информации о книге с полями Title и Author. Напишите функцию,
// которая принимает эту анонимную структуру и выводит информацию о книге.
func printBookInfo(book struct {
	Title  string
	Author string
}) {
	fmt.Printf("Title: %s, Author: %s\n", book.Title, book.Author)
}

// Временная структура для хранения информации о продукте
// Создайте анонимную структуру для хранения информации о продукте с полями Name и Price. Напишите функцию,
// которая принимает эту анонимную структуру и возвращает строку с описанием продукта.
func describeProduct(product struct {
	Name  string
	Price float64
}) string {
	return fmt.Sprintf("Product: %s, Price: %.2f", product.Name, product.Price)
}

// Временная структура для хранения информации о событии
// Создайте анонимную структуру для хранения информации о событии с полями EventName и Date.
// Напишите функцию, которая принимает эту анонимную структуру и возвращает строку с описанием события.
func describeEvent(event struct {
	EventName string
	Date      string
}) string {
	return fmt.Sprintf("Event: %s, Date: %s", event.EventName, event.Date)
}
