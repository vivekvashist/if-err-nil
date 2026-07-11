package main

import (
	"fmt"
)

func main() {
	for x := 0; x < 10; x++ {
		fmt.Printf("Hello\n")
	}

	x := 0
	for x % 2 == 0 {
		fmt.Printf("%d\n", x)
		fmt.Printf("world\n")
		x++
	}
}
