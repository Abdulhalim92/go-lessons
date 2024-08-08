package main

import (
	"fmt"
)

// Функция для печати информации о фигуре с использованием переключателя типов
func printShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())

	switch v := s.(type) {
	case Circle:
		fmt.Printf("This is a Circle with radius %.2f\n", v.Radius)
	case Rectangle:
		fmt.Printf("This is a Rectangle with width %.2f and height %.2f\n", v.Width, v.Height)
	default:
		fmt.Println("Unknown shape")
	}
}

func typeSwitches() {
	c := Circle{Radius: 5}
	r := Rectangle{Width: 3, Height: 4}

	printShapeInfo(c) // Output: Area: 78.50
	// This is a Circle with radius 5.00

	printShapeInfo(r) // Output: Area: 12.00
	// This is a Rectangle with width 3.00 and height 4.00

}
