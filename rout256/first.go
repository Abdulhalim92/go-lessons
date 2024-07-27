package main

import (
	"bufio"
	"fmt"
	"os"
)

func first() {
	var out *bufio.Writer
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fprint(out, "OK")
}
