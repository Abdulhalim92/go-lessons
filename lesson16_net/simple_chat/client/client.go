package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Подключение к серверу
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Ошибка подключения к серверу:", err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println("Подключено к серверу.")

	// Создание ридера для чтения сообщений из консоли
	clientReader := bufio.NewReader(os.Stdin)

	for {
		// Чтение ввода от клиента
		fmt.Print("Вы (клиент): ")
		clientMessage, err := clientReader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка чтения ввода клиента:", err)
			return
		}

		// Отправка сообщения серверу
		_, err = conn.Write([]byte(clientMessage))
		if err != nil {
			fmt.Println("Ошибка отправки серверу:", err)
			return
		}

		// Чтение сообщения от сервера
		serverMessage, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка чтения от сервера:", err)
			return
		}
		fmt.Print("Сообщение от сервера: ", serverMessage)
	}
}
