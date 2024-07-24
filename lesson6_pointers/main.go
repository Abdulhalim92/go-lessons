package main

import (
	"fmt"
)

func main() {
	var x int = 10
	var ptr *int // объявление указателя

	ptr = &x // присвоение адреса переменной x указателю ptr

	fmt.Println("Значение x:", x)                            // вывод значения x
	fmt.Println("Адрес x:", &x)                              // вывод адреса переменной x
	fmt.Println("Значение ptr:", ptr)                        // вывод значения указателя ptr
	fmt.Println("Значение, на которое указывает ptr:", *ptr) // разыменование указателя

	beginnerLvl()
	middleLvl()
	advancedLvl()
}

func beginnerLvl() {
	x := 5
	fmt.Println("Значение x до изменения:", x)
	updateValue(&x)                               // передаем указатель на x в функцию updateValue
	fmt.Println("Значение x после изменения:", x) // x теперь равно 15

	x = 5
	y := 10
	fmt.Println("Значения до обмена:", x, y)
	swap(&x, &y)
	fmt.Println("Значения после обмена:", x, y) // теперь x=10, y=5

	var ptr *int
	x = 42
	printValue(ptr) // "Указатель пуст"
	ptr = &x
	printValue(ptr) // "Значение, на которое указывает указатель: 42"

	x = 10
	fmt.Println("Значение x до увеличения:", x)
	increment(&x)
	fmt.Println("Значение x после увеличения:", x) // x теперь равно 11

	x = 5
	y = -10
	fmt.Println("Число x положительное?", isPositive(&x)) // true
	fmt.Println("Число y положительное?", isPositive(&y)) // false

	str := "Hello"
	fmt.Println("Строка до изменения:", str)
	changeString(&str)
	fmt.Println("Строка после изменения:", str) // "Hello Go"

	x = 6
	fmt.Println("Значение x до удвоения:", x)
	double(&x)
	fmt.Println("Значение x после удвоения:", x) // x теперь равно 12

	x = 7
	fmt.Println("Число x четное?", isEven(&x)) // false
	y = 12
	fmt.Println("Число y четное?", isEven(&y)) // true

	var ptr2 *int
	x = 42
	fmt.Println("Указатель ptr nil?", checkNil(ptr2)) // true
	ptr2 = &x
	fmt.Println("Указатель ptr nil?", checkNil(ptr2)) // false
}

func middleLvl() {
	var name string
	var age int
	namePtr := &name
	agePtr := &age
	storeName(namePtr, "Иван")
	storeAge(agePtr, 30)
	printPerson(namePtr, agePtr)

	var name1, name2 string
	var attended1, attended2 bool
	namePtr1 := &name1
	namePtr2 := &name2
	attendedPtr1 := &attended1
	attendedPtr2 := &attended2
	markAttendance(namePtr1, attendedPtr1)
	printAttendance(namePtr1, attendedPtr1)
	printAttendance(namePtr2, attendedPtr2)

	var scores []int
	scoresPtr := &scores
	addTestScore(scoresPtr, 85)
	addTestScore(scoresPtr, 92)
	addTestScore(scoresPtr, 78)
	fmt.Printf("Средний балл: %.2f\n", averageScore(scoresPtr))

	var candidate1Votes, candidate2Votes int
	candidate1VotesPtr := &candidate1Votes
	candidate2VotesPtr := &candidate2Votes
	var (
		cand1 = "Кандидат 1"
		cand2 = "Кандидат 2"
	)
	vote(&cand1, candidate1VotesPtr)
	vote(&cand2, candidate2VotesPtr)
	vote(&cand1, candidate1VotesPtr)
	vote(&cand1, candidate1VotesPtr)
	fmt.Println("Победитель выборов:", winner(candidate1VotesPtr, candidate2VotesPtr))

	var workplaces int
	workplacesPtr := &workplaces
	addWorkplace(workplacesPtr)
	addWorkplace(workplacesPtr)
	increaseWorkplaces(workplacesPtr, 3)
	fmt.Println("Текущее количество рабочих мест:", *workplacesPtr)

	var balance float64
	balancePtr := &balance
	addFunds(balancePtr, 1000.0)
	addFunds(balancePtr, 500.0)
	fmt.Println("Текущий баланс кредитной карты:", checkBalance(balancePtr))

	var totalExpenses float64
	totalPtr := &totalExpenses
	addExpense(totalPtr, 85.5)
	addExpense(totalPtr, 120.0)
	addExpense(totalPtr, 45.75)
	printTotalExpenses(totalPtr)

	var math, physics, chemistry int
	mathPtr := &math
	physicsPtr := &physics
	chemistryPtr := &chemistry
	addGrade(mathPtr, 85)
	addGrade(physicsPtr, 90)
	addGrade(chemistryPtr, 78)
	fmt.Printf("Средний балл по предметам: %.2f\n", printAverageGrade(mathPtr, physicsPtr, chemistryPtr))
}

