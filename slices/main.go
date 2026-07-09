package main

import (
	"fmt"
)

func main() {
	// cart := []string{"apple", "banana", "milk"}
	// fmt.Printf("%q\n", cart)
	// fruit := cart[:2]
	// fmt.Printf("%q\n", fruit)
	// cart[0] = "pear"
	// fmt.Printf("%q\n", cart) // ["pear" "banana" "milk"]
	// fmt.Printf("%q\n", fruit) // ["pear" "banana"]
	a := []int{1, 2, 3}
	b := a[:]
	a = append(a, 4)
	a[1] = 99
	fmt.Printf("%v\n", a) // [1 99 3 4]
	fmt.Printf("%v\n", b) // [1 2 3]
}
