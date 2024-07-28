package main

import (
	"bufio"
	"fmt"
	"os"
)

// Корень дерева
// Прочитаем последовательность чисел.
// Соберем информацию о всех вершинах и их сыновьях.
// Найдём вершину, которая не присутствует в качестве сына ни у одной другой вершины.
func third() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var m int
		fmt.Fscan(in, &m)

		code := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &code[j])
		}

		// Множество всех вершин и их сыновей
		vertices := make(map[int]struct{})
		children := make(map[int]struct{})

		// Обрабатываем код дерева
		j := 0
		for j < m {
			v := code[j]
			degree := code[j+1]
			vertices[v] = struct{}{}

			for k := 0; k < degree; k++ {
				child := code[j+2+k]
				children[child] = struct{}{}
			}

			j += 2 + degree
		}

		// Находим корень
		var root int
		for v := range vertices {
			if _, found := children[v]; !found {
				root = v
				break
			}
		}

		fmt.Fprintln(out, root)
	}
}
