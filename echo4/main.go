package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	n := flag.Bool("n", false, "omit trailing newline")
	sep := flag.String("s", " ", "separator")

	fmt.Printf("%T\n", n)   // *bool
	fmt.Printf("%T\n", sep) // *string
	fmt.Printf("%p\n", n)   // addr
	fmt.Printf("%p\n", sep) // addr

	flag.Parse()
	fmt.Printf("%s\n", strings.Join(flag.Args(), *sep))
	fmt.Printf("%v\n", *sep)

	if !*n {
		fmt.Printf("\n")
	}
}
