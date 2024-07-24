package main

import (
	"fmt"
	"strings"
)

func main() {
	PrintGreeting()
	PrintWeather()
	PrintTrafficLight()

	grade := GetGrade()
	fmt.Println("Оценка:", grade)

	discount := GetDiscount()
	fmt.Println("Скидка:", discount)

	description := GetTemperatureDescription()
	fmt.Println("Температура:", description)

	CheckNumber(5)
	CheckAge(20)
	CheckPassword("1234")

	sum := Add(3, -5)
	fmt.Println("Сумма:", sum)

	result := CompareStrings("hello", "world")
	fmt.Println("Строки", result)

	maximum := Max(3, 7)
	fmt.Println("Максимальное число:", maximum)

	diff := operation(10, 5)
	fmt.Println("Разность:", diff)

	connResult := concat("Hello", "World")
	fmt.Println("Конкатенация строк:", connResult)

	multiResult := multiply(4, 5)
	fmt.Println("Произведение:", multiResult)

	add := func(a, b int) int {
		return a + b
	}
	ApplyOperation(3, -4, add)

	isPositive := func(num int) bool {
		return num > 0
	}
	CheckCondition(-5, isPositive)

	toUpperCase := func(s string) string {
		return strings.ToUpper(s)
	}
	FormatAndPrint("hello", toUpperCase)

	multiplyBy3 := CreateMultiplier(3)
	result2 := multiplyBy3(4)
	fmt.Println("Результат умножения:", result2)

	greeter := CreateGreeter("Здравствуйте")
	greetResult := greeter("Алексей")
	fmt.Println(greetResult)

	validate := CreateValidator("1234")
	validateResult := validate("1234")
	fmt.Println("Пароль верный?", validateResult)
}
