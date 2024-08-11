package main

import (
	"fmt"
	"os"
)

func ioWriter() {
	// Открытие или создание файла для записи
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Данные для записи
	data := []byte("Hello, Writer!\n")

	// Запись данных в файл
	n, err := file.Write(data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Printf("Wrote %d bytes to file.\n", n)
}
