package main

import (
	"fmt"
)

func main() {
	//var i int
	//for i = 0; i <= 9; i++ {
	//	fmt.Println(i)
	//}
	//
	//fmt.Println(i)

	//var num = 546
	//for num > 0 {
	//	a := num % 10
	//	fmt.Println(a)
	//	num = num / 10
	//}
	//
	//for n := 546; n > 0; n /= 10 {
	//	a := n % 10
	//	fmt.Println(a)
	//}
	//
	//for {
	//	//time.Sleep(1 * time.Second)
	//	fmt.Println("Hello")
	//}

	//for i := 1; i <= 100; i++ {
	//	if i%3 == 0 {
	//		break
	//	}
	//	fmt.Println(i)
	//}
	//
	//for i := 1; i < 10; i++ {
	//	for j := 1; j < 10; j++ {
	//		if j == 5 {
	//			continue
	//		}
	//		fmt.Printf("%d * %d = %d\n", i, j, i*j)
	//	}
	//	fmt.Println("-------------------------")
	//}
	series2()
}

func series2() {
	var p float64 = 1
	var n float64
	for i := 0; i < 10; i++ {
		fmt.Scan(&n)
		p = p * n
	} // 1*1*2*1*2*1*2*1*2*1*2

	fmt.Println(p)
}
