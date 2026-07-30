package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("a: %v\n", a)
	fmt.Printf("len(a):%d cap(a):%d\n", len(a), cap(a))

	b := a[:3]
	fmt.Printf("b:=a[:3] : %v\n", b)

	b[0] = 10
	b = append(b, 15)
	fmt.Printf("a: %v\n", a) // [10 2 3 15 5 6]
	fmt.Printf("b: %v\n", b) // [10 2 3 15]

	b = a[2:] // Fixed: use '=' instead of ':='
	fmt.Printf("b: %v\n", b)
	fmt.Printf("len(b):%d cap(b):%d\n", len(b), cap(b))

	b[0] = 100
	b = append(b, 15)

	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("len(b):%d cap(b):%d\n", len(b), cap(b))

	fmt.Printf("\n")
}
