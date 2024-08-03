package main

import (
	"fmt"
	"strings"
	"time"
)

// Книга и Автор
// Описание: Реализуйте структуру Book с полями title и author, а также метод GetDetails, который возвращает
// строку с деталями книги. Реализуйте структуру Library с массивом книг и метод AddBook, добавляющий книгу
// в библиотеку.
// Методы:
// GetDetails() string для структуры Book
// AddBook(book Book) для структуры Library

type Book struct {
	title  string
	author string
}

// GetDetails Метод для получения деталей книги
func (b Book) GetDetails() string {
	return fmt.Sprintf("Title: %s, Author: %s", b.title, b.author)
}

type Library struct {
	books []Book
}

// AddBook Метод для добавления книги в библиотеку
func (l *Library) AddBook(book Book) {
	l.books = append(l.books, book)
}

// Оценки студента
// Описание: Реализуйте структуру Student с полем grades (срез оценок). Реализуйте метод AddGrade, добавляющий
// оценку, и метод AverageGrade, возвращающий среднее значение оценок.
// Методы:
// AddGrade(grade int)
// AverageGrade() float64

type Student struct {
	grades []int
}

// AddGrade Метод для добавления оценки
func (s *Student) AddGrade(grade int) {
	s.grades = append(s.grades, grade)
}

// AverageGrade Метод для вычисления среднего значения оценок
func (s Student) AverageGrade() float64 {
	sum := 0
	for _, grade := range s.grades {
		sum += grade
	}
	return float64(sum) / float64(len(s.grades))
}

// Круг и Площадь
// Описание: Реализуйте структуру Circle с полем radius. Реализуйте методы Area и Circumference для
// вычисления площади и периметра круга.
// Методы:
// Area() float64
// Circumference() float64

type Circle struct {
	radius float64
}

// Area Метод для вычисления площади круга
func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

// Circumference Метод для вычисления периметра круга
func (c Circle) Circumference() float64 {
	return 2 * 3.14 * c.radius
}

// Контейнер для товаров
// Описание: Реализуйте структуру Item с полями name и price. Реализуйте структуру Inventory с срезом
// товаров и методы для добавления товара и получения общего значения товаров в инвентаре.
// Методы:
// AddItem(item Item)
// TotalValue() float64

type Item struct {
	name  string
	price float64
}

type Inventory struct {
	items []Item
}

// AddItem Метод для добавления товара в инвентарь
func (i *Inventory) AddItem(item Item) {
	i.items = append(i.items, item)
}

// TotalValue Метод для получения общего значения товаров
func (i Inventory) TotalValue() float64 {
	total := 0.0
	for _, item := range i.items {
		total += item.price
	}
	return total
}

// Обработка транзакций
// Описание: Реализуйте структуру Transaction с полями amount и description. Реализуйте метод Process,
// который выводит информацию о транзакции. Реализуйте структуру Account с методами AddTransaction и Balance.
// Методы:
// Process()
// AddTransaction(amount float64, description string)

type Transaction struct {
	amount      float64
	description string
}

// Process Метод для обработки транзакции
func (t Transaction) Process() {
	fmt.Printf("Processing transaction: %s - Amount: %.2f\n", t.description, t.amount)
}

type Account struct {
	transactions []Transaction
}

// AddTransaction Метод для добавления транзакции
func (a *Account) AddTransaction(amount float64, description string) {
	t := Transaction{amount: amount, description: description}
	a.transactions = append(a.transactions, t)
	t.Process()
}

// Balance Метод для получения баланса
func (a Account) Balance() float64 {
	sum := 0.0
	for _, t := range a.transactions {
		sum += t.amount
	}
	return sum
}

// Управление задачами
// Описание: Реализуйте структуру Task с полями title и completed. Реализуйте структуру TaskList
// с срезом задач и методы для добавления задачи и получения количества завершённых задач.
// Методы:
// AddTask(title string)
// CompletedTasks() int

type Task struct {
	title     string
	completed bool
}

type TaskList struct {
	tasks []Task
}

// AddTask Метод для добавления задачи
func (tl *TaskList) AddTask(title string) {
	tl.tasks = append(tl.tasks, Task{title: title})
}

