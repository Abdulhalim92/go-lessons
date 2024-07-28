package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// Упаковка коробок
func sixth() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscanf(reader, "%d\n", &t)

	for i := 0; i < t; i++ {
		var n, k int
		fmt.Fscanf(reader, "%d %d\n", &n, &k)

		var m int
		fmt.Fscanf(reader, "%d\n", &m)

		boxes := make([]int, m)
		for j := 0; j < m; j++ {
			var a int
			fmt.Fscanf(reader, "%d", &a)
			boxes[j] = 1 << a
		}

		sort.Sort(sort.Reverse(sort.IntSlice(boxes)))

		trips := 0
		for len(boxes) > 0 {
			trips++
			remaining := k
			for i := 0; i < n && len(boxes) > 0; {
				if boxes[0] <= remaining {
					remaining -= boxes[0]
					boxes = boxes[1:]
				} else {
					i++
					remaining = k
				}
			}
		}

		fmt.Fprintln(writer, trips)
	}
}
