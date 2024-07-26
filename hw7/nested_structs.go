package main

import "fmt"

// Адрес человека
// Создайте структуру Address с полями Street и City. Создайте структуру Person с полями
// Name и Address. Напишите функцию, которая принимает Person и выводит его имя и адрес.

type Address struct {
	Street string
	City   string
}

type Person2 struct {
	Name    string
	Address Address
}

// Функция для вывода информации о человеке
func describePerson2(p Person2) {
	fmt.Printf("Name: %s, Address: %s, %s\n", p.Name, p.Address.Street, p.Address.City)
}

// Описание компании
// Создайте структуру Company с полями Name и Location (структура Address). Напишите функцию,
// которая принимает Company и выводит информацию о компании.

type Company struct {
	Name     string
	Location Address
}

// Функция для вывода информации о компании
func describeCompany(c Company) {
	fmt.Printf("Company: %s, Location: %s, %s\n", c.Name, c.Location.Street, c.Location.City)
}

// Описание курса
// Создайте структуру Course с полями Title и Instructor (структура Person). Напишите функцию,
// которая принимает Course и выводит информацию о курсе.

type Person3 struct {
	Name    string
	Address Address
}

type Course struct {
	Title      string
	Instructor Person3
}

// Функция для вывода информации о курсе
func describeCourse(c Course) {
	fmt.Printf("Course: %s, Instructor: %s, %s, %s\n", c.Title, c.Instructor.Name, c.Instructor.Address.Street, c.Instructor.Address.City)
}

// Описание события
// Создайте структуру Event с полями Title и Location (структура Address). Напишите функцию,
// которая принимает Event и выводит информацию о событии.

type Event struct {
	Title    string
	Location Address
}

// Функция для вывода информации о событии
func describeEvent(e Event) {
	fmt.Printf("Event: %s, Location: %s, %s\n", e.Title, e.Location.Street, e.Location.City)
}

// Управление проектом
// Создайте структуру Project с полями Name и Manager (структура Person, где Person имеет поле
// Name и Address (структура Address)). Напишите функцию, которая принимает Project и выводит
// информацию о проекте и менеджере.

type Project struct {
	Name    string
	Manager Person3
}

// Функция для вывода информации о проекте
func describeProject(p Project) {
	fmt.Printf("Project: %s, Manager: %s, %s, %s\n", p.Name, p.Manager.Name, p.Manager.Address.Street, p.Manager.Address.City)
}
