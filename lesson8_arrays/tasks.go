package main

import "fmt"

// Создайте массив из 7 целых чисел и инициализируйте его значениями от 1 до 7.
func task1() {
	arr := [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arr) // [1 2 3 4 5 6 7]
}

// Создайте массив из 5 строк и инициализируйте его значениями "a", "b", "c", "d", "e".
func task2() {
	arr := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(arr) // [a b c d e]
}

// Создайте массив из 4 логических значений и инициализируйте его значениями true, false, true, false.
func task3() {
	arr := [4]bool{true, false, true, false}
	fmt.Println(arr) // [true false true false]
}

// Создайте массив из 10 целых чисел без инициализации и выведите его на экран.
func task4() {
	arr := [10]int{}
	fmt.Println(arr) // [0 0 0 0 0 0 0 0 0 0]
}

// Создайте массив из 5 вещественных чисел и инициализируйте его значениями 1.1, 2.2, 3.3, 4.4, 5.5.
func task5() {
	arr := [5]float64{1.1, 2.2, 3.3, 4.4, 5.5}
	fmt.Println(arr) // [1.1 2.2 3.3 4.4 5.5]
}

// Создайте массив из 3 строковых значений с автоматическим определением длины массива.
func task6() {
	arr := [...]string{"a", "b", "c"}
	fmt.Println(arr) // [a b c]
}

// Выведите на экран первый и последний элемент массива [5]int{10, 20, 30, 40, 50}.
func task7() {
	arr := [5]int{10, 20, 30, 40, 50}
	fmt.Println(arr[0], arr[len(arr)-1]) // 10 50
}

// Создайте массив из 4 логических значений и выведите значения по индексам 1 и 3.
func task8() {
	arr := [4]bool{true, false, true, false}
	fmt.Println(arr[1], arr[3]) // false false
}

// Создайте массив из 6 вещественных чисел и выведите значение элемента по индексу 2.
func task9() {
	arr := [6]float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6}
	fmt.Println(arr[2]) // 3.3
}

// Создайте массив из 5 целых чисел и измените значение второго элемента на 100.
func task10() {
	arr := [5]int{1, 2, 3, 4, 5}
	arr[1] = 100
	fmt.Println(arr) // [1 100 3 4 5]
}

// Создайте массив из 4 строк и измените значение последнего элемента на "Go".
func task11() {
	arr := [4]string{"Python", "Java", "C++", "C#"}
	arr[len(arr)-1] = "Go"
	fmt.Println(arr) // [Python Java C++ Go]
}

// Создайте массив из 3 логических значений и измените значение первого элемента на false.
func task12() {
	arr := [3]bool{true, false, true}
	arr[0] = false
	fmt.Println(arr) // [false false true]
}

// Создайте массив из 7 целых чисел и выведите его длину и вместимость.
func task13() {
	arr := [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(len(arr), cap(arr)) // 7 7
}

// Создайте массив из 3 строк и выведите его длину и вместимость.
func task14() {
	arr := [3]string{"a", "b", "c"}
	fmt.Println(len(arr), cap(arr)) // 3 3
}

// Создайте массив из 5 логических значений и выведите его длину и вместимость.
func task15() {
	arr := [5]bool{true, false, true, false, true}
	fmt.Println(len(arr), cap(arr)) // 5 5
}

// Создайте массив из 5 целых чисел и выведите все его элементы с помощью цикла for.
func task16() {
	arr := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

// Создайте массив из 4 строк и выведите все его элементы с помощью цикла for range.
func task17() {
	arr := [4]string{"a", "b", "c", "d"}
	for i, s := range arr {
		fmt.Println(i, s)
	}
}

// Создайте массив из 3 логических значений и выведите все его элементы с помощью цикла for.
func task18() {
	arr := [3]bool{true, false, true}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

// Создайте срез из 5 целых чисел и инициализируйте его значениями от 1 до 5.
func task19() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(arr) // [1 2 3 4 5]
}

// Создайте пустой срез из строк и добавьте в него значения "a", "b", "c".
func task20() {
	arr := []string{}
	arr = append(arr, "a", "b", "c")
	fmt.Println(arr) // [a b c]
}

// Создайте срез из 4 логических значений и инициализируйте его значениями true, false, true, false.
func task21() {
	arr := []bool{true, false, true, false}
	fmt.Println(arr) // [true false true false]
}

// Создайте срез из 6 целых чисел и инициализируйте его значениями от 1 до 6.
func task22() {
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr) // [1 2 3 4 5 6]
}

// Создайте пустой срез из строк и добавьте в него значения "hello", "world".
func task23() {
	arr := []string{}
	arr = append(arr, "hello", "world")
	fmt.Println(arr) // [hello world]
}

// Создайте срез из 5 вещественных чисел и инициализируйте его значениями 1.1, 2.2, 3.3, 4.4, 5.5.
func task24() {
	arr := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	fmt.Println(arr) // [1.1 2.2 3.3 4.4 5.5]
}

// Создайте срез из 7 целых чисел и выведите его длину и вместимость.
func task25() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(len(arr), cap(arr)) // 7 7
}

