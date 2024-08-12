package main

import (
	"math"
	"strconv"
	"strings"
	"time"
)

// Длина строки и Количество слов
// Описание: Реализуйте интерфейс StringProcessor, который будет содержать методы Length()
// и WordCount(). Реализуйте структуру MyString, которая будет работать с текстом и
// реализует этот интерфейс.
// Методы:
// Length() int для получения длины строки.
// WordCount() int для подсчета количества слов.

type StringProcessor interface {
	Length() int
	WordCount() int
}

type MyString struct {
	text string
}

func (ms MyString) Length() int {
	return len(ms.text)
}

func (ms MyString) WordCount() int {
	return len(strings.Fields(ms.text))
}

// Форматирование строки
// Описание: Реализуйте интерфейс Formatter с методами ToUpper() и ToLower(). Реализуйте
// структуру MyFormatter, которая будет работать со строкой и реализует этот интерфейс.
// Методы:
// ToUpper() string для преобразования строки в верхний регистр.
// ToLower() string для преобразования строки в нижний регистр.

type Formatter interface {
	ToUpper() string
	ToLower() string
}

type MyFormatter struct {
	text string
}

func (mf MyFormatter) ToUpper() string {
	return strings.ToUpper(mf.text)
}

func (mf MyFormatter) ToLower() string {
	return strings.ToLower(mf.text)
}

// Работа с указателями на числа
// Описание: Реализуйте интерфейс PointerOperations с методами Increment() и Decrement().
// Реализуйте структуру IntPointer с указателем на число, которая реализует этот интерфейс.
// Методы:
// Increment() для увеличения значения числа на 1.
// Decrement() для уменьшения значения числа на 1.

type PointerOperations interface {
	Increment()
	Decrement()
}

type IntPointer struct {
	num *int
}

func (ip *IntPointer) Increment() {
	*ip.num++
}

func (ip *IntPointer) Decrement() {
	*ip.num--
}

// Удаление пробелов в строке
// Описание: Реализуйте интерфейс StringCleaner с методами TrimSpaces() и RemoveSpaces().
// Реализуйте структуру TextCleaner, которая будет работать со строками и реализует этот
// интерфейс.
// Методы:
// TrimSpaces() string для удаления пробелов с начала и конца строки.
// RemoveSpaces() string для удаления всех пробелов из строки.

type StringCleaner interface {
	TrimSpaces() string
	RemoveSpaces() string
}

type TextCleaner struct {
	text string
}

func (tc TextCleaner) TrimSpaces() string {
	return strings.TrimSpace(tc.text)
}

func (tc TextCleaner) RemoveSpaces() string {
	return strings.ReplaceAll(tc.text, " ", "")
}

// Конкатенация строк
// Описание: Реализуйте интерфейс StringConcatenator с методами Concat() и Join().
// Реализуйте структуру StringJoiner, которая будет работать с массивами строк и
// реализует этот интерфейс.
// Методы:
// Concat() string для конкатенации всех строк в массиве.
// Join(separator string) string для объединения всех строк с заданным разделителем.

type StringConcatenator interface {
	Concat() string
	Join(separator string) string
}

type StringJoiner struct {
	strings []string
}

func (sj StringJoiner) Concat() string {
	return strings.Join(sj.strings, "")
}

func (sj StringJoiner) Join(separator string) string {
	return strings.Join(sj.strings, separator)
}

// Извлечение подстроки
// Описание: Реализуйте интерфейс SubstringExtractor с методами Substr(start, length int)
// string и Suffix(length int) string. Реализуйте структуру TextExtractor, которая
// работает со строками и реализует этот интерфейс.
// Методы:
// Substr(start, length int) string для извлечения подстроки заданной длины с указанной позиции.
// Suffix(length int) string для извлечения последних n символов строки.

type SubstringExtractor interface {
	Substr(start, length int) string
	Suffix(length int) string
}

type TextExtractor struct {
	text string
}

func (te TextExtractor) Substr(start, length int) string {
	if start < 0 || start+length > len(te.text) {
		return ""
	}
	return te.text[start : start+length]
}

