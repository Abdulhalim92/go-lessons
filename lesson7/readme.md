# Производные типы в Go

### Именованные типы и производные
    Именованные типы в Go позволяют создать новый тип на основе существующего типа. 
    Это полезно для повышения читаемости кода и создания более специфичных типов для 
    определенных целей.

```go
package main

import "fmt"

// Age 
//  Создаем именованный тип на основе int
type Age int

func main() {
	var myAge Age = 25
	fmt.Println(myAge)
}

```

### Функция как тип данных
    В Go функции являются первоклассными гражданами, что означает, что их можно 
    присваивать переменным, передавать в качестве аргументов и возвращать из других функций.

```go
package main

import "fmt"

// Adder 
// Определяем тип функции
type Adder func(int, int) int

// Реализуем функцию, соответствующую типу Adder
func add(a int, b int) int {
	return a + b
}

func main() {
	var myAdder Adder = add
	result := myAdder(2, 3)
	fmt.Println(result)
}

```

### Псевдонимы
    Псевдонимы типов позволяют создать альтернативное имя для существующего типа. 
    Они не создают новый тип, а просто дают другое имя для существующего типа.

```go
package main

import "fmt"

// MyInt 
// Создаем псевдоним для типа int
type MyInt = int

func main() {
	var x MyInt = 10
	fmt.Println(x)
}

```

### Структуры
    Структуры в Go используются для объединения данных разных типов в одну сущность.
    Поля структуры могут быть разного типа, что позволяет хранить связанные данные вместе.

```go
package main

import "fmt"

// Person 
// Определяем структуру
type Person struct {
	Name string
	Age  int
}

func main() {
	// Создаем и инициализируем структуру
	p := Person{Name: "Alice", Age: 30}
	fmt.Println(p)
}

```

### Создание и инициализация структур
    Структуры можно создавать и инициализировать несколькими способами.
- Использование литерала структуры:
```go
p := Person{Name: "Alice", Age: 30}
```
- Без указания имен полей (требуется указать все поля в порядке объявления):
```go
p := Person{"Alice", 30}
```
- Создание структуры с использованием ключевого слова new:
```go
p := new(Person)
p.Name = "Alice"
p.Age = 30
```
- Создание структуры через указатель:
```go
p := &Person{Name: "Alice", Age: 30}
```
- Использование анонимной структуры:
```go
p := struct {
    Name    string
    Age     int
    Married bool
    Salary  float64
}{
    Name:    "Alice",
    Age:     30,
    Married: true,
    Salary:  50000.0,
}
```

# Рекомендации
- Используйте именованную инициализацию полей для повышения читаемости 
и избежания ошибок:
```go
p := Person{
    Name:    "Alice",
    Age:     30,
    Married: true,
    Salary:  50000.0,
}
```
- Для больших или часто используемых структур создавайте функции-конструкторы:
```go
func NewPerson(name string, age int, married bool, salary float64) Person {
    return Person{
        Name:    name,
        Age:     age,
        Married: married,
        Salary:  salary,
    }
}

p := NewPerson("Alice", 30, true, 50000.0)
```

### Обращение к полям структуры
    Обращение к полям структуры осуществляется с помощью оператора точки (.).
```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func main() {
    p := Person{Name: "Alice", Age: 30}
    fmt.Println(p.Name)
    fmt.Println(p.Age)
}
```

### Указатели на структуры
    Указатели на структуры позволяют изменять значения полей структуры без необходимости создания копии структуры.
```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func main() {
    p := &Person{Name: "Alice", Age: 30}
    p.Age = 31
    fmt.Println(p)
}
```

### Вложенные структуры
    Структуры могут содержать другие структуры в качестве полей, что позволяет создавать более сложные данные.
```go
package main

import "fmt"

type Address struct {
    City  string
    State string
}

type Person struct {
    Name    string
    Age     int
    Address Address
}

func main() {
    p := Person{
        Name: "Alice",
        Age:  30,
        Address: Address{
            City:  "New York",
            State: "NY",
        },
    }
    fmt.Println(p)
}
```

### Хранение ссылки на структуру того же типа
    Структуры могут содержать указатели на другие структуры того же типа, что позволяет создавать
    связные структуры данных, такие как деревья или графы.
```go
package main

import "fmt"

type Node struct {
    Value int
    Next  *Node
}

func main() {
    n1 := &Node{Value: 1}
    n2 := &Node{Value: 2}
    n3 := &Node{Value: 3}
    
    n1.Next = n2
    n2.Next = n3

    // Печатаем значения узлов
    current := n1
    for current != nil {
        fmt.Println(current.Value)
        current = current.Next
    }
}
```
    В данном примере мы создаем простую односвязную структуру, где каждый узел (Node)
    содержит значение и указатель на следующий узел. Это позволяет нам хранить ссылки 
    на другие структуры того же типа, создавая цепочку узлов.

### Анонимные структуры
    Анонимные структуры позволяют создавать структуры без явного определения типа. Они полезны для 
    создания временных структур или структур, используемых в ограниченном контексте.
```go
package main

import "fmt"

func main() {
    person := struct {
        Name string
        Age  int
    }{
        Name: "Alice",
        Age:  30,
    }
    fmt.Println(person)
}
```

### Нулевое значение структуры
    В Go дефолтное (нуль) значение для структуры означает, что все ее поля инициализируются нуль 
    значениями соответствующих типов данных. Вот пример структуры с различными 
    типами полей и их дефолтными значениями:
```go
type Person struct {
    Name    string  // дефолтное значение ""
    Age     int     // дефолтное значение 0
    Married bool    // дефолтное значение false
    Salary  float64 // дефолтное значение 0.0
}
```

    Когда структура объявляется без явной инициализации, ее поля принимают нуль значения:
```go
var p Person
fmt.Println(p) // Вывод: { 0 false 0}
```

    Если вы хотите создать структуру с собственными дефолтными значениями, вы можете использовать 
    функцию-конструктор. Эта функция будет возвращать структуру, инициализированную вашими 
    собственными значениями по умолчанию. Вот пример:
```go
package main

import "fmt"

type Person struct {
	Name    string
	Age     int
	Married bool
	Salary  float64
}

// NewPerson 
// Функция-конструктор для создания Person с дефолтными значениями
func NewPerson() Person {
	return Person{
		Name:    "John Doe",
		Age:     25,
		Married: false,
		Salary:  30000.0,
	}
}

func main() {
	p := NewPerson()
	fmt.Println(p) // Вывод: {John Doe 25 false 30000}
}
```