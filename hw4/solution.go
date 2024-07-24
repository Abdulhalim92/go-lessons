package main

import "fmt"

// PrintGreeting
// Создайте функцию PrintGreeting, которая печатает "Доброе утро!", если dayType равен "утро"; "Добрый день!",
// если dayType равен "день"; и "Добрый вечер!", если dayType равен "вечер". Переменную  dayType вводить с консоли внутри функции.
func PrintGreeting() {
	dayType := "утро" // Можно менять на "день" или "вечер"
	switch dayType {
	case "утро":
		fmt.Println("Доброе утро!")
	case "день":
		fmt.Println("Добрый день!")
	case "вечер":
		fmt.Println("Добрый вечер!")
	default:
		fmt.Println("Привет!")
	}
}

// PrintWeather
// Создайте функцию PrintWeather, которая печатает "Солнечно", если weatherType равен "солнечно"; "Облачно",
// если weatherType равен "облачно"; и "Дождливо", если weatherType равен "дождливо". Переменную  weatherType вводить с консоли внутри функции.
func PrintWeather() {
	weatherType := "солнечно" // Можно менять на "облачно" или "дождливо"
	switch weatherType {
	case "солнечно":
		fmt.Println("Солнечно")
	case "облачно":
		fmt.Println("Облачно")
	case "дождливо":
		fmt.Println("Дождливо")
	default:
		fmt.Println("Погода неизвестна")
	}
}

// PrintTrafficLight
// Создайте функцию PrintTrafficLight, которая печатает "Стоп", если lightColor равен "красный"; "Внимание",
// если lightColor равен "желтый"; и "Идите", если lightColor равен "зеленый". Переменную lightColor вводить с консоли внутри функции.
func PrintTrafficLight() {
	lightColor := "красный" // Можно менять на "желтый" или "зеленый"
	switch lightColor {
	case "красный":
		fmt.Println("Стоп")
	case "желтый":
		fmt.Println("Внимание")
	case "зеленый":
		fmt.Println("Идите")
	default:
		fmt.Println("Цвет неизвестен")
	}
}

// GetGrade
// Создайте функцию GetGrade, которая возвращает оценку "A", "B" или "C" в зависимости от значения переменной score.
// Переменную scope вводить с консоли внутри функции.
func GetGrade() string {
	score := 85 // Можно менять на другие значения
	switch {
	case score >= 90:
		return "A"
	case score >= 75:
		return "B"
	default:
		return "C"
	}
}

// GetDiscount
// Создайте функцию GetDiscount, которая возвращает скидку "10%" - при amount > 1000, "5%" - при 500 < amount <=1000
// или "0%" - при amount <= 500 в зависимости от суммы покупки amount. Переменную amount вводить с консоли внутри функции.
func GetDiscount() string {
	amount := 1500 // Можно менять на другие значения
	switch {
	case amount > 1000:
		return "10%"
	case amount > 500:
		return "5%"
	default:
		return "0%"
	}
}

// GetTemperatureDescription
// Создайте функцию GetTemperatureDescription, которая возвращает "Холодно" (меньше 10) , "Тепло" (с 10 до 25)
// или "Жарко" в зависимости от значения переменной temperature. Переменную temperature вводить с консоли внутри функции.
func GetTemperatureDescription() string {
	temperature := 15 // Можно менять на другие значения
	switch {
	case temperature < 10:
		return "Холодно"
	case temperature <= 25:
		return "Тепло"
	default:
		return "Жарко"
	}
}

// CheckNumber
// Создайте функцию CheckNumber, которая принимает целое число и печатает "Положительное", если число больше нуля;
// "Отрицательное", если меньше нуля; и "Ноль", если равно нулю.
func CheckNumber(num int) {
	switch {
	case num > 0:
		fmt.Println("Положительное")
	case num < 0:
		fmt.Println("Отрицательное")
	default:
		fmt.Println("Ноль")
	}
}

// CheckAge
// Создайте функцию CheckAge, которая принимает возраст и печатает "Совершеннолетний", если возраст 18 и старше; "Несовершеннолетний", если младше 18.
func CheckAge(age int) {
	switch {
	case age >= 18:
		fmt.Println("Совершеннолетний")
	default:
		fmt.Println("Несовершеннолетний")
	}
}

// CheckPassword
// Создайте функцию CheckPassword, которая принимает строку пароля и печатает "Пароль верный", если пароль равен "1234"; и "Пароль неверный" в противном случае.
func CheckPassword(password string) {
	switch password {
	case "1234":
		fmt.Println("Пароль верный")
	default:
		fmt.Println("Пароль неверный")
	}
}