func advancedLvl() {
	var usdToEur float64 = 0.85
	var amountToConvert float64 = 3000.0
	usdToEurPtr := &usdToEur
	amountPtr := &amountToConvert
	updateExchangeRate(usdToEurPtr, 0.88)
	convertCurrency(amountPtr, *usdToEurPtr)
	fmt.Printf("Конвертированная сумма: %.2f EUR\n", *amountPtr)

	var depositAmount float64 = 100000.0
	totalPtr := &depositAmount
	addInterestRate(totalPtr, 5.0)
	addInterestRate(totalPtr, 4.5)
	addInterestRate(totalPtr, 4.75)
	finalAmount := calculateFinalAmount(totalPtr)
	fmt.Printf("Итоговая сумма по вкладу: %.2f RUB\n", finalAmount)
	if finalAmount > 150000 {
		fmt.Println("Достигнут желаемый результат")
	}

	var loanAmount float64 = 5000000.0
	remainingLoanPtr := &loanAmount
	addMonthlyPayment(remainingLoanPtr, 50000.0)
	addMonthlyPayment(remainingLoanPtr, 60000.0)
	addMonthlyPayment(remainingLoanPtr, 80000.0)
	fmt.Println("Оставшийся остаток по кредиту:", checkRemainingLoan(remainingLoanPtr))

	var totalOperations int
	countPtr := &totalOperations
	addOperationCount(countPtr)
	addOperationCount(countPtr)
	addOperationCount(countPtr)
	printTotalOperations(countPtr)

	var totalCommissions float64
	totalPtr2 := &totalCommissions
	addCommission(totalPtr, 10.5)
	addCommission(totalPtr, 500.0)
	addCommission(totalPtr, 2500.0)
	printTotalCommissions(totalPtr2)

	var eurToUsd float64 = 1.2
	var amountToConvert2 float64 = 8000.0
	eurToUsdPtr := &eurToUsd
	amountPtr3 := &amountToConvert2
	applyCurrencyDiscount(eurToUsdPtr, 5.0)
	convertCurrency2(amountPtr, *eurToUsdPtr)
	fmt.Printf("Конвертированная сумма: %.2f USD\n", *amountPtr3)

	var yearlyIncome float64 = 1200000.0
	incomePtr := &yearlyIncome
	averageMonthlyIncome := calculateAverageMonthlyIncome(incomePtr)
	deductTax(incomePtr, 13.0)
	fmt.Printf("Средний ежемесячный доход: %.2f RUB\n", averageMonthlyIncome)
	if averageMonthlyIncome < 90000 {
		fmt.Println("Доход ниже уровня прожиточного минимума")
	}

	var initialStockPrice float64 = 1000.0
	var totalProfit float64 = initialStockPrice
	totalProfitPtr := &totalProfit
	addProfit(totalProfitPtr, 500.0)
	addProfit(totalProfitPtr, -300.0)
	addProfit(totalProfitPtr, 200.0)
	finalProfit := checkTotalProfit(totalProfitPtr)
	fmt.Printf("Общая прибыль: %.2f RUB\n", finalProfit)
	if finalProfit < 0 {
		fmt.Println("Убыток")
	}

	var dailyExpenses float64
	dailyExpensesPtr := &dailyExpenses
	addDailyExpense(dailyExpensesPtr, 1000.0)
	addDailyExpense(dailyExpensesPtr, 1500.0)
	addDailyExpense(dailyExpensesPtr, 700.0)
	fmt.Printf("Общие расходы за день: %.2f RUB\n", checkDailyExpenses(dailyExpensesPtr))

	var initialInvestment float64 = 500000.0
	var totalIncome float64
	totalIncomePtr := &totalIncome
	addInvestmentIncome(totalIncomePtr, 75000.0)
	addInvestmentIncome(totalIncomePtr, -100000.0)
	addInvestmentIncome(totalIncomePtr, 200000.0)
	finalIncome := checkTotalIncome(totalIncomePtr, initialInvestment)
	fmt.Printf("Общий доход от инвестиций: %.2f RUB\n", finalIncome)
	if finalIncome < initialInvestment {
		fmt.Println("Убыток")
	}
}
