package main

import (
	"fmt"
	"time"
)

func main() {
	s := MyString{text: "Hello World from Go"}
	fmt.Println("Length:", s.Length())        // Output: Length: 17
	fmt.Println("Word Count:", s.WordCount()) // Output: Word Count: 4

	f := MyFormatter{text: "Hello World"}
	fmt.Println("Upper:", f.ToUpper()) // Output: Upper: HELLO WORLD
	fmt.Println("Lower:", f.ToLower()) // Output: Lower: hello world

	n := 10
	ip := IntPointer{num: &n}
	ip.Increment()
	fmt.Println("After Increment:", *ip.num) // Output: After Increment: 11
	ip.Decrement()
	fmt.Println("After Decrement:", *ip.num) // Output: After Decrement: 10

	tc := TextCleaner{text: "  Hello World  "}
	fmt.Println("Trimmed:", tc.TrimSpaces())   // Output: Trimmed: Hello World
	fmt.Println("Removed:", tc.RemoveSpaces()) // Output: Removed: HelloWorld

	sj := StringJoiner{strings: []string{"Hello", "World", "from", "Go"}}
	fmt.Println("Concatenated:", sj.Concat()) // Output: Concatenated: HelloWorldfromGo
	fmt.Println("Joined:", sj.Join(" "))      // Output: Joined: Hello World from Go

	te := TextExtractor{text: "HelloWorld"}
	fmt.Println("Substring:", te.Substr(2, 5)) // Output: Substring: lloWo
	fmt.Println("Suffix:", te.Suffix(3))       // Output: Suffix: rld

	tr := TextReverser{text: "Hello World from Go"}
	fmt.Println("Reversed:", tr.Reverse())            // Output: Reversed: oG morf dlroW olleH
	fmt.Println("Words Reversed:", tr.ReverseWords()) // Output: Words Reversed: Go from World Hello

	ti := TextInserter{text: "Hello World from Go"}
	fmt.Println("InsertAt:", ti.InsertAt(6, "Amazing "))                 // Output: InsertAt: Hello Amazing World from Go
	fmt.Println("InsertAfterWord:", ti.InsertAfterWord("World", "Wide")) // Output: InsertAfterWord: Hello WorldWide from Go

	pc := TextPalindromeChecker{text: "madam"}
	fmt.Println("Is Palindrome:", pc.IsPalindrome()) // Output: Is Palindrome: true
	pc.text = "madam in eden im adam"
	fmt.Println("Is Word Palindrome:", pc.IsWordPalindrome()) // Output: Is Word Palindrome: true

	dr := TextDuplicateRemover{text: "Go Go Go is awesome Awesome"}
	fmt.Println("Remove Duplicates:", dr.RemoveDuplicates())                                 // Output: Remove Duplicates: Go is awesome
	fmt.Println("Remove Duplicates Case Insensitive:", dr.RemoveDuplicatesCaseInsensitive()) // Output: Remove Duplicates Case Insensitive: go is awesome

	budget := &MonthlyBudget{}
	budget.AddIncome(5000)
	budget.AddExpense(2000)
	fmt.Println("Current Balance:", budget.GetBalance()) // Output: Current Balance: 3000

	exchange := Exchange{rateToUSD: 1.1, rateToEUR: 0.9}
	amount := 100.0
	fmt.Printf("%.2f EUR to USD: %.2f\n", amount, exchange.ToUSD(amount)) // Output: 100.00 EUR to USD: 110.00
	fmt.Printf("%.2f USD to EUR: %.2f\n", amount, exchange.ToEUR(amount)) // Output: 100.00 USD to EUR: 90.00

	account := &Account{}
	account.Deposit(1000)
	account.Withdraw(500)
	fmt.Println("Current Balance:", account.GetBalance()) // Output: Current Balance: 500

	taxCalc := SimpleTaxCalculator{incomeTaxRate: 20, vatRate: 15}
	fmt.Println("Income Tax on 1000:", taxCalc.CalculateIncomeTax(1000)) // Output: Income Tax on 1000: 200
	fmt.Println("VAT on 500:", taxCalc.CalculateVAT(500))                // Output: VAT on 500: 75

	invoice := &SimpleInvoice{}
	invoice.AddItem("Laptop", 1000)
	invoice.AddItem("Mouse", 50)
	fmt.Println("Total Invoice:", invoice.Total()) // Output: Total Invoice: 1050

	portfolio := &Portfolio{holdings: make(map[string]int)}
	portfolio.Buy("AAPL", 10)
	portfolio.Buy("GOOGL", 5)
	portfolio.Sell("AAPL", 3)
	fmt.Println("Current Holdings:", portfolio.GetHoldings()) // Output: Current Holdings: map[AAPL:7 GOOGL:5]

	atm := &ATMSystem{}
	atm.Deposit(1000)
	atm.Withdraw(500)
	fmt.Println("ATM Balance:", atm.GetBalance()) // Output: ATM Balance: 500

	payroll := &CompanyPayroll{employees: make(map[string]float64)}
	payroll.AddEmployee("Alice", 50000)
	payroll.AddEmployee("Bob", 60000)
	fmt.Println("Alice's Salary:", payroll.GetSalary("Alice")) // Output: Alice's Salary: 50000
	fmt.Println("Total Payroll:", payroll.TotalPayroll())      // Output: Total Payroll: 110000

	library := &BookLibrary{}
	library.AddBook("1984", "George Orwell")
	library.AddBook("To Kill a Mockingbird", "Harper Lee")
	fmt.Println("Books in Library:", library.GetBooks()) // Output: Books in Library: [1984 To Kill a Mockingbird]

	store := &OnlineStore{}
	store.AddOrder("1001", 250.75)
	store.AddOrder("1002", 500.25)
	fmt.Println("Total Sales:", store.GetTotalSales()) // Output: Total Sales: 751

	account2 := &BankAccount2{}
	account2.AddTransaction(500)
	account2.AddTransaction(-200)
	account2.AddTransaction(300)

	for _, s := range account2.GetStatement() {
		fmt.Println(s)
	}
	// Output:
	// Transaction: 500.00, Balance: 500.00
	// Transaction: -200.00, Balance: 300.00
	// Transaction: 300.00, Balance: 600.00

	catalog := &Catalog{products: make(map[string]float64)}
	catalog.AddProduct("Laptop", 1000)
	catalog.AddProduct("Mouse", 25)
	fmt.Println("Price of Laptop:", catalog.GetProductPrice("Laptop")) // Output: Price of Laptop: 1000
	fmt.Println("All Products:", catalog.GetProducts())                // Output: All Products: map[Laptop:1000 Mouse:25]

	restaurant := &Restaurant{orders: make(map[string][]string)}
	restaurant.AddOrder("1001", []string{"Pizza", "Salad"})
	restaurant.AddOrder("1002", []string{"Burger", "Fries"})
	fmt.Println("Order 1001:", restaurant.GetOrderDetails("1001")) // Output: Order 1001: [Pizza Salad]
	fmt.Println("All Orders:", restaurant.GetAllOrders())          // Output: All Orders: map[1001:[Pizza Salad] 1002:[Burger Fries]]

	mortgage := Mortgage{}
	principal := 200000.0
	annualRate := 5.0
	years := 30
	monthlyPayment := mortgage.CalculateMonthlyPayment(principal, annualRate, years)
	totalPayment := mortgage.CalculateTotalPayment(principal, annualRate, years)
	fmt.Printf("Monthly Payment: %.2f\n", monthlyPayment) // Output: Monthly Payment: 1073.64
	fmt.Printf("Total Payment: %.2f\n", totalPayment)     // Output: Total Payment: 386510.35

	bank := &Bank{customers: make(map[string]string)}
	bank.AddCustomer("Alice", "123456")
	bank.AddCustomer("Bob", "654321")
	accountNumber, name := bank.GetCustomerDetails("123456")
	fmt.Printf("Customer: %s, Account: %s\n", name, accountNumber) // Output: Customer: Alice, Account: 123456
	fmt.Println("All Customers:", bank.GetAllCustomers())          // Output: All Customers: map[123456:Alice 654321:Bob]

	workTime := &WorkTime{
		clockInTimes:  make(map[string]time.Time),
		clockOutTimes: make(map[string]time.Time),
	}
	workTime.ClockIn("E001")
	time.Sleep(2 * time.Second)
	workTime.ClockOut("E001")
	fmt.Printf("Working Hours for E001: %.2f\n", workTime.GetWorkingHours("E001")) // Output: Working Hours for E001: 0.00 (approximately)

	university := &University{students: make(map[string]string)}
	university.AddStudent("John Doe", "S001")
	university.AddStudent("Jane Smith", "S002")
	studentID, name := university.GetStudentDetails("S001")
	fmt.Printf("Student: %s, ID: %s\n", name, studentID)      // Output: Student: John Doe, ID: S001
	fmt.Println("All Students:", university.GetAllStudents()) // Output: All Students: map[S001:John Doe S002:Jane Smith]

	enterprise := &Enterprise{expenses: make(map[string]float64)}
	enterprise.AddExpense("Office Supplies", 200)
	enterprise.AddExpense("Travel", 500)
	enterprise.AddExpense("Office Supplies", 150)
	fmt.Println("Total Office Supplies Expenses:", enterprise.GetTotalExpenses("Office Supplies")) // Output: Total Office Supplies Expenses: 350
	fmt.Println("All Expenses:", enterprise.GetAllExpenses())

	institution := &EducationInstitution{courses: make(map[string]int)}
	institution.AddCourse("Mathematics", 5)
	institution.AddCourse("Physics", 4)
	fmt.Println("Credits for Mathematics:", institution.GetCourseCredits("Mathematics")) // Output: Credits for Mathematics: 5
	fmt.Println("All Courses:", institution.GetAllCourses())

	household := &Household{bills: make(map[string]float64)}
	household.AddBill("Electricity", 100)
	household.AddBill("Water", 50)
	household.AddBill("Electricity", 150)
	fmt.Println("Total Electricity Bills:", household.GetTotalBills("Electricity")) // Output: Total Electricity Bills: 250
	fmt.Println("All Bills:", household.GetAllBills())
}
