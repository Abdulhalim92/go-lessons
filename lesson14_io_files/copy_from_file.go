package main

import (
	"fmt"
	"io"
	"os"
)

func copyFromFile() {
	// Открытие файла для чтения
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Копирование данных из файла в консоль (os.Stdout)
	_, err = io.Copy(os.Stdout, file)
	if err != nil {
		fmt.Println("Error copying data:", err)
		return
	}
}
