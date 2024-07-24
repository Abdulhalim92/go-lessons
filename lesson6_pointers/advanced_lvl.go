package main

import "fmt"

// Напишите программу для конвертации суммы из долларов в евро по текущему курсу. Начальный курс равен 0.85 евро за доллар.
// Реализуйте функции для обновления курса валюты и конвертации суммы. Если конвертированная сумма превышает 5000 евро,
// выведите сообщение "Превышен лимит конвертации".
func updateExchangeRate(usdToEurPtr *float64, newRate float64) {
	*usdToEurPtr = newRate
}

func convertCurrency(amountPtr *float64, exchangeRate float64) {
	*amountPtr *= exchangeRate
	if *amountPtr > 5000 {
		fmt.Println("Превышен лимит конвертации")
	}
}

// Напишите программу для рассчета сложных процентов по вкладу. Начальная сумма вклада равна 100000 рублей.
// Реализуйте функции для добавления процентной ставки и расчета итоговой суммы через 3 года. Если итоговая сумма
// превышает 150000 рублей, выведите сообщение "Достигнут желаемый результат".
func addInterestRate(totalPtr *float64, interestRate float64) {
	*totalPtr *= 1 + interestRate/100.0
}

func calculateFinalAmount(totalPtr *float64) float64 {
	return *totalPtr
}

// Напишите программу для учета ежемесячных выплат по ипотеке. Начальная сумма кредита равна 5000000 рублей.
// Реализуйте функции для добавления ежемесячного платежа и проверки оставшейся суммы. Если оставшаяся сумма становится
// меньше 1000000 рублей, выведите сообщение "Оплачен половинный остаток кредита".
func addMonthlyPayment(remainingLoanPtr *float64, payment float64) {
	*remainingLoanPtr -= payment
	if *remainingLoanPtr < 1000000 {
		fmt.Println("Оплачен половинный остаток кредита")
	}
}

func checkRemainingLoan(remainingLoanPtr *float64) float64 {
	return *remainingLoanPtr
}

// Напишите программу для учета операций по кредитной карте. Начальный лимит операций равен 50.
// Реализуйте функции для добавления новой операции и вывода количества выполненных операций.
// Если количество операций превысит 100, выведите сообщение "Превышено допустимое количество операций".
func addOperationCount(countPtr *int) {
	*countPtr++
	if *countPtr > 100 {
		fmt.Println("Превышено допустимое количество операций")
	}
}

func printTotalOperations(countPtr *int) {
	fmt.Printf("Общее количество операций: %d\n", *countPtr)
}

// Напишите программу для учета комиссий за транзакции на банковском счете. Начальная сумма комиссий равна 0.
// Реализуйте функции для добавления комиссии и вывода общей суммы комиссий. Если общая сумма комиссий превысит 3000 рублей,
// выведите сообщение "Превышен лимит комиссий".
func addCommission(totalPtr *float64, commission float64) {
	*totalPtr += commission
	if *totalPtr > 3000 {
		fmt.Println("Превышен лимит комиссий")
	}
}

func printTotalCommissions(totalPtr *float64) {
	fmt.Printf("Общая сумма комиссий: %.2f\n", *totalPtr)
}

// Напишите программу для конвертации суммы из евро в доллары по текущему курсу. Начальный курс равен 1.2 доллара за евро.
// Реализуйте функции для применения скидки курса валюты и конвертации суммы. Если конвертированная сумма превышает 10000 долларов,
// выведите сообщение "Превышен лимит конвертации".
func applyCurrencyDiscount(ratePtr *float64, discount float64) {
	*ratePtr *= 1 - discount/100.0
}

func convertCurrency2(amountPtr *float64, exchangeRate float64) {
	*amountPtr *= exchangeRate
	if *amountPtr > 10000 {
		fmt.Println("Превышен лимит конвертации")
	}
}

// Напишите программу для расчета среднего ежемесячного дохода. Начальная сумма дохода за год равна 1200000 рублей.
// Реализуйте функции для вычисления среднего дохода за месяц и вычета налогов в размере 13%. Если средний доход за месяц
// становится меньше 90000 рублей, выведите сообщение "Доход ниже уровня прожиточного минимума".
func calculateAverageMonthlyIncome(incomePtr *float64) float64 {
	return *incomePtr / 12.0
}

func deductTax(incomePtr *float64, taxRate float64) {
	*incomePtr *= 1 - taxRate/100.0
}

// Напишите программу для учета покупки и продажи акций. Начальная стоимость акции равна 1000 рублей.
// Реализуйте функции для добавления прибыли и проверки общей прибыли после продажи всех акций. Если общая прибыль
// становится отрицательной, выведите сообщение "Убыток".
func addProfit(totalProfitPtr *float64, profit float64) {
	*totalProfitPtr += profit
}

func checkTotalProfit(totalProfitPtr *float64) float64 {
	return *totalProfitPtr
}

// Напишите программу для учета ежедневных расходов. Начальный лимит расходов равен 2000 рублей в день.
// Реализуйте функции для добавления нового расхода и проверки общей суммы расходов за день.
// Если общая сумма расходов превышает 3000 рублей, выведите сообщение "Превышен лимит расходов за день".
func addDailyExpense(dailyExpensesPtr *float64, expense float64) {
	*dailyExpensesPtr += expense
	if *dailyExpensesPtr > 3000 {
		fmt.Println("Превышен лимит расходов за день")
	}
}

func checkDailyExpenses(dailyExpensesPtr *float64) float64 {
	return *dailyExpensesPtr
}

// Напишите программу для учета доходов от инвестиций. Начальная сумма инвестиций равна 500000 рублей. Реализуйте
// функции для добавления прибыли и проверки общего дохода от инвестиций. Если общий доход становится меньше начальной
// суммы инвестиций, выведите сообщение "Убыток".
func addInvestmentIncome(totalIncomePtr *float64, income float64) {
	*totalIncomePtr += income
}

func checkTotalIncome(totalIncomePtr *float64, initialInvestment float64) float64 {
	return *totalIncomePtr + initialInvestment
}