// CompletedTasks Метод для получения количества завершённых задач
func (tl TaskList) CompletedTasks() int {
	count := 0
	for _, task := range tl.tasks {
		if task.completed {
			count++
		}
	}
	return count
}

// Работа с температурой
// Описание: Реализуйте структуру Temperature с полем celsius. Реализуйте методы для преобразования
// температуры в Фаренгейт и Кельвин.
// Методы:
// ToFahrenheit() float64
// ToKelvin() float64

type Temperature struct {
	celsius float64
}

// ToFahrenheit Метод для преобразования температуры в Фаренгейт
func (t Temperature) ToFahrenheit() float64 {
	return t.celsius*9/5 + 32
}

// ToKelvin Метод для преобразования температуры в Кельвин
func (t Temperature) ToKelvin() float64 {
	return t.celsius + 273.15
}

// Управление списком студентов
// Описание: Реализуйте структуру Student с полем name и age. Реализуйте структуру Classroom с срезом
// студентов и методы для добавления студента и получения средней возрастной группы.
// Методы:
// AddStudent(name string, age int)
// AverageAge() float64

type Student2 struct {
	name string
	age  int
}

type Classroom struct {
	students []Student2
}

// AddStudent Метод для добавления студента
func (c *Classroom) AddStudent(name string, age int) {
	c.students = append(c.students, Student2{name: name, age: age})
}

// AverageAge Метод для вычисления средней возрастной группы
func (c Classroom) AverageAge() float64 {
	totalAge := 0
	for _, student := range c.students {
		totalAge += student.age
	}
	return float64(totalAge) / float64(len(c.students))
}

// Склад товаров
// Описание: Реализуйте структуру Product с полями name и quantity. Реализуйте структуру Warehouse
// с срезом продуктов и методы для добавления продукта и получения общего количества товаров на складе.
// Методы:
// AddProduct(name string, quantity int)
// TotalQuantity() int

type Product struct {
	name     string
	quantity int
}

type Warehouse struct {
	products []Product
}

// AddProduct Метод для добавления продукта
func (w *Warehouse) AddProduct(name string, quantity int) {
	w.products = append(w.products, Product{name: name, quantity: quantity})
}

// TotalQuantity Метод для получения общего количества товаров
func (w Warehouse) TotalQuantity() int {
	total := 0
	for _, product := range w.products {
		total += product.quantity
	}
	return total
}

//  Заметки и метки
// Описание: Реализуйте структуру Note с полем content и tags. Реализуйте структуру Notebook с
// срезом заметок и методы для добавления заметки и получения всех заметок с указанной меткой.
// Методы:
// AddNote(content string, tags []string)
// NotesWithTag(tag string) []Note

type Note struct {
	content string
	tags    []string
}

type Notebook struct {
	notes []Note
}

// AddNote Метод для добавления заметки
func (n *Notebook) AddNote(content string, tags []string) {
	n.notes = append(n.notes, Note{content: content, tags: tags})
}

// NotesWithTag Метод для получения всех заметок с указанной меткой
func (n Notebook) NotesWithTag(tag string) []Note {
	var result []Note
	for _, note := range n.notes {
		for _, t := range note.tags {
			if t == tag {
				result = append(result, note)
				break
			}
		}
	}
	return result
}

// Управление зарплатами
// Описание: Реализуйте структуру Employee с полями name и salary. Реализуйте структуру Payroll
// с срезом сотрудников и методы для добавления сотрудника и получения средней зарплаты.
// Методы:
// AddEmployee(name string, salary float64)
// AverageSalary() float64

type Employee struct {
	name   string
	salary float64
}

type Payroll struct {
	employees []Employee
}

// AddEmployee Метод для добавления сотрудника
func (p *Payroll) AddEmployee(name string, salary float64) {
	p.employees = append(p.employees, Employee{name: name, salary: salary})
}

// AverageSalary Метод для вычисления средней зарплаты
func (p Payroll) AverageSalary() float64 {
	totalSalary := 0.0
	for _, employee := range p.employees {
		totalSalary += employee.salary
	}
	return totalSalary / float64(len(p.employees))
}

