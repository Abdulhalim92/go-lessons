package main

import "fmt"

// Напишите программу для управления персональными данными. Реализуйте функции для хранения и вывода имени и возраста человека с использованием указателей.
func storeName(namePtr *string, name string) {
	*namePtr = name
}

func storeAge(agePtr *int, age int) {
	*agePtr = age
}

func printPerson(namePtr *string, agePtr *int) {
	fmt.Printf("Имя: %s, Возраст: %d лет\n", *namePtr, *agePtr)
}

// Напишите программу для учета посещаемости студентов на курсе. Реализуйте функции для отметки студента как посетившего курс и вывода списка студентов с их статусом посещения.
func markAttendance(namePtr *string, attendedPtr *bool) {
	*attendedPtr = true
}

func printAttendance(namePtr *string, attendedPtr *bool) {
	status := "не посетил"
	if *attendedPtr {
		status = "посетил"
	}
	fmt.Printf("Студент %s - %s\n", *namePtr, status)
}

// Напишите программу для учета результатов тестирования студентов. Реализуйте функции для добавления результатов теста и вывода среднего балла.
func addTestScore(scoresPtr *[]int, score int) {
	*scoresPtr = append(*scoresPtr, score)
}

func averageScore(scoresPtr *[]int) float64 {
	sum := 0
	count := len(*scoresPtr)
	for _, score := range *scoresPtr {
		sum += score
	}
	if count > 0 {
		return float64(sum) / float64(count)
	}
	return 0.0
}

// Напишите программу для учета голосов на выборах. Реализуйте функции для подсчета голосов за каждого кандидата и определения победителя.
func vote(candidatePtr *string, votesPtr *int) {
	*votesPtr++
}

func winner(candidate1VotesPtr *int, candidate2VotesPtr *int) string {
	if *candidate1VotesPtr > *candidate2VotesPtr {
		return "Кандидат 1"
	} else if *candidate2VotesPtr > *candidate1VotesPtr {
		return "Кандидат 2"
	} else {
		return "Ничья"
	}
}

// Напишите программу для управления производственным процессом. Реализуйте функции для добавления нового рабочего места и увеличения числа рабочих мест.
func addWorkplace(workplacesPtr *int) {
	*workplacesPtr++
}

func increaseWorkplaces(workplacesPtr *int, delta int) {
	*workplacesPtr += delta
}

// Напишите программу для учета баланса кредитной карты. Реализуйте функции для добавления средств и проверки текущего баланса.
func addFunds(balancePtr *float64, amount float64) {
	*balancePtr += amount
}

func checkBalance(balancePtr *float64) float64 {
	return *balancePtr
}

// Напишите программу для учета расходов на семейный бюджет. Реализуйте функции для добавления новой записи расходов и вывода общей суммы расходов.
func addExpense(totalPtr *float64, amount float64) {
	*totalPtr += amount
}

func printTotalExpenses(totalPtr *float64) {
	fmt.Printf("Общая сумма расходов: %.2f\n", *totalPtr)
}

// Напишите программу для учета оценок по предметам студентов. Реализуйте функции для добавления оценки по предмету и вывода среднего балла.
func addGrade(mathPtr *int, grade int) {
	*mathPtr = grade
}

func printAverageGrade(mathPtr *int, physicsPtr *int, chemistryPtr *int) float64 {
	sum := *mathPtr + *physicsPtr + *chemistryPtr
	return float64(sum) / 3.0
}
