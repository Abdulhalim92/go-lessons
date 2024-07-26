package main

import "fmt"

// Простой узел
// Создайте структуру Node, которая содержит целое значение и указатель на следующий узел.
// Напишите функцию, которая создает два связанных узла и выводит их значения.

type Node struct {
	Value int
	Next  *Node
}

// Функция для создания и вывода связанных узлов
func createAndPrintNodes() {
	node1 := &Node{Value: 1}
	node2 := &Node{Value: 2}
	node1.Next = node2

	fmt.Println("Node 1:", node1.Value)
	fmt.Println("Node 2:", node1.Next.Value)
}

// Три связанных узла
// Создайте структуру Node, которая содержит целое значение и указатель на следующий узел.
// Напишите функцию, которая создает три связанных узла и выводит их значения.

// Функция для создания и вывода трех связанных узлов
func createAndPrintThreeNodes() {
	node1 := &Node{Value: 1}
	node2 := &Node{Value: 2}
	node3 := &Node{Value: 3}
	node1.Next = node2
	node2.Next = node3

	fmt.Println("Node 1:", node1.Value)
	fmt.Println("Node 2:", node1.Next.Value)
	fmt.Println("Node 3:", node1.Next.Next.Value)
}

// Обход связного списка
// Создайте структуру Node, которая содержит целое значение и указатель на следующий узел.
// Напишите функцию, которая принимает указатель на первый узел и обходит связный список,
// выводя значения всех узлов.

// Функция для обхода связного списка и вывода значений узлов
func traverseAndPrintList(head *Node) {
	current := head
	for current != nil {
		fmt.Println(current.Value)
		current = current.Next
	}
}

// Добавление узла в конец списка
// Создайте структуру Node, которая содержит целое значение и указатель на следующий узел.
// Напишите функцию, которая принимает указатель на первый узел и добавляет новый узел в конец списка.

// Функция для добавления нового узла в конец списка
func addNodeToEnd(head *Node, newValue int) {
	current := head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &Node{Value: newValue}
}

// Удаление узла из списка
// Создайте структуру Node, которая содержит целое значение и указатель на следующий узел.
// Напишите функцию, которая принимает указатель на первый узел и значение, которое нужно
// удалить, и удаляет узел с этим значением из списка.

// Функция для удаления узла из списка
func removeNode(head **Node, valueToRemove int) {
	current := *head
	if current != nil && current.Value == valueToRemove {
		*head = current.Next
		return
	}

	for current != nil && current.Next != nil {
		if current.Next.Value == valueToRemove {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
}