// Разделение бюджета
// Описание: Реализуйте структуру Budget с полем amount. Реализуйте метод Allocate, который
// выделяет сумму из бюджета, если сумма доступна. Реализуйте метод Remaining для получения
// оставшегося бюджета.
// Методы:
// Allocate(amount float64) bool
// Remaining() float64

type Budget struct {
	amount float64
}

// Allocate Метод для выделения суммы из бюджета
func (b *Budget) Allocate(amount float64) bool {
	if amount > b.amount {
		return false
	}
	b.amount -= amount
	return true
}

// Remaining Метод для получения оставшегося бюджета
func (b Budget) Remaining() float64 {
	return b.amount
}

// Отслеживание заказов
// Описание: Реализуйте структуру Order с полями id и totalAmount. Реализуйте структуру OrderManager
// с срезом заказов и методы для добавления заказа и получения общей суммы всех заказов.
// Методы:
// AddOrder(id int, totalAmount float64)
// TotalOrdersAmount() float64

type Order struct {
	id          int
	totalAmount float64
}

type OrderManager struct {
	orders []Order
}

// AddOrder Метод для добавления заказа
func (o *OrderManager) AddOrder(id int, totalAmount float64) {
	o.orders = append(o.orders, Order{id: id, totalAmount: totalAmount})
}

// TotalOrdersAmount Метод для получения общей суммы всех заказов
func (o OrderManager) TotalOrdersAmount() float64 {
	total := 0.0
	for _, order := range o.orders {
		total += order.totalAmount
	}
	return total
}

// Управление аккаунтами
// Описание: Реализуйте структуру Account с полем balance. Реализуйте методы Deposit и Withdraw,
// а также метод GetBalance для получения текущего баланса.
// Методы:
// Deposit(amount float64)
// Withdraw(amount float64) bool
// GetBalance() float64

type Account2 struct {
	balance float64
}

// Deposit Метод для пополнения баланса
func (a *Account2) Deposit(amount float64) {
	a.balance += amount
}

// Withdraw Метод для снятия средств с баланса
func (a *Account2) Withdraw(amount float64) bool {
	if amount > a.balance {
		return false
	}
	a.balance -= amount
	return true
}

// GetBalance Метод для получения текущего баланса
func (a Account2) GetBalance() float64 {
	return a.balance
}

// Операции над строками
// Описание: Реализуйте структуру Text с полем content. Реализуйте методы для добавления текста,
// замены слова и получения длины текста.
// Методы:
// AddText(text string)
// ReplaceWord(oldWord, newWord string)
// Length() int

type Text struct {
	content string
}

// AddText Метод для добавления текста
func (t *Text) AddText(text string) {
	t.content += text
}

// ReplaceWord Метод для замены слова
func (t *Text) ReplaceWord(oldWord, newWord string) {
	t.content = strings.ReplaceAll(t.content, oldWord, newWord)
}

// Length Метод для получения длины текста
func (t Text) Length() int {
	return len(t.content)
}

// Управление списком покупок
// Описание: Реализуйте структуру ShoppingItem с полями name и price. Реализуйте структуру ShoppingList с
// срезом покупок и методы для добавления покупки и получения общего значения покупок.
// Методы:
// AddItem(name string, price float64)
// TotalPrice() float64

type ShoppingItem struct {
	name  string
	price float64
}

type ShoppingList struct {
	items []ShoppingItem
}

// AddItem Метод для добавления покупки
func (s *ShoppingList) AddItem(name string, price float64) {
	s.items = append(s.items, ShoppingItem{name: name, price: price})
}

// TotalPrice Метод для получения общего значения покупок
func (s ShoppingList) TotalPrice() float64 {
	total := 0.0
	for _, item := range s.items {
		total += item.price
	}
	return total
}

// Учет времени
// Описание: Реализуйте структуру Event с полями name и date (в формате строк). Реализуйте методы для
// добавления события и получения всех событий после указанной даты.
// Методы:
// AddEvent(name string, date string)
// EventsAfterDate(date string) []Event

type Event struct {
	name string
	date string
}

