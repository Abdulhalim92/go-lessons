# Методы структур

### Что такое метод
    Метод в языке Go – это функция, которая связана с определённым типом данных.
    Методы позволяют определять поведение объектов и предоставляют интерфейс для
    взаимодействия с этими объектами. Они объявляются так же, как и функции, но 
    с дополнительным параметром, называемым receiver (приёмник), который
    указывает на тип данных, к которому метод относится.

### Создание метода
    Методы создаются так же, как и функции, с дополнительным параметром receiver.
    Receiver указывается между ключевыми словами func и именем метода. Receiver 
    может быть либо значением (value receiver), либо указателем (pointer receiver).

```go
package main

import "fmt"

// Rectangle 
// Определение структуры
type Rectangle struct {
	width, height float64
}

// Area 
// Метод с value receiver
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Scale 
// Метод с pointer receiver
func (r *Rectangle) Scale(factor float64) {
	r.width *= factor
	r.height *= factor
}

func main() {
	rect := Rectangle{width: 10, height: 5}
	fmt.Println("Area:", rect.Area()) // Output: Area: 50

	// Изменение размера с использованием pointer receiver
	rect.Scale(2)
	fmt.Println("Scaled Area:", rect.Area()) // Output: Scaled Area: 200
}
```

### Value или Pointer receiver

#### Value receiver
    Value receiver использует копию значения receiver. Это значит, что любые 
    изменения, сделанные в методе, не повлияют на исходный объект.

```go
package main

import "fmt"

type Counter struct {
    value int
}

func (c Counter) Increment() {
    c.value++
}

func main() {
    c := Counter{value: 10}
    c.Increment()
    fmt.Println("Value after Increment:", c.value) // Output: Value after Increment: 10
}
```

#### Pointer receiver
    Pointer receiver использует указатель на receiver. Это значит, что метод может
    изменять значения полей receiver, и эти изменения будут видны за пределами 
    метода.

```go
package main

import "fmt"

type Counter struct {
    value int
}

func (c *Counter) Increment() {
    c.value++
}

func main() {
    c := Counter{value: 10}
    c.Increment()
    fmt.Println("Value after Increment:", c.value) // Output: Value after Increment: 11
}
```

### Когда использовать value или pointer receiver
- Если нужно менять значения, то pointer receiver: Если метод должен изменять значения
полей структуры, он должен использовать pointer receiver.
- Много указателей заставляют Garbage Collector (GC) больше работать: Использование
большого количества указателей может увеличивать нагрузку на сборщик мусора.
- Если копия значения большого размера, то лучше использовать указатель: Если 
структура содержит много данных, копирование её значений может быть дорогостоящим.
- Значения пакета sync передавать только по указателю: Пакет sync (например,
sync.Mutex) предназначен для синхронизации и его значения должны передаваться по указателю.

### Ограничения для создания методов
- Тип receiver должен быть определён в том же пакете: Методы можно определять только для
типов, которые объявлены в том же пакете.
- Receiver не может быть интерфейсом: Методы нельзя определять для интерфейсных типов.
- Методы не могут быть перегружены: В Go не поддерживается перегрузка методов,
поэтому в одном типе не может быть двух методов с одинаковыми именами, даже если они
имеют разные параметры.

### Имя метода – модификатор доступа: private, Public
    В Go модификаторы доступа определяются по первому символу имени метода:
- Public: Если имя метода начинается с заглавной буквы, метод является экспортируемым
(публичным) и доступен за пределами пакета.
- private: Если имя метода начинается с маленькой буквы, метод является 
неэкспортируемым (приватным) и доступен только внутри пакета.

```go
package main

import "fmt"

type Rectangle struct {
	width, height float64
}

// Area Public метод
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// private метод
func (r Rectangle) scale(factor float64) {
	r.width *= factor
	r.height *= factor
}

func main() {
	rect := Rectangle{width: 10, height: 5}
	fmt.Println("Area:", rect.Area()) // Output: Area: 50

	// Вызов приватного метода изнутри пакета
	rect.scale(2)
	fmt.Println("Scaled Area:", rect.Area()) // Output: Scaled Area: 200
}
```

### Дополнительно
- Методы могут принимать не только указатели и значения структур, но и любые
пользовательские типы, такие как массивы, срезы, карты и т.д.
- Методы могут иметь различные сигнатуры, включающие параметры и 
возвращаемые значения.
- Методы могут использоваться для реализации интерфейсов, что позволяет 
структурированно и гибко организовать код.