func (te TextExtractor) Suffix(length int) string {
	if length > len(te.text) {
		return te.text
	}
	return te.text[len(te.text)-length:]
}

// Переворот строки
// Описание: Реализуйте интерфейс StringReverser с методами Reverse() string и
// ReverseWords() string. Реализуйте структуру TextReverser, которая будет работать
// со строками и реализует этот интерфейс.
// Методы:
// Reverse() string для переворота строки.
// ReverseWords() string для переворота слов в строке.

type StringReverser interface {
	Reverse() string
	ReverseWords() string
}

type TextReverser struct {
	text string
}

func reverseRunes(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func (tr TextReverser) Reverse() string {
	return reverseRunes(tr.text)
}

func (tr TextReverser) ReverseWords() string {
	words := strings.Fields(tr.text)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

// Вставка подстроки
// Описание: Реализуйте интерфейс StringInserter с методами InsertAt(index int, substring
// string) string и InsertAfterWord(word, substring string) string. Реализуйте структуру
// TextInserter, которая работает со строками и реализует этот интерфейс.
// Методы:
// InsertAt(index int, substring string) string для вставки подстроки в указанную позицию.
// InsertAfterWord(word, substring string) string для вставки подстроки после указанного слова.

type StringInserter interface {
	InsertAt(index int, substring string) string
	InsertAfterWord(word, substring string) string
}

type TextInserter struct {
	text string
}

func (ti TextInserter) InsertAt(index int, substring string) string {
	if index < 0 || index > len(ti.text) {
		return ti.text
	}
	return ti.text[:index] + substring + ti.text[index:]
}

func (ti TextInserter) InsertAfterWord(word, substring string) string {
	pos := strings.Index(ti.text, word)
	if pos == -1 {
		return ti.text
	}
	return ti.text[:pos+len(word)] + substring + ti.text[pos+len(word):]
}

// Проверка на палиндром
// Описание: Реализуйте интерфейс PalindromeChecker с методами IsPalindrome() bool и
// IsWordPalindrome() bool. Реализуйте структуру TextPalindromeChecker, которая будет
// работать со строками и реализует этот интерфейс.
// Методы:
// IsPalindrome() bool для проверки, является ли строка палиндромом.
// IsWordPalindrome() bool для проверки, является ли строка палиндромом по словам.

type PalindromeChecker interface {
	IsPalindrome() bool
	IsWordPalindrome() bool
}

type TextPalindromeChecker struct {
	text string
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func (pc TextPalindromeChecker) IsPalindrome() bool {
	cleanedText := strings.ReplaceAll(pc.text, " ", "")
	return cleanedText == reverseString(cleanedText)
}

func (pc TextPalindromeChecker) IsWordPalindrome() bool {
	words := strings.Fields(pc.text)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		if words[i] != words[j] {
			return false
		}
	}
	return true
}

// Удаление повторяющихся слов
// Описание: Реализуйте интерфейс DuplicateRemover с методами RemoveDuplicates() и
// RemoveDuplicatesCaseInsensitive(). Реализуйте структуру TextDuplicateRemover,
// которая работает со строками и реализует этот интерфейс.
// Методы:
// RemoveDuplicates() string для удаления дубликатов слов в строке.
// RemoveDuplicatesCaseInsensitive() string для удаления дубликатов слов без учета регистра.

type DuplicateRemover interface {
	RemoveDuplicates() string
	RemoveDuplicatesCaseInsensitive() string
}

type TextDuplicateRemover struct {
	text string
}

func removeDuplicates(words []string) []string {
	encountered := map[string]bool{}
	result := []string{}

	for _, word := range words {
		if !encountered[word] {
			encountered[word] = true
			result = append(result, word)
		}
	}
	return result
}

func (dr TextDuplicateRemover) RemoveDuplicates() string {
	words := strings.Fields(dr.text)
	result := removeDuplicates(words)
	return strings.Join(result, " ")
}

func (dr TextDuplicateRemover) RemoveDuplicatesCaseInsensitive() string {
	words := strings.Fields(dr.text)
	lowerWords := []string{}
	for _, word := range words {
		lowerWords = append(lowerWords, strings.ToLower(word))
	}
	result := removeDuplicates(lowerWords)
	return strings.Join(result, " ")
}

//  Бюджетирование
// Описание: Реализуйте интерфейс Budget с методами AddIncome(amount float64) и
// AddExpense(amount float64). Реализуйте структуру MonthlyBudget, которая будет
// работать с месячным бюджетом и реализует этот интерфейс.
// Методы:
// AddIncome(amount float64) для добавления дохода.
// AddExpense(amount float64) для добавления расхода.
// GetBalance() float64 для получения текущего баланса.

type Budget interface {
	AddIncome(amount float64)
	AddExpense(amount float64)
	GetBalance() float64
}

type MonthlyBudget struct {
	income  float64
	expense float64
}

func (mb *MonthlyBudget) AddIncome(amount float64) {
	mb.income += amount
}

func (mb *MonthlyBudget) AddExpense(amount float64) {
	mb.expense += amount
}

func (mb MonthlyBudget) GetBalance() float64 {
	return mb.income - mb.expense
}

// Валютообменник
// Описание: Реализуйте интерфейс CurrencyConverter с методами ToUSD(amount float64)
// и ToEUR(amount float64). Реализуйте структуру Exchange, которая будет работать
// с конвертацией валют и реализует этот интерфейс.
// Методы:
// ToUSD(amount float64) float64 для конвертации в доллары США.
// ToEUR(amount float64) float64 для конвертации в евро.

type CurrencyConverter interface {
	ToUSD(amount float64) float64
	ToEUR(amount float64) float64
}

type Exchange struct {
	rateToUSD float64
	rateToEUR float64
}

func (e Exchange) ToUSD(amount float64) float64 {
	return amount * e.rateToUSD
}

func (e Exchange) ToEUR(amount float64) float64 {
	return amount * e.rateToEUR
}

// Банковский счет
// Описание: Реализуйте интерфейс BankAccount с методами Deposit(amount float64) и
// Withdraw(amount float64). Реализуйте структуру Account, которая будет работать
// с банковским счетом и реализует этот интерфейс.
// Методы:
// Deposit(amount float64) для депозита средств на счет.
// Withdraw(amount float64) для снятия средств со счета.
// GetBalance() float64 для получения текущего баланса.

type BankAccount interface {
	Deposit(amount float64)
	Withdraw(amount float64)
	GetBalance() float64
}

type Account struct {
	balance float64
}

func (a *Account) Deposit(amount float64) {
	a.balance += amount
}

func (a *Account) Withdraw(amount float64) {
	if amount <= a.balance {
		a.balance -= amount
	}
}

func (a Account) GetBalance() float64 {
	return a.balance
}

// Налоговый расчет
// Описание: Реализуйте интерфейс TaxCalculator с методами CalculateIncomeTax(income
// float64) и CalculateVAT(amount float64). Реализуйте структуру SimpleTaxCalculator,
// которая будет работать с расчетом налогов и реализует этот интерфейс.
// Методы:
// CalculateIncomeTax(income float64) float64 для расчета налога на доход.
// CalculateVAT(amount float64) float64 для расчета НДС.

type TaxCalculator interface {
	CalculateIncomeTax(income float64) float64
	CalculateVAT(amount float64) float64
}

type SimpleTaxCalculator struct {
	incomeTaxRate float64
	vatRate       float64
}

func (tc SimpleTaxCalculator) CalculateIncomeTax(income float64) float64 {
	return income * tc.incomeTaxRate / 100
}

func (tc SimpleTaxCalculator) CalculateVAT(amount float64) float64 {
	return amount * tc.vatRate / 100
}

// Ведение счета-фактуры
// Описание: Реализуйте интерфейс Invoice с методами AddItem(name string, price float64)
// и Total() float64. Реализуйте структуру SimpleInvoice, которая будет работать со
// счетами и реализует этот интерфейс.
// Методы:
// AddItem(name string, price float64) для добавления товара в счет.
// Total() float64 для получения общей суммы счета.

type Invoice interface {
	AddItem(name string, price float64)
	Total() float64
}

type SimpleInvoice struct {
	items []item
}

type item struct {
	name  string
	price float64
}

func (si *SimpleInvoice) AddItem(name string, price float64) {
	si.items = append(si.items, item{name: name, price: price})
}

func (si SimpleInvoice) Total() float64 {
	total := 0.0
	for _, item := range si.items {
		total += item.price
	}
	return total
}

// Управление портфелем акций
// Описание: Реализуйте интерфейс StockPortfolio с методами Buy(symbol string, quantity int)
// и Sell(symbol string, quantity int). Реализуйте структуру Portfolio, которая будет
// работать с управлением акциями и реализует этот интерфейс.
// Методы:
// Buy(symbol string, quantity int) для покупки акций.
// Sell(symbol string, quantity int) для продажи акций.
// GetHoldings() map[string]int для получения текущих владений.

type StockPortfolio interface {
	Buy(symbol string, quantity int)
	Sell(symbol string, quantity int)
	GetHoldings() map[string]int
}

type Portfolio struct {
	holdings map[string]int
}

func (p *Portfolio) Buy(symbol string, quantity int) {
	if _, exists := p.holdings[symbol]; !exists {
		p.holdings[symbol] = 0
	}
	p.holdings[symbol] += quantity
}

func (p *Portfolio) Sell(symbol string, quantity int) {
	if _, exists := p.holdings[symbol]; exists && p.holdings[symbol] >= quantity {
		p.holdings[symbol] -= quantity
	}
}

func (p Portfolio) GetHoldings() map[string]int {
	return p.holdings
}

// Банкомат
// Описание: Реализуйте интерфейс ATM с методами Deposit(amount float64) и
// Withdraw(amount float64). Реализуйте структуру ATMSystem, которая будет
// работать с банкоматом и реализует этот интерфейс.
// Методы:
// Deposit(amount float64) для депозита средств в банкомат.
// Withdraw(amount float64) для снятия средств из банкомата.
// GetBalance() float64 для получения текущего баланса.

type ATM interface {
	Deposit(amount float64)
	Withdraw(amount float64)
	GetBalance() float64
}

type ATMSystem struct {
	balance float64
}

func (atm *ATMSystem) Deposit(amount float64) {
	atm.balance += amount
}

func (atm *ATMSystem) Withdraw(amount float64) {
	if amount <= atm.balance {
		atm.balance -= amount
	}
}

func (atm ATMSystem) GetBalance() float64 {
	return atm.balance
}

// Управление зарплатой сотрудников
// Описание: Реализуйте интерфейс Payroll с методами AddEmployee(name string, salary
// float64) и GetSalary(name string) float64. Реализуйте структуру CompanyPayroll,
// которая будет работать с зарплатами сотрудников и реализует этот интерфейс.
// Методы:
// AddEmployee(name string, salary float64) для добавления сотрудника и его зарплаты.
// GetSalary(name string) float64 для получения зарплаты сотрудника.
// TotalPayroll() float64 для получения общей суммы зарплат.

type Payroll interface {
	AddEmployee(name string, salary float64)
	GetSalary(name string) float64
	TotalPayroll() float64
}

type CompanyPayroll struct {
	employees map[string]float64
}

func (cp *CompanyPayroll) AddEmployee(name string, salary float64) {
	cp.employees[name] = salary
}

func (cp CompanyPayroll) GetSalary(name string) float64 {
	return cp.employees[name]
}

func (cp CompanyPayroll) TotalPayroll() float64 {
	total := 0.0
	for _, salary := range cp.employees {
		total += salary
	}
	return total
}

// Управление библиотекой
// Описание: Реализуйте интерфейс Library с методами AddBook(title, author string)
// и GetBooks() []string. Реализуйте структуру BookLibrary, которая будет работать
// с библиотекой книг и реализует этот интерфейс.
// Методы:
// AddBook(title, author string) для добавления книги в библиотеку.
// GetBooks() []string для получения списка книг в библиотеке.

type Library interface {
	AddBook(title, author string)
	GetBooks() []string
}

type Book struct {
	title  string
	author string
}

type BookLibrary struct {
	books []Book
}

func (bl *BookLibrary) AddBook(title, author string) {
	bl.books = append(bl.books, Book{title: title, author: author})
}

func (bl BookLibrary) GetBooks() []string {
	var titles []string
	for _, book := range bl.books {
		titles = append(titles, book.title)
	}
	return titles
}

// Управление заказами в интернет-магазине
// Описание: Реализуйте интерфейс OrderManager с методами AddOrder(orderID string, amount
// float64) и GetTotalSales() float64. Реализуйте структуру OnlineStore, которая будет
// работать с заказами и реализует этот интерфейс.
// Методы:
// AddOrder(orderID string, amount float64) для добавления заказа.
// GetTotalSales() float64 для получения общей суммы продаж.

type OrderManager interface {
	AddOrder(orderID string, amount float64)
	GetTotalSales() float64
}

type Order struct {
	orderID string
	amount  float64
}

type OnlineStore struct {
	orders []Order
}

func (os *OnlineStore) AddOrder(orderID string, amount float64) {
	os.orders = append(os.orders, Order{orderID: orderID, amount: amount})
}

func (os OnlineStore) GetTotalSales() float64 {
	total := 0.0
	for _, order := range os.orders {
		total += order.amount
	}
	return total
}

// Управление банковскими транзакциями
// Описание: Реализуйте интерфейс TransactionManager с методами AddTransaction(amount float64)
// и GetStatement() []string. Реализуйте структуру BankAccount, которая будет работать с
// банковскими транзакциями и реализует этот интерфейс.
// Методы:
// AddTransaction(amount float64) для добавления транзакции.
// GetStatement() []string для получения выписки по счету.

type TransactionManager interface {
	AddTransaction(amount float64)
	GetStatement() []string
}

type Transaction struct {
	amount  float64
	balance float64
}

type BankAccount2 struct {
	balance      float64
	transactions []Transaction
}

func (ba *BankAccount2) AddTransaction(amount float64) {
	ba.balance += amount
	ba.transactions = append(ba.transactions, Transaction{amount: amount, balance: ba.balance})
}

func (ba BankAccount2) GetStatement() []string {
	var statement []string
	for _, t := range ba.transactions {
		statement = append(statement, "Transaction: "+strconv.FormatFloat(t.amount, 'f', 2, 64)+", Balance: "+strconv.FormatFloat(t.balance, 'f', 2, 64))
	}
	return statement
}

// Управление каталогом продуктов
// Описание: Реализуйте интерфейс ProductCatalog с методами A
// ddProduct(name string, price float64) и GetProductPrice(name string) float64.
// Реализуйте структуру Catalog, которая будет работать с каталогом продуктов и
// реализует этот интерфейс.
// Методы:
// AddProduct(name string, price float64) для добавления продукта в каталог.
// GetProductPrice(name string) float64 для получения цены продукта.
// GetProducts() map[string]float64 для получения всех продуктов и их цен.

type ProductCatalog interface {
	AddProduct(name string, price float64)
	GetProductPrice(name string) float64
	GetProducts() map[string]float64
}

type Catalog struct {
	products map[string]float64
}

func (c *Catalog) AddProduct(name string, price float64) {
	c.products[name] = price
}

func (c Catalog) GetProductPrice(name string) float64 {
	return c.products[name]
}

func (c Catalog) GetProducts() map[string]float64 {
	return c.products
}

// Управление заказами в ресторане
// Описание: Реализуйте интерфейс RestaurantOrderManager с методами
// AddOrder(orderID string, items []string) и GetOrderDetails(orderID string) []string.
// Реализуйте структуру Restaurant, которая будет работать с заказами и реализует
// этот интерфейс.
// Методы:
// AddOrder(orderID string, items []string) для добавления заказа.
// GetOrderDetails(orderID string) []string для получения деталей заказа.
// GetAllOrders() map[string][]string для получения всех заказов.

type RestaurantOrderManager interface {
	AddOrder(orderID string, items []string)
	GetOrderDetails(orderID string) []string
	GetAllOrders() map[string][]string
}

type Restaurant struct {
	orders map[string][]string
}

func (r *Restaurant) AddOrder(orderID string, items []string) {
	r.orders[orderID] = items
}

func (r Restaurant) GetOrderDetails(orderID string) []string {
	return r.orders[orderID]
}

func (r Restaurant) GetAllOrders() map[string][]string {
	return r.orders
}

// Расчет ипотечного кредита
// Описание: Реализуйте интерфейс MortgageCalculator с методами
// CalculateMonthlyPayment(principal, annualRate float64, years int) и
// CalculateTotalPayment(principal, annualRate float64, years int). Реализуйте структуру
// Mortgage, которая будет работать с расчетом ипотечного кредита и реализует этот интерфейс.
// Методы:
// CalculateMonthlyPayment(principal, annualRate float64, years int) float64 для расчета ежемесячного платежа.
// CalculateTotalPayment(principal, annualRate float64, years int) float64 для расчета общей суммы выплат.

type MortgageCalculator interface {
	CalculateMonthlyPayment(principal, annualRate float64, years int) float64
	CalculateTotalPayment(principal, annualRate float64, years int) float64
}

type Mortgage struct{}

func (m Mortgage) CalculateMonthlyPayment(principal, annualRate float64, years int) float64 {
	monthlyRate := annualRate / 12 / 100
	months := float64(years * 12)
	return principal * monthlyRate * math.Pow(1+monthlyRate, months) / (math.Pow(1+monthlyRate, months) - 1)
}

func (m Mortgage) CalculateTotalPayment(principal, annualRate float64, years int) float64 {
	return m.CalculateMonthlyPayment(principal, annualRate, years) * float64(years*12)
}

// Управление клиентами банка
// Описание: Реализуйте интерфейс CustomerManager с методами
// AddCustomer(name string, accountNumber string) и
// GetCustomerDetails(accountNumber string) (string, string).
// Реализуйте структуру Bank, которая будет работать с клиентами банка и реализует этот интерфейс.
// Методы:
// AddCustomer(name string, accountNumber string) для добавления клиента.
// GetCustomerDetails(accountNumber string) (string, string) для получения деталей клиента.
// GetAllCustomers() map[string]string для получения всех клиентов.

type CustomerManager interface {
	AddCustomer(name string, accountNumber string)
	GetCustomerDetails(accountNumber string) (string, string)
	GetAllCustomers() map[string]string
}

type Bank struct {
	customers map[string]string
}

func (b *Bank) AddCustomer(name string, accountNumber string) {
	b.customers[accountNumber] = name
}

func (b Bank) GetCustomerDetails(accountNumber string) (string, string) {
	return accountNumber, b.customers[accountNumber]
}

func (b Bank) GetAllCustomers() map[string]string {
	return b.customers
}

// Учет рабочего времени сотрудников
// Описание: Реализуйте интерфейс TimeTracker с методами ClockIn(employeeID string) и
// ClockOut(employeeID string). Реализуйте структуру WorkTime, которая будет работать
// с учетом рабочего времени сотрудников и реализует этот интерфейс.
// Методы:
// ClockIn(employeeID string) для регистрации начала работы.
// ClockOut(employeeID string) для регистрации окончания работы.
// GetWorkingHours(employeeID string) float64 для получения отработанных часов.

type TimeTracker interface {
	ClockIn(employeeID string)
	ClockOut(employeeID string)
	GetWorkingHours(employeeID string) float64
}

type WorkTime struct {
	clockInTimes  map[string]time.Time
	clockOutTimes map[string]time.Time
}

func (wt *WorkTime) ClockIn(employeeID string) {
	wt.clockInTimes[employeeID] = time.Now()
}

func (wt *WorkTime) ClockOut(employeeID string) {
	wt.clockOutTimes[employeeID] = time.Now()
}

func (wt WorkTime) GetWorkingHours(employeeID string) float64 {
	clockIn := wt.clockInTimes[employeeID]
	clockOut := wt.clockOutTimes[employeeID]
	return clockOut.Sub(clockIn).Hours()
}

// Система учета студентов
// Описание: Реализуйте интерфейс StudentManager с методами
// AddStudent(name string, studentID string) и
// GetStudentDetails(studentID string) (string, string). Реализуйте структуру
// University, которая будет работать с учетной записью студентов и реализует этот
// интерфейс.
// Методы:
// AddStudent(name string, studentID string) для добавления студента.
// GetStudentDetails(studentID string) (string, string) для получения данных студента.
// GetAllStudents() map[string]string для получения списка всех студентов.

type StudentManager interface {
	AddStudent(name string, studentID string)
	GetStudentDetails(studentID string) (string, string)
	GetAllStudents() map[string]string
}

type University struct {
	students map[string]string
}

func (u *University) AddStudent(name string, studentID string) {
	u.students[studentID] = name
}

func (u University) GetStudentDetails(studentID string) (string, string) {
	return studentID, u.students[studentID]
}

func (u University) GetAllStudents() map[string]string {
	return u.students
}

// Управление расходами предприятия
// Описание: Реализуйте интерфейс ExpenseManager с методами
// AddExpense(category string, amount float64) и GetTotalExpenses(category string) float64.
// Реализуйте структуру Enterprise, которая будет работать с расходами предприятия и
// реализует этот интерфейс.
// Методы:
// AddExpense(category string, amount float64) для добавления расхода.
// GetTotalExpenses(category string) float64 для получения общей суммы расходов по категории.
// GetAllExpenses() map[string]float64 для получения всех расходов.

type ExpenseManager interface {
	AddExpense(category string, amount float64)
	GetTotalExpenses(category string) float64
	GetAllExpenses() map[string]float64
}

type Enterprise struct {
	expenses map[string]float64
}

func (e *Enterprise) AddExpense(category string, amount float64) {
	e.expenses[category] += amount
}

func (e Enterprise) GetTotalExpenses(category string) float64 {
	return e.expenses[category]
}

func (e Enterprise) GetAllExpenses() map[string]float64 {
	return e.expenses
}

// Управление учебными курсами
// Описание: Реализуйте интерфейс CourseManager с методами
// AddCourse(courseName string, credits int) и GetCourseCredits(courseName string) int.
// Реализуйте структуру EducationInstitution, которая будет работать с учебными курсами
// и реализует этот интерфейс.
// Методы:
// AddCourse(courseName string, credits int) для добавления курса.
// GetCourseCredits(courseName string) int для получения количества кредитов курса.
// GetAllCourses() map[string]int для получения всех курсов.

type CourseManager interface {
	AddCourse(courseName string, credits int)
	GetCourseCredits(courseName string) int
	GetAllCourses() map[string]int
}

type EducationInstitution struct {
	courses map[string]int
}

func (ei *EducationInstitution) AddCourse(courseName string, credits int) {
	ei.courses[courseName] = credits
}

func (ei EducationInstitution) GetCourseCredits(courseName string) int {
	return ei.courses[courseName]
}

func (ei EducationInstitution) GetAllCourses() map[string]int {
	return ei.courses
}

// Учет счетов за коммунальные услуги
// Описание: Реализуйте интерфейс UtilityBillManager с методами
// AddBill(utilityType string, amount float64) и GetTotalBills(utilityType string) float64.
// Реализуйте структуру Household, которая будет работать с учетом коммунальных услуг и
// реализует этот интерфейс.
// Методы:
// AddBill(utilityType string, amount float64) для добавления счета за коммунальные услуги.
// GetTotalBills(utilityType string) float64 для получения общей суммы по типу коммунальной услуги.
// GetAllBills() map[string]float64 для получения всех счетов.

type UtilityBillManager interface {
	AddBill(utilityType string, amount float64)
	GetTotalBills(utilityType string) float64
	GetAllBills() map[string]float64
}

type Household struct {
	bills map[string]float64
}

func (h *Household) AddBill(utilityType string, amount float64) {
	h.bills[utilityType] += amount
}

func (h Household) GetTotalBills(utilityType string) float64 {
	return h.bills[utilityType]
}

func (h Household) GetAllBills() map[string]float64 {
	return h.bills
}
