package main

import "fmt"

type Counter struct {
	value int
}

// Incrementer
// Метод, возвращающий функцию-замыкание
func (c Counter) Incrementer() func() int {
	c.value++
	return func() int {
		c.value++
		return c.value
	}
}

func main() {
	counter := Counter{}
	inc := counter.Incrementer() // func()int
	fmt.Println(counter.value)

	fmt.Println(inc()) // Output: 1
	fmt.Println(inc()) // Output: 2
	fmt.Println(inc()) // Output: 3
}
