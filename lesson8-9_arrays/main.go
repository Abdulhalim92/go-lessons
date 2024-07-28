package main

import "fmt"

func main() {
	// Создание и инициализация массива
	var arr [5]int
	arr2 := [3]string{"a", "b", "c"}
	arr3 := [...]int{1, 2, 3, 4, 5}

	// Доступ к элементам массива
	fmt.Println(arr[0])
	arr[1] = 10

	// Длина и вместимость массива
	fmt.Println(len(arr), cap(arr))

	// Итерация по массиву
	for i, s := range arr2 {
		fmt.Println(i, s)
	}

	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	// Создание среза на основе массива
	slice := arr3[1:4]
	fmt.Println(slice)

	// Длина и вместимость среза
	fmt.Println(len(slice), cap(slice))

	// Создание среза с помощью make
	slice2 := make([]int, 5)
	slice3 := make([]int, 5, 10)

	// Использование функции append
	slice4 := []int{1, 2, 3}
	slice4 = append(slice4, 4, 5)

	// Операция слайсинг
	slice5 := slice4[1:4]

	// Использование функции copy
	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	n := copy(dst, src)

	fmt.Println(slice2, slice3, slice4, slice5, dst, n)

}
