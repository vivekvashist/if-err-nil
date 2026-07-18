// go is pass-by-value including pointers
// when &a is passed to a func it receives a copy of the pointer value i.e address
// new(int) point to a different memory addr
package main

import "fmt"

func main() {
	a := 10
	b := &a
	c := a

	fmt.Printf("a: %d, b: %p, *b: %d, c: %d\n", a, b, *b, c)

	a = 20
	fmt.Printf("a: %d, b: %p, *b: %d, c: %d\n", a, b, *b, c)

	*b = 30
	fmt.Printf("a: %d, b: %p, *b: %d, c: %d\n", a, b, *b, c)

	c = 40
	fmt.Printf("a: %d, b: %p, *b: %d, c: %d\n", a, b, *b, c)

	var bb *int
	fmt.Printf("%T\n", bb) // *int

	// zero-value for a pointer is nil
	fmt.Printf("%v\n", bb) // <nil>

	// reading a nil value cause panic
	// panic: runtime error: invalid memory address or nil pointer dereference
	// you cannot read of write to a nil pointer
	// fmt.Printf("%v\n", *bb)

	// new can be used to create a pointer
	// it allocate memory as well and does not cause panic

	xx := new(int)
	fmt.Printf("%v\n", xx)  // 0xxxxx addr
	fmt.Printf("%T\n", xx)  // *int
	fmt.Printf("%v\n", *xx) // 0 zero-value of an int type

	aa := 20
	fmt.Printf("%d\n", aa)
	setTo10(&aa)
	fmt.Printf("%d\n", aa)

}
func setTo10(aa *int) {
	aa = new(int) // call by value this a a new pointer
	*aa = 10
}
