package main

import (
	"fmt"
)

func addOne(a int) int {
	return a + 1
}

func main() {
	// assigning func to a local variable
	// note no () we are not invoking addOne we are referencing it

	myAddOne := addOne
	fmt.Printf("%d\n", addOne(1))
	fmt.Printf("%d\n", myAddOne(1))
}
