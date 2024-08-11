package main

import (
	"fmt"
	"io"
	"strings"
)

func ioReader() {
	// Создание нового Reader из строки
	r := strings.NewReader("Hello, Reader!")

	// Буфер для чтения данных
	buf := make([]byte, 8)

	// Чтение данных из Reader в буфер
	for {
		n, err := r.Read(buf)
		if err == io.EOF {
			break // Выход из цикла при достижении конца данных
		}
		fmt.Printf("Read %d bytes: %s\n", n, buf[:n])
	}
}
