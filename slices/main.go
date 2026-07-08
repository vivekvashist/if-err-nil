package main

import (
	"fmt"
)

func main() {
	cart := []string{"apple", "banana", "milk"}
	fmt.Printf("%q\n", cart)
	fruit := cart[:2]
	fmt.Printf("%q\n", fruit)
	cart[0] = "pear"
	fmt.Printf("%q\n", cart) // ["pear" "banana" "milk"]
	fmt.Printf("%q\n", fruit) // ["pear" "banana"]

}
