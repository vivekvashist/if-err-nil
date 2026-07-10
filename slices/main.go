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
	//	a := []int{1, 2, 3}
	//	b := a[:]
	//	a = append(a, 4)
	//	a[1] = 99
	//	fmt.Printf("%v\n", a) // [1 99 3 4]
	//	fmt.Printf("%v\n", b) // [1 2 3]

	var intSlice []int = []int{3, 1, 4, 3, 9, 6}
	stringSlice := []string{"The", "sky", "was", "the", "color"}

	if n := len(intSlice); n > 0 {
		fmt.Printf("intSlice has %d elements\n", n)
	} else {
		fmt.Printf("intSlice has no elements\n")
	}

	if n := len(stringSlice); n > 0 {
		fmt.Printf("stringSlice has %d elements\n", n)
	} else {
		fmt.Printf("stringSlice has no elements\n")
	}
}