// Создайте пустой срез из строк и добавьте в него 3 элемента, затем выведите его длину и вместимость.
func task26() {
	arr := []string{}
	arr = append(arr, "a", "b", "c")
	fmt.Println(len(arr), cap(arr)) // 3 3
}

// Создайте срез из 4 вещественных чисел и выведите его длину и вместимость.
func task27() {
	arr := []float64{1.1, 2.2, 3.3, 4.4}
	fmt.Println(len(arr), cap(arr)) // 4 4
}

// Создайте срез из 5 целых чисел с помощью функции make.
func task28() {
	arr := make([]int, 5)
	fmt.Println(arr) // [0 0 0 0 0]
}

// Создайте срез из 3 строк с длиной 3 и вместимостью 5.
func task29() {
	arr := make([]string, 3, 5)
	fmt.Println(arr) // [ 0 0 0 ]
}

// Создайте пустой срез из логических значений с вместимостью 4.
func task30() {
	arr := make([]bool, 0, 4)
	fmt.Println(arr) // []
}

// Создайте срез из 3 целых чисел и добавьте в него значение 4.
func task31() {
	arr := []int{1, 2, 3}
	arr = append(arr, 4)
	fmt.Println(arr) // [1 2 3 4]
}

// Создайте пустой срез из строк и добавьте в него значения "a", "b", "c".
func task32() {
	arr := []string{}
	arr = append(arr, "a", "b", "c")
	fmt.Println(arr) // [a b c]
}

// Создайте срез из 2 вещественных чисел и добавьте в него 3 новых значения.
func task33() {
	arr := []float64{1.1, 2.2}
	arr = append(arr, 3.3, 4.4, 5.5)
	fmt.Println(arr) // [1.1 2.2 3.3 4.4 5.5]
}

// Создайте массив из 6 целых чисел и создайте срез, содержащий элементы с индексами от 1 до 4.
func task34() {
	arr := [6]int{1, 2, 3, 4, 5, 6}
	arr2 := arr[1:5]
	fmt.Println(arr2) // [2 3 4 5]
}

// Создайте срез из 5 строк и создайте новый срез, содержащий последние 3 элемента.
func task35() {
	arr := []string{"a", "b", "c", "d", "e"}
	arr2 := arr[len(arr)-3:]
	fmt.Println(arr2) // [d e]
}

// Создайте массив из 4 логических значений и создайте срез, содержащий первые 2 элемента.
func task36() {
	arr := [4]bool{true, false, true, false}
	arr2 := arr[:2]
	fmt.Println(arr2) // [true false]
}

// Создайте два среза из 3 целых чисел и скопируйте значения первого среза во второй.
func task37() {
	arr1 := []int{1, 2, 3}
	arr2 := []int{4, 5, 6}
	copy(arr2, arr1)
	fmt.Println(arr2) // [1 2 3]
}

// Создайте два среза из 4 строк и скопируйте значения второго и третьего элемента первого среза во второй.
func task38() {
	arr1 := []string{"a", "b", "c", "d"}
	arr2 := []string{"e", "f", "g"}
	copy(arr2[1:], arr1[2:])
	fmt.Println(arr2) // [e f c d]
}

