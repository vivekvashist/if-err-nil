package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// $ go doc fmt.Fprintf | grep func
	// func Fprintf(w io.Writer, format string, a ...any) (n int, err error)

	fmt.Fprintf(os.Stdout, "%s, ", "hello")

	// $ go doc bufio.NewWriter | grep func
	// func NewWriter(w io.Writer) *Writer

	buf := bufio.NewWriter(os.Stdout)
	fmt.Fprintf(buf, "%s\n", "world")
	buf.Flush()
}
