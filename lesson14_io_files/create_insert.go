package main

import (
	"fmt"
	"os"
)

func creayeInsert() {
	// Создание файла
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Запись строки в файл
	_, err = file.WriteString("Hello, Go!")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("File created and data written successfully")
}
