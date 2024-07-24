package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func relation() {
	n1 := &Node{Value: 1}
	n2 := &Node{Value: 2}
	n3 := &Node{Value: 3}

	n1.Next = n2
	n2.Next = n3
	//n3.Next = n1

	// Печатаем значения узлов
	current := n1
	for current != nil {
		fmt.Println(current.Value)
		current = current.Next
	}
}