```go
package main

import "fmt"

type MySlice []int

// Sum 
// Метод для пользовательского типа MySlice
func (s MySlice) Sum() int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}

func main() {
	s := MySlice{1, 2, 3, 4, 5}
	fmt.Println("Sum:", s.Sum()) // Output: Sum: 15
}
```

### Методы для встроенных типов
    В Go можно определять методы не только для структур, но и для встроенных
    типов, таких как массивы, срезы, карты и пользовательские типы (alias типов).

```go
package main

import "fmt"

// MyInt 
// Определение нового типа на основе встроенного типа int
type MyInt int

// IsEven 
// Метод для нового типа MyInt
func (m MyInt) IsEven() bool {
	return m%2 == 0
}

func main() {
	num := MyInt(4)
	fmt.Println("Is Even:", num.IsEven()) // Output: Is Even: true
}
```

### Композиция методов
    Методы можно использовать для композиции, когда один метод вызывает другой
    метод.

```go
package main

import "fmt"

type Rectangle struct {
	width, height float64
}

// Area 
// Метод для вычисления площади
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Perimeter 
// Метод для вычисления периметра
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

// Describe 
// Метод для вывода всех характеристик прямоугольника
func (r Rectangle) Describe() {
	fmt.Printf("Width: %v, Height: %v, Area: %v, Perimeter: %v\n", r.width, r.height, r.Area(), r.Perimeter())
}

func main() {
	rect := Rectangle{width: 10, height: 5}
	rect.Describe() // Output: Width: 10, Height: 5, Area: 50, Perimeter: 30
}
```

### Методы и интерфейсы
    Методы играют ключевую роль в реализации интерфейсов. В Go интерфейсы 
    определяют набор методов, и любой тип, который реализует эти методы, 
    автоматически реализует интерфейс.

```go
package main

import "fmt"

// Shape 
// Определение интерфейса
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width, height float64
}

// Area 
// Реализация методов интерфейса Shape для Rectangle
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

type Circle struct {
	radius float64
}

// Area 
// Реализация методов интерфейса Shape для Circle
func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.radius
}

func main() {
	r := Rectangle{width: 10, height: 5}
	c := Circle{radius: 7}

	// Функция, работающая с любым типом, реализующим интерфейс Shape
	describeShape := func(s Shape) {
		fmt.Printf("Area: %v, Perimeter: %v\n", s.Area(), s.Perimeter())
	}

	describeShape(r) // Output: Area: 50, Perimeter: 30
	describeShape(c) // Output: Area: 153.86, Perimeter: 43.96
}
```

### Методы и встраивание (embedding)
    Go поддерживает концепцию встраивания, которая позволяет включать один тип
    в другой. Это полезно для создания композиций и переопределения методов.

```go
package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) Introduce() {
	fmt.Printf("Hi, I'm %s and I'm %d years old.\n", p.name, p.age)
}

type Employee struct {
	Person
	position string
}

// Introduce 
// Переопределение метода Introduce для Employee
func (e Employee) Introduce() {
	fmt.Printf("Hi, I'm %s, a %s, and I'm %d years old.\n", e.name, e.position, e.age)
}

func main() {
	p := Person{name: "Alice", age: 30}
	e := Employee{Person: Person{name: "Bob", age: 40}, position: "Manager"}

	p.Introduce() // Output: Hi, I'm Alice and I'm 30 years old.
	e.Introduce() // Output: Hi, I'm Bob, a Manager, and I'm 40 years old.
}
```

### Методы для срезов
    Методы могут быть определены для срезов, что позволяет добавлять дополнительное
    поведение к коллекциям.

```go
package main

import "fmt"

type IntSlice []int

// Sum 
// Метод для суммирования элементов среза
func (s IntSlice) Sum() int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}

func main() {
	s := IntSlice{1, 2, 3, 4, 5}
	fmt.Println("Sum:", s.Sum()) // Output: Sum: 15
}
```

### Методы с различными сигнатурами
    Методы могут иметь различные сигнатуры, включающие параметры и 
    возвращаемые значения.

```go
package main

import "fmt"

type Calculator struct{}

// Add 
// Метод с двумя параметрами и одним возвращаемым значением
func (c Calculator) Add(a, b int) int {
	return a + b
}

// Divide 
// Метод с двумя параметрами и двумя возвращаемыми значениями
func (c Calculator) Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func main() {
	calc := Calculator{}
	sum := calc.Add(10, 5)
	fmt.Println("Sum:", sum) // Output: Sum: 15

	quotient, err := calc.Divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Quotient:", quotient) // Output: Quotient: 5
	}

	_, err = calc.Divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err) // Output: Error: division by zero
	}
}
```

