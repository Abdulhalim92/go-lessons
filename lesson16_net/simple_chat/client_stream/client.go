package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func handleClient(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Клиент подключен.")

	// Создание ридера для чтения сообщений из консоли
	serverReader := bufio.NewReader(os.Stdin)

	for {
		// Чтение сообщения от клиента
		clientMessage, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка чтения от клиента:", err)
			return
		}
		fmt.Print("Сообщение от клиента: ", clientMessage)

		// Чтение ввода от сервера
		fmt.Print("Вы (сервер): ")
		serverMessage, err := serverReader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка чтения ввода сервера:", err)
			return
		}

		// Отправка сообщения клиенту
		_, err = conn.Write([]byte(serverMessage))
		if err != nil {
			fmt.Println("Ошибка отправки клиенту:", err)
			return
		}
	}
}

func main() {
	// Запуск сервера на порту 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("Сервер запущен, ожидает подключения...")

	for {
		// Ожидание подключения клиента
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Ошибка при подключении:", err)
			continue
		}

		// Запуск новой goroutine для обработки клиента
		go handleClient(conn)
	}
}
