# Создание клиента в Go и выполнение запросов к серверу

В Go стандартная библиотека `net/http` предоставляет мощный инструментарий для создания
HTTP-клиентов и выполнения запросов к серверам. В этом разделе мы рассмотрим, как 
создать HTTP-клиента, выполнять различные типы запросов (GET, POST, PUT, DELETE) и 
обрабатывать ответы.

## Создание клиента и выполнение HTTP-запросов

Для выполнения HTTP-запросов в Go используется функция `http.Get`, `http.Post` и другие
методы, предоставляемые пакетом `net/http`. Эти функции позволяют отправлять запросы к
серверу и получать ответы.

### Выполнение GET-запроса

Метод **GET** используется для запроса данных с сервера без изменения его состояния. В Go
для выполнения GET-запроса используется функция `http.Get`.

**Пример выполнения GET-запроса:**

```go
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// Выполнение GET-запроса к серверу
	resp, err := http.Get("https://api.example.com/data")
	if err != nil {
		fmt.Println("Ошибка запроса:", err)
		return
	}
	defer resp.Body.Close() // Закрытие тела ответа после завершения

	// Чтение тела ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения ответа:", err)
		return
	}

	// Вывод полученных данных
	fmt.Println(string(body))
}
```

### Выполнение POST-запроса

Метод **POST** используется для отправки данных на сервер (например, отправка формы или 
загрузка файла). В Go для выполнения POST-запроса используется функция `http.Post` или
`http.PostForm`.

**Пример выполнения POST-запроса:**

```go
package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {
	// Подготовка данных для отправки
	jsonData := []byte(`{"name": "John", "age": 30}`)

	// Выполнение POST-запроса
	resp, err := http.Post("https://api.example.com/users", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Ошибка запроса:", err)
		return
	}
	defer resp.Body.Close()

	// Чтение и вывод тела ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения ответа:", err)
		return
	}

	fmt.Println(string(body))
}
```

### Выполнение PUT-запроса

Метод **PUT** используется для замены текущего ресурса данными, отправленными в запросе.
В Go для выполнения PUT-запроса можно использовать метод `http.NewRequest` с указанием
метода `PUT` и стандартный HTTP-клиент.

**Пример выполнения PUT-запроса:**

```go
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	client := &http.Client{}
	// Создание нового DELETE-запроса
	req, err := http.NewRequest(http.MethodDelete, "https://api.example.com/users/123", nil)
	if err != nil {
		fmt.Println("Ошибка создания запроса:", err)
		return
	}

	// Выполнение запроса
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка выполнения запроса:", err)
		return
	}
	defer resp.Body.Close()

	// Чтение и вывод тела ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения ответа:", err)
		return
	}

	fmt.Println(string(body))
}
```

### Выполнение DELETE-запроса

Метод **DELETE** используется для удаления ресурса на сервере. Для выполнения
DELETE-запроса в Go также используется метод `http.NewRequest` с указанием метода 
`DELETE`.

**Пример выполнения DELETE-запроса:**

```go
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	client := &http.Client{}
	// Создание нового DELETE-запроса
	req, err := http.NewRequest(http.MethodDelete, "https://api.example.com/users/123", nil)
	if err != nil {
		fmt.Println("Ошибка создания запроса:", err)
		return
	}

	// Выполнение запроса
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка выполнения запроса:", err)
		return
	}
	defer resp.Body.Close()

	// Чтение и вывод тела ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения ответа:", err)
		return
	}

	fmt.Println(string(body))
}
```