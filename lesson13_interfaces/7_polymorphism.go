package main

import (
	"fmt"
)

// Shape
// Интерфейс Shape с методом Area
type Shape interface {
	Area() float64
}

// Circle
// Тип Circle реализует метод Area
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Rectangle
// Тип Rectangle реализует метод Area
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// TotalArea
// Функция для вычисления общей площади фигур
func TotalArea(shapes ...Shape) float64 {
	var total float64
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

func polymorphism() {
	c := Circle{Radius: 5}
	r := Rectangle{Width: 3, Height: 4}

	fmt.Println("Total Area:", TotalArea(c, r)) // Output: Total Area: 78.5
}
