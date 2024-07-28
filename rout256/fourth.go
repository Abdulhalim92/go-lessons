package main

import (
	"bufio"
	"fmt"
	"os"
)

// Сломанный сервер
// Для решения данной задачи можно использовать метод двух указателей (или скользящего окна).
// Основная идея заключается в том, чтобы поддерживать окно, содержащее не более двух различных
// идентификаторов клиентов, и по мере необходимости расширять или сужать его, чтобы найти
// максимальную длину такого окна.
func fourth() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for ; t > 0; t-- {
		var n int
		fmt.Fscan(in, &n)

		ids := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &ids[i])
		}

		maxLen := longestSubarrayWithTwoDistinct(ids)
		fmt.Fprintln(out, maxLen)
	}
}

func longestSubarrayWithTwoDistinct(ids []int) int {
	counts := make(map[int]int)
	left, right, maxLen := 0, 0, 0

	for right < len(ids) {
		counts[ids[right]]++
		for len(counts) > 2 {
			counts[ids[left]]--
			if counts[ids[left]] == 0 {
				delete(counts, ids[left])
			}
			left++
		}
		if currentLen := right - left + 1; currentLen > maxLen {
			maxLen = currentLen
		}
		right++
	}

	return maxLen
}
