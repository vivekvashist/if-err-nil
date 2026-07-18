// closures - local func that has access to the variables that exist in the env where
// it was declared
package main

import "fmt"

func main() {
	b := 2
	myAddOne := func(a int) {
		// referring b inside the anonymous func
		// we can read a value
		// return a + b
		// we can write a value too
		b = a + b
	}
	myAddOne(1)
	fmt.Printf("%d\n", b)
}
