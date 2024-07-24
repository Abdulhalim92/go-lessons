package main

import "fmt"

func main() {
	var balance float64 = 100000.0
	depositAmount := 50000.0
	balancePtr := &balance
	for i := 0; i < 10; i++ {
		monthlyDeposit(balancePtr, depositAmount)
	}

	var loanAmount float64 = 3000000.0
	annualInterestRate := 10.0
	numMonths := 60 // Предположим, что кредит погашается за 5 лет
	monthlyPayment := calculateMonthlyPayment(loanAmount, annualInterestRate, numMonths)
	balancePtr2 := &loanAmount
	for i := 0; i < numMonths; i++ {
		makeMonthlyPayment(balancePtr2, monthlyPayment)
	}

	var transferBalance float64 = 500000.0
	transfer(&transferBalance, 100000.0)

	var deposit float64 = 100000.0
	annualInterestRate = 5.0
	for i := 0; i < 12; i++ { // Начисление процентов за 12 месяцев
		applyMonthlyInterest(&deposit, annualInterestRate)
	}

	var rate float64 = 74.5
	amount := 100.0
	ratePtr := &rate
	amountPtr := &amount
	for i := 0; i < 6; i++ {
		newRate := 70.0 + float64(i) // пример обновления курса
		updateExchangeRate(ratePtr, newRate)
		convertToRub(amountPtr, rate)
		fmt.Printf("Конвертированная сумма: %.2f RUB\n", *amountPtr)
	}

	balance = 200000.0
	interestRate := 5.0
	balancePtr3 := &balance
	for i := 0; i < 12; i++ {
		addMonthlyInterest(balancePtr3, interestRate)
	}

	var dailyExpenses float64
	expensesPtr := &dailyExpenses
	expense := 800.0
	for i := 0; i < 30; i++ {
		expense = expense + float64(i)
		addDailyExpense(expensesPtr, expense)
	}

	balance = 500000.0
	interestRate = 6.0
	balancePtr4 := &balance
	for i := 0; i < 5; i++ {
		addYearlyInterest(balancePtr4, interestRate)
	}

	balance = 100000.0
	commissionRate := 1.0
	balancePtr5 := &balance
	addTransactionWithCommission(balancePtr5, 50000.0, commissionRate)

	balance = 300000.0
	incomeRate := 7.0
	balancePtr6 := &balance
	for i := 0; i < 5; i++ {
		addYearlyIncome(balancePtr6, incomeRate)
	}

	var depositBal float64 = 10000.0
	depositAmount = 5000.0
	depositBalance(&depositBal, depositAmount)

	var withdrawBalance float64 = 20000.0
	withdrawAmount := 25000.0
	withdraw(&withdrawBalance, withdrawAmount)

	var checkBal float64 = 15000.0
	checkBalance(&checkBal)

	var interestBal float64 = 50000.0
	interestRate = 2.0
	applyInterest(&interestBal, interestRate)

	var balance1 float64 = 30000.0
	var balance2 float64 = 15000.0
	transferAmount := 10000.0
	transferFunds(&balance1, &balance2, transferAmount)

	var balance7 float64 = 40000.0
	withdrawAmount2 := 15000.0
	withdrawLimit2 := 10000.0
	withdrawWithLimit(&balance7, withdrawAmount2, withdrawLimit2)
}
