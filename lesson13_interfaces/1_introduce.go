package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Person struct {
	Name string
}

func (p Person) Speak() string {
	return "Hello, my name is " + p.Name
}

func Greet(s Speaker) {
	fmt.Println(s.Speak())
}

func introduce() {
	p := Person{Name: "Alice"}
	Greet(p) // Output: Hello, my name is Alice
}
