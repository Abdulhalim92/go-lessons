package main

import (
	"bufio"
	"fmt"
	"os"
)

// Ошибка округления
func second() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var n, p int
		fmt.Fscan(in, &n, &p)
		totalLoss := 0.0

		for j := 0; j < n; j++ {
			var ai int
			fmt.Fscan(in, &ai)

			// Реальная комиссия
			actualCommission := float64(ai) * float64(p) / 100

			// Округленная до рублей комиссия (ошибка)
			roundedCommission := float64(ai * p / 100)

			// Упущенная комиссия
			loss := actualCommission - roundedCommission
			totalLoss += loss
		}

		// Выводим результат с двумя знаками после запятой
		fmt.Fprintf(out, "%.2f\n", totalLoss)
	}
}