type Calendar struct {
	events []Event
}

// AddEvent Метод для добавления события
func (c *Calendar) AddEvent(name, date string) {
	c.events = append(c.events, Event{name: name, date: date})
}

// EventsAfterDate Метод для получения всех событий после указанной даты
func (c Calendar) EventsAfterDate(date string) []Event {
	var result []Event
	referenceDate, _ := time.Parse("2006-01-02", date)
	for _, event := range c.events {
		eventDate, _ := time.Parse("2006-01-02", event.date)
		if eventDate.After(referenceDate) {
			result = append(result, event)
		}
	}
	return result
}

// Управление заказами в магазине
// Описание: Реализуйте структуру Order с полями orderID, product и quantity. Реализуйте структуру
// Store с срезом заказов и методы для добавления заказа и получения общего количества товаров по
// каждому продукту.
// Методы:
// AddOrder(orderID int, product string, quantity int)
// TotalQuantityByProduct(product string) int

type Order2 struct {
	orderID  int
	product  string
	quantity int
}

type Store struct {
	orders []Order2
}

// AddOrder Метод для добавления заказа
func (s *Store) AddOrder(orderID int, product string, quantity int) {
	s.orders = append(s.orders, Order2{orderID: orderID, product: product, quantity: quantity})
}

// TotalQuantityByProduct Метод для получения общего количества товаров по продукту
func (s Store) TotalQuantityByProduct(product string) int {
	total := 0
	for _, order := range s.orders {
		if order.product == product {
			total += order.quantity
		}
	}
	return total
}

// Расчет стоимости поездки
// Описание: Реализуйте структуру Trip с полями distance и costPerMile. Реализуйте методы для
// установки стоимости за милю и расчета общей стоимости поездки.
// Методы:
// SetCostPerMile(cost float64)
// TotalCost() float64

type Trip struct {
	distance    float64
	costPerMile float64
}

// SetCostPerMile Метод для установки стоимости за милю
func (t *Trip) SetCostPerMile(cost float64) {
	t.costPerMile = cost
}

// TotalCost Метод для расчета общей стоимости поездки
func (t Trip) TotalCost() float64 {
	return t.distance * t.costPerMile
}

// Рейтинг фильмов
// Описание: Реализуйте структуру Movie с полями title и rating. Реализуйте структуру MovieList
// с срезом фильмов и методы для добавления фильма и получения среднего рейтинга.
// Методы:
// AddMovie(title string, rating float64)
// AverageRating() float64

type Movie struct {
	title  string
	rating float64
}

type MovieList struct {
	movies []Movie
}

// AddMovie Метод для добавления фильма
func (ml *MovieList) AddMovie(title string, rating float64) {
	ml.movies = append(ml.movies, Movie{title: title, rating: rating})
}

// AverageRating Метод для получения среднего рейтинга
func (ml MovieList) AverageRating() float64 {
	totalRating := 0.0
	for _, movie := range ml.movies {
		totalRating += movie.rating
	}
	return totalRating / float64(len(ml.movies))
}

// Система учёта заказов с учетом скидок
// Описание: Реализуйте структуру Order с полями id, product, quantity и price. Реализуйте структуру
// Discount с полем percentage. Реализуйте методы для применения скидки и получения общей суммы
// заказа с учетом скидки.
// Методы:
// ApplyDiscount(discount Discount)
// TotalAmount() float64

type Order3 struct {
	id       int
	product  string
	quantity int
	price    float64
}

type Discount struct {
	percentage float64
}

func (o *Order3) ApplyDiscount(d Discount) {
	o.price -= o.price * (d.percentage / 100)
}

func (o Order3) TotalAmount() float64 {
	return o.price * float64(o.quantity)
}

// Многократные переводы денег
// Описание: Реализуйте структуру Account с полями balance и history (срез транзакций).
// Реализуйте методы для пополнения, снятия и получения истории транзакций.
// Методы:
// Deposit(amount float64)
// Withdraw(amount float64) bool
// TransactionHistory() []string

type Account3 struct {
	balance float64
	history []string
}

