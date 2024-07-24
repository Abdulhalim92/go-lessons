package main

import (
	"fmt"
	"math/rand"
	"time"
)

// GenerateRandomPassword
// Создайте функцию GenerateRandomPassword, которая генерирует и возвращает случайный пароль из 8 символов, состоящий из букв и цифр.
func GenerateRandomPassword() {
	rand.Seed(time.Now().UnixNano())

	var password string
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	for i := 0; i < 8; i++ {
		index := rand.Intn(len(characters))
		password += string(characters[index])
	}

	fmt.Println("Generated password:", password)
}

// Создайте функцию GenerateInvoiceNumber, которая генерирует и печатает уникальный номер счета в формате "INV-2024-07-001".
var invoiceCounter int

func GenerateInvoiceNumber() {
	year := time.Now().Year()
	month := time.Now().Month()
	invoiceCounter++

	invoiceNumber := fmt.Sprintf("INV-%d-%02d-%03d", year, month, invoiceCounter)
	fmt.Println("Generated invoice number:", invoiceNumber)
}

// PrintFormattedDate
// Создайте функцию PrintFormattedDate, которая печатает текущую дату в формате "2 January 2006".
func PrintFormattedDate() {
	currentDate := time.Now().Format("2 January 2006")
	fmt.Println("Current date:", currentDate)
}
