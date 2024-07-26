package main

// Проверка температуры
// Определите тип Temperature на основе float64. Напишите функцию, которая принимает температуру
// и возвращает сообщение о том, является ли она ниже, выше или равной нулю.

type Temperature float64

// Функция для проверки температуры
func checkTemperature(temp Temperature) string {
	if temp < 0 {
		return "Below zero"
	} else if temp > 0 {
		return "Above zero"
	}
	return "Zero"
}

// Проверка возраста
// Определите тип Age на основе int. Напишите функцию, которая принимает возраст и возвращает
// сообщение о том, является ли человек подростком (возраст от 13 до 19 лет включительно) или нет.

type Age int

// Функция для проверки возраста
func isTeenager(age Age) string {
	if age >= 13 && age <= 19 {
		return "Teenager"
	}
	return "Not a teenager"
}

// Проверка скорости
// Определите тип Speed на основе float64. Напишите функцию, которая принимает скорость и
// возвращает сообщение о том, является ли она допустимой (от 0 до 120 км/ч включительно) или нет.

type Speed float64

// Функция для проверки скорости
func isValidSpeed(speed Speed) string {
	if speed >= 0 && speed <= 120 {
		return "Valid speed"
	}
	return "Invalid speed"
}

// Проверка счета
// Определите тип Score на основе int. Напишите функцию, которая принимает счет и возвращает
// сообщение о том, является ли он положительным, отрицательным или нулевым.

type Score int

// Функция для проверки счета
func checkScore(score Score) string {
	if score > 0 {
		return "Positive score"
	} else if score < 0 {
		return "Negative score"
	}
	return "Zero score"
}

// Классификация BMI
// Определите тип BMI на основе float64. Напишите функцию, которая принимает значение BMI и
// возвращает сообщение о том, в какой категории оно находится: "Underweight", "Normal weight",
// "Overweight", или "Obesity".

type BMI float64

// Функция для классификации BMI
func classifyBMI(bmi BMI) string {
	if bmi < 18.5 {
		return "Underweight"
	} else if bmi >= 18.5 && bmi < 24.9 {
		return "Normal weight"
	} else if bmi >= 25 && bmi < 29.9 {
		return "Overweight"
	}
	return "Obesity"
}