// Add
// Создайте функцию Add, которая принимает два целых числа и возвращает сумму их модулей. Используйте условные операторы для проверки значений чисел.
func Add(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	return a + b
}

// CompareStrings
// Создайте функцию CompareStrings, которая принимает две строки и возвращает "равны", если строки одинаковы,
// и "не равны" в противном случае. Используйте условные операторы для сравнения строк.
func CompareStrings(str1, str2 string) string {
	if str1 == str2 {
		return "равны"
	} else {
		return "не равны"
	}
}

// Max
// Создайте функцию Max, которая принимает два целых числа и возвращает большее из них. Используйте условные операторы для сравнения чисел.
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Создайте переменную operation, которая является функцией, принимающей два целых числа и возвращающей их разность.
// Используйте условные операторы внутри функции для проверки значений чисел.
var operation func(int, int) int = func(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

// Создайте переменную concat, которая является функцией, принимающей две строки и возвращающей их конкатенацию с учетом того,
// что если строки не пустые добавить между ними разделение по пробелу. Используйте условные операторы внутри функции для проверки значений строк.
var concat func(string, string) string = func(a, b string) string {
	if a == "" {
		return b
	}
	if b == "" {
		return a
	}
	return a + " " + b
}

// Создайте переменную multiply, которая является функцией, принимающей два целых числа и возвращающей произведение их модулей.
// Используйте условные операторы внутри функции для проверки значений чисел.
var multiply func(int, int) int = func(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	} else if a < 0 {
		a = -a
	} else if b < 0 {
		b = -b
	}
	return a * b
}

// ApplyOperation
// Создайте функцию ApplyOperation, которая принимает два целых числа и функцию, которая выполняет операцию над модулями этих чисел.
// Примените переданную функцию и выведите результат. Используйте условные операторы внутри функции для проверки значений чисел.
func ApplyOperation(a, b int, operation func(int, int) int) {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	result := operation(a, b)
	fmt.Println("Результат операции:", result)
}

// CheckCondition
// Создайте функцию CheckCondition, которая принимает целое число и функцию, возвращающую true или false. Если функция возвращает true,
// выведите "Условие выполнено", иначе "Условие не выполнено". Используйте условные операторы для проверки значений.
func CheckCondition(num int, condition func(int) bool) {
	if condition(num) {
		fmt.Println("Условие выполнено")
	} else {
		fmt.Println("Условие не выполнено")
	}
}

// FormatAndPrint
// Создайте функцию FormatAndPrint, которая принимает строку и функцию, возвращающую отформатированную строку.
// Если строка пустая, callback – функция должна вывести “Пустая строка”, в противном случае, вывести строку которая
// получается от предыдущей преобразованием в заглавную всех символов. Примените переданную функцию и выведите результат.
// Используйте условные операторы для проверки значений строк.
func FormatAndPrint(s string, formatter func(string) string) {
	if s == "" {
		fmt.Println("Строка пустая")
		return
	}
	result := formatter(s)
	fmt.Println("Отформатированная строка:", result)
}

// CreateMultiplier
// Создайте функцию CreateMultiplier, которая принимает число и возвращает функцию, умножающую переданное число на модуль этого числа.
// Используйте условные операторы внутри возвращаемой функции для проверки значений чисел.
func CreateMultiplier(factor int) func(int) int {
	return func(num int) int {
		if factor == 0 || num == 0 {
			return 0
		} else if factor < 0 {
			return (-1) * num * factor
		}
		return num * factor
	}
}

// CreateGreeter
// Создайте функцию CreateGreeter, которая принимает строку и возвращает функцию, которая принимает имя и возвращает приветствие с
// использованием переданной строки. Если строка пустая, то функция должна возвращать приветствие гостя. Используйте условные операторы
// внутри возвращаемой функции для проверки значений строк.
func CreateGreeter(greeting string) func(string) string {
	return func(name string) string {
		if name == "" {
			return greeting + ", гость"
		}
		return greeting + ", " + name
	}
}

// CreateValidator
// Создайте функцию CreateValidator, которая принимает строку пароля и возвращает функцию, которая принимает строку и возвращает true,
// если строка совпадает с паролем, и false в противном случае. Используйте условные операторы внутри возвращаемой функции для проверки значений строк.
func CreateValidator(password string) func(string) bool {
	return func(input string) bool {
		if input == "" {
			return false
		}
		return input == password
	}
}
