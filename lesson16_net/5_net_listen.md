# Использование `net.Listen` для создания TCP-сервера в Go

`net.Listen` — это функция из стандартной библиотеки Go, которая используется для
создания сервера, способного принимать входящие соединения по определенному протоколу и
на определенном порту. В контексте TCP-серверов эта функция позволяет приложению 
"слушать" определенный порт и ждать входящих соединений от клиентов.

## Основы работы с `net.Listen`

Функция `net.Listen` создает объект типа `net.Listener`, который может принимать 
входящие соединения. Эта функция принимает два аргумента:

1. **network**: строка, определяющая тип сети. Для TCP-соединений используется значение 
`"tcp"`.
2. **address**: строка, определяющая адрес и порт для прослушивания (например, `":8080"`
для всех IP-адресов на порту 8080).

### Пример создания TCP-сервера с использованием `net.Listen`

В этом примере мы создадим простой TCP-сервер, который будет прослушивать порт `8080` и
обрабатывать входящие соединения в отдельных горутинах.

```go
package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	// Создание слушателя на определенном порту
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Ошибка создания слушателя:", err)
		return
	}
	defer ln.Close()

	fmt.Println("TCP-сервер запущен на порту 8080...")

	// Обработка входящих соединений
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Ошибка соединения:", err)
			continue
		}

		// Обработка каждого соединения в отдельной горутине
		go handleConnection(conn)
	}
}

// Функция для обработки входящего соединения
func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Чтение данных из соединения
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка чтения:", err)
		return
	}

	fmt.Printf("Получено сообщение: %s", message)

	// Отправка ответа клиенту
	conn.Write([]byte("Сообщение получено\n"))
}
```

### Как работает пример

1. **Создание слушателя**: Сервер использует `net.Listen("tcp", ":8080")` для создания 
слушателя, который прослушивает все входящие соединения на порту `8080`.

2. **Принятие соединений**: В цикле `for` сервер ждет входящих соединений с помощью
`ln.Accept()`. Когда соединение установлено, функция возвращает объект `net.Conn`, 
представляющий это соединение.

3. **Обработка соединений в отдельных горутинах**: Чтобы сервер мог обрабатывать 
несколько клиентов одновременно, каждое входящее соединение передается в новую горутину с
помощью `go handleConnection(conn)`. Это позволяет серверу не блокироваться на одном 
соединении.

4. **Чтение и запись данных**: Функция `handleConnection` читает данные от клиента с
помощью `bufio.NewReader` и отправляет ответ обратно клиенту.

### Пример TCP-клиента

Для завершения демонстрации создадим TCP-клиента, который будет подключаться к нашему 
серверу и отправлять сообщение.

```go
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Установка соединения с сервером
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Ошибка подключения к серверу:", err)
		return
	}
	defer conn.Close()

	// Чтение сообщения из консоли
	fmt.Print("Введите сообщение: ")
	message, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	// Отправка сообщения на сервер
	fmt.Fprintf(conn, message + "\n")

	// Получение ответа от сервера
	response, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print("Ответ от сервера: " + response)
}
```

### Как работает клиент

1. **Установка соединения**: Клиент использует `net.Dial("tcp", "localhost:8080")` для
подключения к серверу, который слушает порт `8080`.

2. **Отправка и получение данных**: После подключения клиент отправляет сообщение, 
введенное пользователем, и ждет ответа от сервера.

### Заключение

Функция `net.Listen` и связанный с ней интерфейс `net.Listener` в Go предоставляют 
удобный способ для создания серверов, поддерживающих TCP-соединения. Используя эти
инструменты, можно строить масштабируемые и многопоточные сетевые приложения, которые 
обрабатывают множество соединений одновременно.
