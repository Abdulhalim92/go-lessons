package main

import "fmt"

// Обратный отсчет
// Создайте псевдоним для типа int под названием Count. Напишите функцию, которая принимает
// Count и выводит обратный отсчет до нуля.

type Count = int

// Функция для обратного отсчета
func countdown(c Count) {
	for i := c; i >= 0; i-- {
		fmt.Println(i)
	}
}

// Проверка уровня батареи
// Создайте псевдоним для типа int под названием BatteryLevel. Напишите функцию, которая принимает
// BatteryLevel и возвращает сообщение о том, низкий, средний или высокий уровень заряда.

type BatteryLevel = int

// Функция для проверки уровня батареи
func checkBatteryLevel(level BatteryLevel) string {
	if level < 20 {
		return "Low"
	} else if level < 80 {
		return "Medium"
	}
	return "High"
}

// Определение категории веса
// Создайте псевдоним для типа float64 под названием Weight. Напишите функцию, которая принимает
// Weight и возвращает сообщение о том, в какой категории веса находится объект: "Light", "Medium" или "Heavy".

type Weight = float64

// Функция для определения категории веса
func categorizeWeight(w Weight) string {
	if w < 50 {
		return "Light"
	} else if w < 100 {
		return "Medium"
	}
	return "Heavy"
}

// Оценка успеваемости
// Создайте псевдоним для типа int под названием Grade. Напишите функцию, которая принимает Grade
// и возвращает сообщение о том, является ли оценка удовлетворительной (50 и выше) или нет.

type Grade = int

// Функция для оценки успеваемости
func evaluateGrade(g Grade) string {
	if g >= 50 {
		return "Satisfactory"
	}
	return "Unsatisfactory"
}

// Оценка состояния здоровья
// Создайте псевдонимы для типов float64 под названиями BMI и BloodPressure. Напишите функцию, которая
// принимает BMI и BloodPressure, и возвращает сообщение о состоянии здоровья: "Healthy", "At risk" или "Unhealthy"

type BloodPressure = float64

// Функция для оценки состояния здоровья
func evaluateHealth(bmi BMI, bp BloodPressure) string {
	if bmi >= 18.5 && bmi <= 24.9 && bp >= 90 && bp <= 120 {
		return "Healthy"
	} else if (bmi < 18.5 || bmi > 24.9) && (bp < 90 || bp > 120) {
		return "Unhealthy"
	}
	return "At risk"
}
