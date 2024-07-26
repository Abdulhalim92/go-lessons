package main

import "fmt"

// Age
// Создаем именованный тип на основе int
type Age int

// Adder
// Определяем тип функции
type Adder func(int, int) int

// Реализуем функцию, соответствующую типу Adder
func add(a int, b int) int {
	return a + b
}

// MyInt
// Создаем псевдоним для типа int
type MyInt = int

// Person
// Определяем структуру
type Person struct {
	Name string
	Age  int
}

func main() {
	var myAge Age = 25
	fmt.Println(myAge)

	var myAdder Adder = add
	result := myAdder(2, 3)
	fmt.Println(result)

	var x MyInt = 10
	fmt.Println(x)

	// Создаем и инициализируем структуру
	// Использование литерала структуры
	p1 := Person{Name: "Alice", Age: 30}
	fmt.Println(p1)
	fmt.Printf("%+v\n", p1)

	// Без указания имен полей (порядок важен)
	p2 := Person{"Bob", 40}
	fmt.Println(p2)

	// Создание структуры с использованием ключевого слова new:
	p3 := new(Person)
	p3.Name = "Charlie"
	p3.Age = 50
	fmt.Println(p3)
	fmt.Println(*p3)

	// Создание структуры через указатель
	p4 := &Person{Name: "David", Age: 60}
	fmt.Println(*p4)

	// Обращение к полям структуры
	fmt.Println(p4.Name)
	fmt.Println(p4.Age)

	// Вложенные структуры
	nested()

	// Хранение ссылки на структуру того же типа
	relation()

	// Анонимные структуры
	person := struct {
		Name string
		Age  int
	}{
		Name: "Alice",
		Age:  30,
	}
	fmt.Println(person)
}

func problems() {
	var myAge Age = 20
	result := checkAdult(myAge)
	fmt.Println(result)

	var myNumber Number = 7
	result2 := isEven(myNumber)
	fmt.Println(result2)

	var myScore Score = 85
	result3 := isValidScore(myScore)
	fmt.Println(result3)

	var op Operation
	op = adding
	fmt.Println("Addition:", op(3, 4))
	op = subtract
	fmt.Println("Subtraction:", op(7, 2))
	op = multiply
	fmt.Println("Multiplication:", op(5, 6))

	var cmp Comparator
	cmp = equal
	fmt.Println("Equal:", cmp(3, 3))
	cmp = greater
	fmt.Println("Greater:", cmp(5, 3))
	cmp = less
	fmt.Println("Less:", cmp(2, 4))

	var op2 UnaryOperation
	op2 = double
	fmt.Println("Double:", op2(4))
	op2 = triple
	fmt.Println("Triple:", op2(3))

	var c Count = 5
	countdown(c)

	var t Temperature = -5.5
	result4 := checkTemperature(t)
	fmt.Println(result4)

	var originalPrice Price = 100.0
	discountedPrice := applyDiscount(originalPrice)
	fmt.Println("Discounted Price:", discountedPrice)

	user := User{Name: "Alice", Age: 30}
	printUser(user)

	product := Product{Name: "Laptop", Price: 999.99}
	printProduct(product)

	book := Book{Title: "1984", Author: "George Orwell", Pages: 328}
	printBook(book)

	p := &Person3{Name: "Alice", Age: 30}
	fmt.Printf("Before update: %s, Age: %d\n", p.Name, p.Age)
	updateAge(p, 35)
	fmt.Printf("After update: %s, Age: %d\n", p.Name, p.Age)

	r := &Rectangle{Width: 5, Height: 10}
	updateArea(r)
	r.Width = 7
	r.Height = 12
	updateArea(r)

	c1 := &Coordinate{X: 10, Y: 20}
	c2 := &Coordinate{X: 10, Y: 20}
	c3 := &Coordinate{X: 15, Y: 25}
	fmt.Println(compareCoordinates(c1, c2))
	fmt.Println(compareCoordinates(c1, c3))

	student := Student{
		Name: "Bob",
		Age:  22,
		Address: Address2{
			Street: "123 Elm St",
			City:   "Springfield",
		},
	}
	printStudent(student)

	employee := Employee{
		Name: "John",
		ID:   12345,
		Department: Department{
			Name: "Engineering",
		},
	}
	printEmployee(employee)

	manager := &Employee2{Name: "Alice", Position: "Senior Manager"}
	employee2 := &Employee2{Name: "Bob", Position: "Junior Developer", Manager: manager}
	printEmployeeHierarchy(employee2)

	book2 := struct {
		Title  string
		Author string
	}{
		Title:  "The Great Gatsby",
		Author: "F. Scott Fitzgerald",
	}
	printBookInfo(book2)

	product2 := struct {
		Name  string
		Price float64
	}{
		Name:  "Smartphone",
		Price: 499.99,
	}
	description := describeProduct(product2)
	fmt.Println(description)

	event := struct {
		EventName string
		Date      string
	}{
		EventName: "Annual Conference",
		Date:      "2024-08-15",
	}
	description2 := describeEvent(event)
	fmt.Println(description2)
}