// Deposit Метод для пополнения баланса
func (a *Account3) Deposit(amount float64) {
	a.balance += amount
	a.history = append(a.history, fmt.Sprintf("Deposited: %.2f", amount))
}

// Withdraw Метод для снятия средств
func (a *Account3) Withdraw(amount float64) bool {
	if amount > a.balance {
		return false
	}
	a.balance -= amount
	a.history = append(a.history, fmt.Sprintf("Withdrew: %.2f", amount))
	return true
}

// TransactionHistory Метод для получения истории транзакций
func (a Account3) TransactionHistory() []string {
	return a.history
}

// Статистика по продажам
// Описание: Реализуйте структуру Sale с полями item и amount. Реализуйте структуру SalesReport с
// методами для добавления продажи и получения статистики по общим продажам.
// Методы:
// AddSale(item string, amount float64)
// TotalSales() float64

type Sale struct {
	item   string
	amount float64
}

type SalesReport struct {
	sales []Sale
}

// AddSale Метод для добавления продажи
func (sr *SalesReport) AddSale(item string, amount float64) {
	sr.sales = append(sr.sales, Sale{item: item, amount: amount})
}

// TotalSales Метод для получения статистики по общим продажам
func (sr SalesReport) TotalSales() float64 {
	total := 0.0
	for _, sale := range sr.sales {
		total += sale.amount
	}
	return total
}

//  Учет кредитов и дебетов
// Описание: Реализуйте структуру Transaction с полями transactionType (кредит или дебет) и amount.
// Реализуйте структуру Account с методами для добавления транзакций и получения баланса.
// Методы:
// AddTransaction(transactionType string, amount float64)
// Balance() float64

type Transaction2 struct {
	transactionType string
	amount          float64
}

type Account4 struct {
	balance      float64
	transactions []Transaction2
}

// AddTransaction Метод для добавления транзакций
func (a *Account4) AddTransaction(transactionType string, amount float64) {
	if transactionType == "credit" {
		a.balance += amount
	} else if transactionType == "debit" {
		a.balance -= amount
	}
	a.transactions = append(a.transactions, Transaction2{transactionType: transactionType, amount: amount})
}

// Balance Метод для получения баланса
func (a Account4) Balance() float64 {
	return a.balance
}

// Учет заказов с учетом возвратов
// Описание: Реализуйте структуру Order с полями id, product, quantity и price. Реализуйте структуру
// OrderManager с методами для добавления заказа, обработки возвратов и получения общего значения
// активных заказов.
// Методы:
// AddOrder(id int, product string, quantity int, price float64)
// ReturnOrder(id int, quantity int)
// TotalActiveOrders() float64

type Order4 struct {
	id       int
	product  string
	quantity int
	price    float64
}

type OrderManager4 struct {
	orders map[int]Order4
}

// AddOrder Метод для добавления заказа
func (om *OrderManager4) AddOrder(id int, product string, quantity int, price float64) {
	om.orders[id] = Order4{id: id, product: product, quantity: quantity, price: price}
}

// ReturnOrder Метод для обработки возвратов
func (om *OrderManager4) ReturnOrder(id int, quantity int) {
	if order, exists := om.orders[id]; exists {
		if quantity >= order.quantity {
			delete(om.orders, id)
		} else {
			om.orders[id] = Order4{id: id, product: order.product, quantity: order.quantity - quantity, price: order.price}
		}
	}
}

// TotalActiveOrders Метод для получения общего значения активных заказов
func (om OrderManager4) TotalActiveOrders() float64 {
	total := 0.0
	for _, order := range om.orders {
		total += float64(order.quantity) * order.price
	}
	return total
}

// Управление проектами
// Описание: Реализуйте структуру Task с полями title, description и status. Реализуйте структуру Project
// с срезом задач и методы для добавления задачи и получения количества задач в каждом статусе.
// Методы:
// AddTask(title, description, status string)
// CountTasksByStatus(status string) int

type Task2 struct {
	title       string
	description string
	status      string
}

type Project struct {
	tasks []Task2
}

// AddTask Метод для добавления задачи
func (p *Project) AddTask(title, description, status string) {
	p.tasks = append(p.tasks, Task2{title: title, description: description, status: status})
}

