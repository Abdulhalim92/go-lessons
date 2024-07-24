package main

import (
	"crypto/rand"
	"fmt"
	"strings"
)

// GenerateUUID
// Создайте функцию GenerateUUID, которая генерирует и возвращает уникальный идентификатор UUID v4 в виде строки.
func GenerateUUID() string {
	uuid := make([]byte, 16)
	_, err := rand.Read(uuid)
	if err != nil {
		panic(err)
	}
	uuid[8] = uuid[8]&^0xc0 | 0x80
	uuid[6] = uuid[6]&^0xf0 | 0x40

	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
}

// ReverseString
// Создайте функцию ReverseString, которая принимает строку и возвращает ее инвертированную версию.
func ReverseString(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// CountWords
// Создайте функцию CountWords, которая принимает строку и возвращает количество слов в ней.
func CountWords(text string) int {
	words := strings.Fields(text)
	return len(words)
}
