package main

import "fmt"

func main() {
	var firstVariable bool
	fmt.Println(firstVariable)

	secondVariable := false
	fmt.Println(secondVariable)

	age := 15

	// Basic if
	if (age < 18) == true { // можно короче: if age < 18 {
		fmt.Println("You are too young (full)")
	}

	//Short syntax
	if isChild := isChildren(age); isChild == true { // можно короче: if isChild := isChildren(age); isChild {
		fmt.Println("You are too young (short)")
		fmt.Println(isChild)
	}
	//fmt.Println(isChild)

	// If ... else
	age = 20
	if age < 18 {
		fmt.Println("you are too young")
	} else {
		fmt.Println("You are an adult")
	}

	// &&
	if age >= 7 && age <= 18 {
		fmt.Println("You are in school")
	}

	// ||
	if age == 14 || age == 20 || age == 40 {
		fmt.Println("You have to get a new passport")
	}

	// !
	if !isChildren(age) {
		fmt.Println("You are an adult")
	}

	num := 6
	if num > 9 && num < 100 {
		fmt.Println("Двузначное")
	}

	n := 6
	if n%2 == 1 {
		fmt.Println("Число нечетное")
	} else {
		fmt.Println("Число четное")
	}

	m := 78
	if m >= 0 && m < 10 {
		fmt.Println("Число однозначное")
	} else if m >= 10 && m < 100 {
		fmt.Println("Число двузначное")
	} else if m >= 100 && m < 1000 {
		fmt.Println("Число трехначное")
	} else {
		fmt.Println("Число слишком большое")
	}

	nn := -1
	switch nn {
	case 1:
		fmt.Println("Positive")
		fallthrough
	case -1:
		fmt.Println("Negative")
	case 0:
		fmt.Println("Unsigned")
	}

	if nn == 1 {
		fmt.Println("Positive")
	} else if nn == -1 {
		fmt.Println("Negative")
	} else if n == 0 {
		fmt.Println("Unsigned")
	}

	var a int
	fmt.Scan(&a)
	fmt.Println(a)
}

func isChildren(age int) bool {
	return age < 18
}
