package main

import (
	"fmt"
)

func divAndRemainder(a, b int) (int, int) {
	return a / b, a % b
}

func main() {
	div, rem := divAndRemainder(2, 3)
	fmt.Printf("%d, %d\n", div, rem)

	// ./main.go:15:9: no new variables on left side of :=
	//	div, _ := divAndRemainder(2, 3)
	div, _ = divAndRemainder(2, 3)
	fmt.Printf("%d\n", div)
	_, rem = divAndRemainder(2, 3)
	fmt.Printf("%d\n", rem)

	// does nothing
	divAndRemainder(100, -100)
}
