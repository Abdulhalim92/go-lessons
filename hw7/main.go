package main

import "fmt"

func main() {
	namedTypes()
	funcAsType()
	aliases()
	structs()
	ptrToStruct()
	nestedStructs()
	relatedLists()
	anonymousStructs()
}

func namedTypes() {
	var temp Temperature = -3.5
	result := checkTemperature(temp)
	fmt.Println(result)

	var age Age = 16
	result2 := isTeenager(age)
	fmt.Println(result2)

	var speed Speed = 130.0
	result3 := isValidSpeed(speed)
	fmt.Println(result3)

	var score Score = -5
	result4 := checkScore(score)
	fmt.Println(result4)

	var bmi BMI = 22.0
	result5 := classifyBMI(bmi)
	fmt.Println(result5)
}

func funcAsType() {
	var op UnaryOperation
	op = square
	fmt.Println("Square:", op(4))

	op = double
	fmt.Println("Double:", op(3))

	var check Check
	check = isEven
	fmt.Println("Is Even:", check(4))
	fmt.Println("Is Even:", check(7))

	check = isPositive
	fmt.Println("Is Positive:", check(3))
	fmt.Println("Is Positive:", check(-1))

	results := applyOperations(6, 2, add, subtract, multiply)
	fmt.Println("Results:", results)
}

func aliases() {
	var c Count = 5
	countdown(c)

	var level BatteryLevel = 45
	result := checkBatteryLevel(level)
	fmt.Println(result)

	var w Weight = 75.0
	result2 := categorizeWeight(w)
	fmt.Println(result2)

	var g Grade = 65
	result3 := evaluateGrade(g)
	fmt.Println(result3)

	var bmi BMI = 22.0
	var bp BloodPressure = 110.0
	result4 := evaluateHealth(bmi, bp)
	fmt.Println(result4)
}

func structs() {
	product := Product{Name: "Laptop", Price: 999.99}
	describeProduct(product)

	person := Person{FirstName: "John", LastName: "Doe", Age: 30}
	describePerson(person)

	product1 := Product{Name: "Laptop", Price: 999.99}
	product2 := Product{Name: "Tablet", Price: 499.99}
	expensiveProduct := compareProducts(product1, product2)
	fmt.Printf("More expensive product: %s, Price: %.2f\n", expensiveProduct.Name, expensiveProduct.Price)

	items := []Item{
		{Name: "Apples", Quantity: 10},
		{Name: "Bananas", Quantity: 5},
		{Name: "Oranges", Quantity: 8},
	}
	total := totalQuantity(items)
	fmt.Println("Total quantity:", total)

	books := []Book{
		{Title: "1984", Author: "George Orwell"},
		{Title: "Brave New World", Author: "Aldous Huxley"},
	}
	library := Library{Name: "City Library", Books: books}
	describeLibrary(library)
}

func ptrToStruct() {
	product := &Product{Name: "Laptop", Price: 999.99}
	fmt.Printf("Before update: %s, Price: %.2f\n", product.Name, product.Price)
	updatePrice(product, 899.99)
	fmt.Printf("After update: %s, Price: %.2f\n", product.Name, product.Price)

	person := &Person{FirstName: "John", LastName: "Doe", Age: 30}
	fmt.Printf("Before increase: %s %s, Age: %d\n", person.FirstName, person.LastName, person.Age)
	increaseAge(person)
	fmt.Printf("After increase: %s %s, Age: %d\n", person.FirstName, person.LastName, person.Age)

	product2 := &Product{Name: "Laptop", Price: 999.99}
	fmt.Printf("Before update: %s, Price: %.2f\n", product2.Name, product2.Price)
	updateProduct(product, "Gaming Laptop", 1299.99)
	fmt.Printf("After update: %s, Price: %.2f\n", product2.Name, product2.Price)

	item := &Item{Name: "Apples", Quantity: 10}
	fmt.Printf("Before increase: %s, Quantity: %d\n", item.Name, item.Quantity)
	increaseQuantity(item, 5)
	fmt.Printf("After increase: %s, Quantity: %d\n", item.Name, item.Quantity)

	tasks := []Task{
		{Title: "Task 1", Completed: false},
		{Title: "Task 2", Completed: false},
	}
	fmt.Println("Before update:", tasks)
	updateTaskStatus(&tasks, 1, true)
	fmt.Println("After update:", tasks)
}

func nestedStructs() {
	person := Person2{Name: "John", Address: Address{Street: "123 Main St", City: "New York"}}
	describePerson2(person)

	company := Company{Name: "Tech Corp", Location: Address{Street: "456 Market St", City: "San Francisco"}}
	describeCompany(company)

	instructor := Person3{Name: "Dr. Smith", Address: Address{Street: "789 Elm St", City: "Los Angeles"}}
	course := Course{Title: "Go Programming", Instructor: instructor}
	describeCourse(course)

	event := Event{Title: "Go Conference", Location: Address{Street: "123 Conference St", City: "Boston"}}
	describeEvent(event)

	manager := Person3{Name: "Alice", Address: Address{Street: "123 Manager St", City: "Seattle"}}
	project := Project{Name: "New App", Manager: manager}
	describeProject(project)
}

func relatedLists() {
	createAndPrintNodes()

	createAndPrintThreeNodes()

	node1 := &Node{Value: 1}
	node2 := &Node{Value: 2}
	node3 := &Node{Value: 3}
	node1.Next = node2
	node2.Next = node3
	traverseAndPrintList(node1)

	node11 := &Node{Value: 1}
	node22 := &Node{Value: 2}
	node11.Next = node22
	addNodeToEnd(node11, 3)
	traverseAndPrintList(node11)

	node111 := &Node{Value: 1}
	node222 := &Node{Value: 2}
	node333 := &Node{Value: 3}
	node111.Next = node222
	node222.Next = node333

	fmt.Println("Before removal:")
	traverseAndPrintList(node111)
	removeNode(&node111, 2)
	fmt.Println("After removal:")
	traverseAndPrintList(node111)
}

func anonymousStructs() {
	product := struct {
		Name  string
		Price float64
	}{
		Name:  "Laptop",
		Price: 999.99,
	}
	describeProduct2(product)

	person := struct {
		FirstName string
		LastName  string
		Age       int
	}{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}
	describePerson3(person)

	car := struct {
		Make  string
		Model string
		Year  int
	}{
		Make:  "Toyota",
		Model: "Corolla",
		Year:  2020,
	}
	describeCar(car)

	book := struct {
		Title  string
		Author string
	}{
		Title:  "1984",
		Author: "George Orwell",
	}
	describeBook(book)

	event := struct {
		Title    string
		Date     string
		Location struct {
			Street string
			City   string
		}
	}{
		Title: "Go Conference",
		Date:  "2022-10-10",
		Location: struct {
			Street string
			City   string
		}{
			Street: "123 Conference St",
			City:   "Boston",
		},
	}
	describeEvent2(event)
}
