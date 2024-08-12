package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Чтение файла построчно

func fileByLine() {
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// Вывод каждой строки на экран
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error during scanning:", err)
	}
}

// Чтение слов из строки

func wordsFromString() {
	text := "Hello, World! Welcome to Go programming."
	scanner := bufio.NewScanner(strings.NewReader(text))

	// Установка функции разбиения на слова
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		// Вывод каждого слова на экран
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error during scanning:", err)
	}
}

// Чтение файла по символам

func fileByCharacter() {
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Установка функции разбиения на символы
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		// Вывод каждого символа на экран
		fmt.Print(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error during scanning:", err)
	}
}

// Ограничение размера буфера

func bufferSizeLimit() {
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Установка максимального размера буфера (например, 128 КБ)
	const maxCapacity = 128 * 1024
	buffer := make([]byte, maxCapacity)
	scanner.Buffer(buffer, maxCapacity)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error during scanning:", err)
	}

}

// scanner.Split(bufio.ScanLines)
// scanner.Split(bufio.ScanWords)
// scanner.Split(bufio.ScanRunes)
// scanner.Split(bufio.ScanRunes)
