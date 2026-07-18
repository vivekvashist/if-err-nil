package main

import (
	"fmt"
)

func main() {
	a := 10

	// ./main.go:10:5: non-boolean condition in if statement
	//  if 0 {
	//	b := a / 2

	// a, b are work in both if and else block
	if b := a / 2; a > 5 {
		fmt.Printf("a is bigger than 5 and a is %d b is %d \n", a, b)
	} else {
		// b will not work in else block
		fmt.Printf("a is less than or equal to and a is %d b is %d \n", a, b)
	}
}
