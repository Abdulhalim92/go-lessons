package main

import "fmt"

type Address struct {
	City  string
	State string
}

type Person2 struct {
	Name    string
	Age     int
	Address Address
}

func nested() {
	p := Person2{
		Name: "Alice",
		Age:  30,
		Address: Address{
			City:  "New York",
			State: "NY",
		},
	}
	fmt.Println(p)
	fmt.Println(p.Name)
	fmt.Println(p.Age)
	fmt.Println(p.Address.City)
	fmt.Println(p.Address.State)

	var address Address = Address{
		City:  "New York",
		State: "NY",
	}
	var p2 Person2 = Person2{
		Name:    "Bob",
		Age:     40,
		Address: address,
	}
	fmt.Println(p2)
}
