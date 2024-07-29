package main

import (
	"fmt"
)

var a = initializeVariable("a")
var b = initializeVariable("b")

func initializeVariable(name string) string {
	fmt.Println("Initializing variable", name)
	return name
}

func init() {
	fmt.Println("Init function 1")
}

func init() {
	fmt.Println("Init function 2")
}

func main() {
	fmt.Println("Main function")
	fmt.Println("a =", a)
	fmt.Println("b =", b)
}

// Output:
// Initializing variable a
// Initializing variable b
// Init function 1
// Init function 2
// Main function
// a = a
// b = b