### Конверсия типов в методах
    Методы могут использоваться для преобразования типов в другие типы (конверсия).

```go
package main

import "fmt"

type Celsius float64
type Fahrenheit float64

func (c Celsius) ToFahrenheit() Fahrenheit {
    return Fahrenheit(c*9/5 + 32)
}

func main() {
    tempC := Celsius(25)
    tempF := tempC.ToFahrenheit()
    fmt.Printf("%.2f°C is %.2f°F\n", tempC, tempF) // Output: 25.00°C is 77.00°F
}
```

### Метод с пустым receiver
    Методы могут быть определены с пустым receiver, что позволяет вызывать метод
    без необходимости создания экземпляра структуры.

```go
package main

import "fmt"

// Logger 
// Пустая структура
type Logger struct{}

// Log 
// Метод с пустым receiver
func (Logger) Log(message string) {
	fmt.Println("Log:", message)
}

func main() {
	Logger{}.Log("This is a log message.") // Output: Log: This is a log message.
}
```

### Методы и интерфейсы с встраиванием
    При встраивании структур, методы встроенных структур также могут быть 
    использованы для удовлетворения интерфейсов.

```go
package main

import "fmt"

type Speaker interface {
    Speak()
}

type Animal struct {
    name string
}

func (a Animal) Speak() {
    fmt.Println(a.name, "says hello!")
}

type Dog struct {
    Animal
    breed string
}

func main() {
    var speaker Speaker
    dog := Dog{Animal: Animal{name: "Buddy"}, breed: "Golden Retriever"}

    speaker = dog
    speaker.Speak() // Output: Buddy says hello!
}
```

### Методы для пользовательских коллекций
    Вы можете создавать методы для пользовательских коллекций, таких как карты 
    и срезы, для добавления специализированного поведения.

```go
package main

import "fmt"

type IntMap map[string]int

// Sum 
// Метод для IntMap
func (m IntMap) Sum() int {
	sum := 0
	for _, v := range m {
		sum += v
	}
	return sum
}

func main() {
	m := IntMap{"a": 1, "b": 2, "c": 3}
	fmt.Println("Sum:", m.Sum()) // Output: Sum: 6
}
```

### Методы с вариативными параметрами
    Методы могут использовать вариативные параметры для обработки переменного
    числа аргументов.

```go
package main

import "fmt"

type Logger struct{}

// Log 
// Метод с вариативными параметрами
func (Logger) Log(messages ...string) {
	for _, message := range messages {
		fmt.Println("Log:", message)
	}
}

func main() {
	logger := Logger{}
	logger.Log("Message 1", "Message 2", "Message 3")   // Output: Log: Message 1
	                                                    //         Log: Message 2
	                                                    //         Log: Message 3
}
```

### Методы и обработка ошибок
    Методы могут возвращать ошибки для обработки исключительных ситуаций.

```go
package main

import (
	"errors"
	"fmt"
)

type Divider struct{}

// Divide 
// Метод для деления с обработкой ошибок
func (Divider) Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	div := Divider{}
	result, err := div.Divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err) // Output: Error: division by zero
	} else {
		fmt.Println("Result:", result)
	}
}
```

### Методы и конструкторы
    В Go нет явной поддержки конструкторов, но вы можете использовать функции
    для инициализации структур.

```go
package main

import "fmt"

type Rectangle struct {
	width, height float64
}

// NewRectangle 
// Функция-конструктор для Rectangle
func NewRectangle(width, height float64) *Rectangle {
	return &Rectangle{width: width, height: height}
}

func main() {
	rect := NewRectangle(10, 5)
	fmt.Printf("Rectangle: %+v\n", rect) // Output: Rectangle: &{width:10 height:5}
}
```

### Методы и замыкания (closures)
    Методы могут использовать замыкания для хранения состояния и создания
    более сложного поведения.

```go
package main

import "fmt"

type Counter struct {
	value int
}

// Incrementer 
// Метод, возвращающий функцию-замыкание
func (c *Counter) Incrementer() func() int {
	return func() int {
		c.value++
		return c.value
	}
}

func main() {
	counter := Counter{}
	inc := counter.Incrementer()

	fmt.Println(inc()) // Output: 1
	fmt.Println(inc()) // Output: 2
	fmt.Println(inc()) // Output: 3
}
```