// Создайте два среза из 5 вещественных чисел и скопируйте все значения первого среза во второй.
func task39() {
	arr1 := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	arr2 := []float64{6.6, 7.7, 8.8, 9.9, 10.10}
	copy(arr2, arr1)
	fmt.Println(arr2) // [1.1 2.2 3.3 4.4 5.5]
}

// Создайте массив из 5 целых чисел, где значения второго и четвертого элемента равны 10 и 20 соответственно.
func task40() {
	arr := [5]int{1, 2, 3, 4, 5}
	arr[1] = 10
	arr[3] = 20
	fmt.Println(arr) // [1 10 3 20 5]
}

// Создайте массив из 6 строк, где первая и последняя строка равны "start" и "end" соответственно.
func task41() {
	arr := [6]string{"start", "a", "b", "c", "d", "end"}
	arr[0] = "start"
	arr[len(arr)-1] = "end"
	fmt.Println(arr) // [start a b c d end]
}

// Создайте массив из 4 логических значений, где третий элемент равен true.
func task42() {
	arr := [4]bool{true, false, true, false}
	arr[2] = true
	fmt.Println(arr) // [true false true true]
}

// Создайте массив из 4 целых чисел и скопируйте его в новый массив, затем измените второй элемент нового массива.
func task43() {
	arr := [4]int{1, 2, 3, 4}
	arr2 := arr
	arr2[1] = 10
	fmt.Println(arr)  // [1 2 3 4]
	fmt.Println(arr2) // [1 10 3 4]
}

// Создайте массив из 5 строк и скопируйте его в новый массив, затем измените первый элемент нового массива.
func task44() {
	arr := [5]string{"a", "b", "c", "d", "e"}
	arr2 := arr
	arr2[0] = "start"
	fmt.Println(arr)  // [start b c d e]
	fmt.Println(arr2) // [start b c d e]
}

// Создайте массив из 3 вещественных чисел и скопируйте его в новый массив, затем измените последний элемент нового массива.
func task45() {
	arr := [3]float64{1.1, 2.2, 3.3}
	arr2 := arr
	arr2[2] = 10.10
	fmt.Println(arr)  // [1.1 2.2 10.10]
	fmt.Println(arr2) // [1.1 2.2 10.10]
}

// Создайте массив из 6 целых чисел и создайте срез, содержащий элементы с индексами от 1 до 3. Выведите длину и вместимость среза.
func task46() {
	arr := [6]int{1, 2, 3, 4, 5, 6}
	arr2 := arr[1:4]
	fmt.Println(len(arr2), cap(arr2)) // 3 5
}

// Создайте срез из 5 строк и создайте новый срез, содержащий элементы с индексами от 2 до 4. Выведите длину и вместимость нового среза.
func task47() {
	arr := []string{"a", "b", "c", "d", "e"}
	arr2 := arr[2:4]
	fmt.Println(len(arr2), cap(arr2)) // 2 3
}

// Создайте массив из 4 логических значений и создайте срез, содержащий элементы с индексами от 0 до 2.
// Измените значение первого элемента среза и выведите исходный массив.
func task48() {
	arr := [4]bool{true, false, true, false}
	arr2 := arr[0:3]
	arr2[0] = false
	fmt.Println(arr) // [false false true false]
}

// Создайте срез из 3 целых чисел и добавьте в него значение 4, затем добавьте еще одно значение 5.
func task49() {
	arr := []int{1, 2, 3}
	arr = append(arr, 4)
	arr = append(arr, 5)
	fmt.Println(arr) // [1 2 3 4 5]
}

// Создайте пустой срез из строк и добавьте в него значения "x", "y", "z".
func task50() {
	arr := []string{}
	arr = append(arr, "x", "y", "z")
	fmt.Println(arr) // [x y z]
}

// Создайте срез из 2 вещественных чисел и добавьте в него значения 3.3, 4.4, 5.5.
func task51() {
	arr := []float64{1.1, 2.2}
	arr = append(arr, 3.3, 4.4, 5.5)
	fmt.Println(arr) // [1.1 2.2 3.3 4.4 5.5]
}
