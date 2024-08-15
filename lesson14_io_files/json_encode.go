package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name  string `json:"name,omitempty"`
	Age   int    `json:"age,omitempty"`
	Email string `json:"email,omitempty"`
}

func jsonEncode() {
	// Создание объекта
	person := Person{
		Name:  "John Doe",
		Age:   30,
		Email: "john.doe@example.com",
	}

	// Открытие файла для записи
	file, err := os.Create("person.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Сериализация объекта в JSON и запись в файл
	encoder := json.NewEncoder(file)
	err = encoder.Encode(person)
	if err != nil {
		fmt.Println("Error encoding JSON to file:", err)
		return
	}

	fmt.Println("JSON data written to file successfully")
}
