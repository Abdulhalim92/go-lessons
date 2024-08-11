package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//type Person struct {
//	Name  string `json:"name"`
//	Age   int    `json:"age"`
//	Email string `json:"email"`
//}

func jsonDecode() {
	// Открытие файла для чтения
	file, err := os.Open("person.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var person Person

	// Десериализация JSON-данных из файла в структуру
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&person)
	if err != nil {
		fmt.Println("Error decoding JSON from file:", err)
		return
	}

	fmt.Printf("Name: %s, Age: %d, Email: %s\n", person.Name, person.Age, person.Email)
}
