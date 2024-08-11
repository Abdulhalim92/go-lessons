package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer func() { _ = out.Flush() }()

	var num int
	fmt.Fscan(in, &num)

	fmt.Fprintln(out, num)
}
