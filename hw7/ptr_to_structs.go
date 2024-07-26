package main

// Обновление цены продукта
// Создайте структуру Product с полями Name и Price. Напишите функцию, которая принимает
// указатель на Product и обновляет его цену.

// Функция для обновления цены продукта
func updatePrice(p *Product, newPrice float64) {
	p.Price = newPrice
}

// Увеличение возраста
// Создайте структуру Person с полями Name и Age. Напишите функцию, которая принимает
// указатель на Person и увеличивает его возраст на 1.

// Функция для увеличения возраста
func increaseAge(p *Person) {
	p.Age++
}

// Обновление информации о продукте
// Создайте структуру Product с полями Name и Price. Напишите функцию, которая принимает
// указатель на Product и обновляет его название и цену.

// Функция для обновления информации о продукте
func updateProduct(p *Product, newName string, newPrice float64) {
	p.Name = newName
	p.Price = newPrice
}

// Увеличение количества предметов
// Создайте структуру Item с полями Name и Quantity. Напишите функцию, которая принимает
// указатель на Item и увеличивает количество на заданное значение.

// Функция для увеличения количества предметов
func increaseQuantity(item *Item, amount int) {
	item.Quantity += amount
}

// Управление списком задач
// Создайте структуру Task с полями Title и Completed. Напишите функцию, которая принимает
// указатель на массив Task и обновляет статус выполнения задачи по заданному индексу.

type Task struct {
	Title     string
	Completed bool
}

// Функция для обновления статуса выполнения задачи
func updateTaskStatus(tasks *[]Task, index int, status bool) {
	if index >= 0 && index < len(*tasks) {
		(*tasks)[index].Completed = status
	}
}
