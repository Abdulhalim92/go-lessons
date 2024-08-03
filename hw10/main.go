package main

import (
	"fmt"
	"time"
)

func main() {
	lib := Library{}
	book1 := Book{title: "1984", author: "George Orwell"}
	book2 := Book{title: "To Kill a Mockingbird", author: "Harper Lee"}
	lib.AddBook(book1)
	lib.AddBook(book2)
	for _, book := range lib.books {
		fmt.Println(book.GetDetails())
	}

	student := Student{}
	student.AddGrade(90)
	student.AddGrade(85)
	student.AddGrade(78)
	fmt.Printf("Average Grade: %.2f\n", student.AverageGrade())

	circle := Circle{radius: 5}
	fmt.Printf("Area: %.2f\n", circle.Area())
	fmt.Printf("Circumference: %.2f\n", circle.Circumference())

	inventory := Inventory{}
	inventory.AddItem(Item{name: "Laptop", price: 999.99})
	inventory.AddItem(Item{name: "Phone", price: 599.99})
	fmt.Printf("Total Value: %.2f\n", inventory.TotalValue())

	acc := Account{}
	acc.AddTransaction(150.75, "Deposit")
	acc.AddTransaction(-50.00, "Withdrawal")

	taskList := TaskList{}
	taskList.AddTask("Task 1")
	taskList.AddTask("Task 2")
	taskList.tasks[1].completed = true // Установка статуса задачи как завершённого
	fmt.Printf("Completed Tasks: %d\n", taskList.CompletedTasks())

	temp := Temperature{celsius: 25}
	fmt.Printf("Fahrenheit: %.2f\n", temp.ToFahrenheit())
	fmt.Printf("Kelvin: %.2f\n", temp.ToKelvin())

	classroom := Classroom{}
	classroom.AddStudent("Alice", 20)
	classroom.AddStudent("Bob", 22)
	classroom.AddStudent("Charlie", 19)
	fmt.Printf("Average Age: %.2f\n", classroom.AverageAge())

	warehouse := Warehouse{}
	warehouse.AddProduct("Widget", 100)
	warehouse.AddProduct("Gadget", 200)
	fmt.Printf("Total Quantity: %d\n", warehouse.TotalQuantity())

	notebook := Notebook{}
	notebook.AddNote("Go is awesome", []string{"programming", "go"})
	notebook.AddNote("Learn Go concurrency", []string{"programming", "go", "concurrency"})
	notes := notebook.NotesWithTag("go")
	for _, note := range notes {
		fmt.Println("Note:", note.content)
	}

	payroll := Payroll{}
	payroll.AddEmployee("John", 60000)
	payroll.AddEmployee("Jane", 65000)
	payroll.AddEmployee("Doe", 70000)
	fmt.Printf("Average Salary: %.2f\n", payroll.AverageSalary())

	budget := Budget{amount: 5000}
	if budget.Allocate(2000) {
		fmt.Println("Allocation successful")
	} else {
		fmt.Println("Allocation failed")
	}
	fmt.Printf("Remaining Budget: %.2f\n", budget.Remaining())

	manager := OrderManager{}
	manager.AddOrder(1, 250.75)
	manager.AddOrder(2, 125.50)
	manager.AddOrder(3, 300.00)
	fmt.Printf("Total Orders Amount: %.2f\n", manager.TotalOrdersAmount())

	account := Account2{balance: 1000}
	account.Deposit(500)
	if account.Withdraw(300) {
		fmt.Println("Withdrawal successful")
	} else {
		fmt.Println("Withdrawal failed")
	}
	fmt.Printf("Current Balance: %.2f\n", account.GetBalance())

	text := Text{}
	text.AddText("Hello world! ")
	text.AddText("Welcome to Go programming.")
	text.ReplaceWord("world", "universe")
	fmt.Printf("Text Length: %d\n", text.Length())
	fmt.Println("Text Content:", text.content)

	list := ShoppingList{}
	list.AddItem("Bread", 2.50)
	list.AddItem("Milk", 1.20)
	list.AddItem("Eggs", 3.00)
	fmt.Printf("Total Price: %.2f\n", list.TotalPrice())

	calendar := Calendar{}
	calendar.AddEvent("Go Conference", "2024-08-15")
	calendar.AddEvent("Go Workshop", "2024-09-10")
	events := calendar.EventsAfterDate("2024-08-01")
	for _, event := range events {
		fmt.Printf("Event: %s on %s\n", event.name, event.date)
	}

	store := Store{}
	store.AddOrder(1, "Laptop", 5)
	store.AddOrder(2, "Phone", 3)
	store.AddOrder(3, "Laptop", 2)
	fmt.Printf("Total Quantity of Laptops: %d\n", store.TotalQuantityByProduct("Laptop"))

	trip := Trip{distance: 100}
	trip.SetCostPerMile(0.75)
	fmt.Printf("Total Trip Cost: %.2f\n", trip.TotalCost())

	movieList := MovieList{}
	movieList.AddMovie("Inception", 8.8)
	movieList.AddMovie("The Matrix", 8.7)
	movieList.AddMovie("Interstellar", 8.6)
	fmt.Printf("Average Rating: %.2f\n", movieList.AverageRating())

	order := Order3{id: 1, product: "Laptop", quantity: 2, price: 1000}
	discount := Discount{percentage: 10}
	order.ApplyDiscount(discount)
	fmt.Printf("Total Amount after Discount: %.2f\n", order.TotalAmount())

	account3 := Account3{balance: 1000}
	account3.Deposit(500)
	account3.Withdraw(200)
	history := account3.TransactionHistory()
	for _, record := range history {
		fmt.Println(record)
	}

	report := SalesReport{}
	report.AddSale("TV", 800)
	report.AddSale("Laptop", 1200)
	fmt.Printf("Total Sales: %.2f\n", report.TotalSales())

	account4 := Account4{balance: 1000}
	account4.AddTransaction("credit", 500)
	account4.AddTransaction("debit", 200)
	fmt.Printf("Current Balance: %.2f\n", account4.Balance())

	manager4 := OrderManager4{orders: make(map[int]Order4)}
	manager4.AddOrder(1, "Laptop", 2, 1000)
	manager4.AddOrder(2, "Phone", 5, 500)
	manager4.ReturnOrder(1, 1)
	fmt.Printf("Total Active Orders: %.2f\n", manager4.TotalActiveOrders())

	project := Project{}
	project.AddTask("Task 1", "Description 1", "completed")
	project.AddTask("Task 2", "Description 2", "in progress")
	project.AddTask("Task 3", "Description 3", "completed")
	fmt.Printf("Completed Tasks: %d\n", project.CountTasksByStatus("completed"))

	userManager := UserManager{}
	userManager.AddUser("Alice", "alice@example.com", 30)
	userManager.AddUser("Bob", "bob@example.com", 25)
	userManager.AddUser("Charlie", "charlie@example.com", 35)
	users := userManager.UsersOlderThan(30)
	for _, user := range users {
		fmt.Println("User:", user.username, "Age:", user.age)
	}

	manager2 := ContractManager{}
	manager2.AddContract(1, "Company A", 5000)
	manager2.AddContract(2, "Company B", 3000)
	manager2.AddContract(3, "Company A", 2000)
	fmt.Printf("Total Amount for Company A: %.2f\n", manager2.TotalAmountForClient("Company A"))

	library := Library3{}
	library.AddBook("The Catcher in the Rye", "J.D. Salinger", 1951)
	library.AddBook("To Kill a Mockingbird", "Harper Lee", 1960)
	library.AddBook("1984", "George Orwell", 1949)
	books := library.BooksByAuthorsAfterYear(1950)
	for _, book := range books {
		fmt.Printf("Book: %s by %s (Published: %d)\n", book.title, book.author, book.year)
	}

	tracker := UserActivityTracker{}
	tracker.AddActivity("Login", time.Now().Add(-1*time.Hour))
	tracker.AddActivity("Logout", time.Now().Add(-30*time.Minute))
	tracker.AddActivity("Upload", time.Now())
	activities := tracker.ActivitiesAfterTime(time.Now().Add(-45 * time.Minute))
	for _, activity := range activities {
		fmt.Printf("Activity: %s at %s\n", activity.activityType, activity.timestamp.Format(time.RFC3339))
	}
}