// CountTasksByStatus Метод для получения количества задач в каждом статусе
func (p Project) CountTasksByStatus(status string) int {
	count := 0
	for _, task := range p.tasks {
		if task.status == status {
			count++
		}
	}
	return count
}

// Управление пользовательскими данными
// Описание: Реализуйте структуру User с полями username, email и age. Реализуйте структуру UserManager
// с срезом пользователей и методы для добавления пользователя и получения всех пользователей старше
// определенного возраста.
// Методы:
// AddUser(username, email string, age int)
// UsersOlderThan(age int) []User

type User struct {
	username string
	email    string
	age      int
}

type UserManager struct {
	users []User
}

// AddUser Метод для добавления пользователя
func (um *UserManager) AddUser(username, email string, age int) {
	um.users = append(um.users, User{username: username, email: email, age: age})
}

// UsersOlderThan Метод для получения всех пользователей старше определенного возраста
func (um UserManager) UsersOlderThan(age int) []User {
	var result []User
	for _, user := range um.users {
		if user.age > age {
			result = append(result, user)
		}
	}
	return result
}

// Управление контрактами
// Описание: Реализуйте структуру Contract с полями contractID, client и amount. Реализуйте структуру
// ContractManager с срезом контрактов и методы для добавления контракта и получения общего объема
// контрактов для клиента.
// Методы:
// AddContract(contractID int, client string, amount float64)
// TotalAmountForClient(client string) float64

type Contract struct {
	contractID int
	client     string
	amount     float64
}

type ContractManager struct {
	contracts []Contract
}

// AddContract Метод для добавления контракта
func (cm *ContractManager) AddContract(contractID int, client string, amount float64) {
	cm.contracts = append(cm.contracts, Contract{contractID: contractID, client: client, amount: amount})
}

// TotalAmountForClient Метод для получения общего объема контрактов для клиента
func (cm ContractManager) TotalAmountForClient(client string) float64 {
	total := 0.0
	for _, contract := range cm.contracts {
		if contract.client == client {
			total += contract.amount
		}
	}
	return total
}

// Управление списком книг
// Описание: Реализуйте структуру Book с полями title, author и year. Реализуйте структуру Library с
// срезом книг и методы для добавления книги и получения всех книг авторов, опубликованных после
// определенного года.
// Методы:
// AddBook(title, author string, year int)
// BooksByAuthorsAfterYear(year int) []Book

type Book3 struct {
	title  string
	author string
	year   int
}

type Library3 struct {
	books []Book3
}

// AddBook Метод для добавления книги
func (l *Library3) AddBook(title, author string, year int) {
	l.books = append(l.books, Book3{title: title, author: author, year: year})
}

// BooksByAuthorsAfterYear Метод для получения всех книг авторов, опубликованных после определенного года
func (l Library3) BooksByAuthorsAfterYear(year int) []Book3 {
	var result []Book3
	for _, book := range l.books {
		if book.year > year {
			result = append(result, book)
		}
	}
	return result
}

// Отслеживание активностей пользователя
// Описание: Реализуйте структуру Activity с полями activityType и timestamp. Реализуйте структуру
// UserActivityTracker с срезом активностей и методы для добавления активности и получения всех
// активностей после определенного времени.
// Методы:
// AddActivity(activityType string, timestamp time.Time)
// ActivitiesAfterTime(timestamp time.Time) []Activity

type Activity struct {
	activityType string
	timestamp    time.Time
}

type UserActivityTracker struct {
	activities []Activity
}

// AddActivity Метод для добавления активности
func (u *UserActivityTracker) AddActivity(activityType string, timestamp time.Time) {
	u.activities = append(u.activities, Activity{activityType: activityType, timestamp: timestamp})
}

// ActivitiesAfterTime Метод для получения всех активностей после определенного времени
func (u UserActivityTracker) ActivitiesAfterTime(timestamp time.Time) []Activity {
	var result []Activity
	for _, activity := range u.activities {
		if activity.timestamp.After(timestamp) {
			result = append(result, activity)
		}
	}
	return result
}
