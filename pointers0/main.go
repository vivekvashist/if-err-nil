package main

import (
	"fmt"
)

func main() {
	i, j := 42, 2701

	fmt.Printf("i: %d\n", i)
	fmt.Printf("j: %d\n", j)
	fmt.Printf("&j: %p\n", &j)

	fmt.Printf("\n")
	fmt.Printf("&i: %p\n", &i)
	p := &i
	fmt.Printf("p (value of p is the address of i &i): %p\n", p)                       // value of p is the address i i.e &i
	fmt.Printf("*p (is the value at that address which is the value at i) : %d\n", *p) // *p is the value at that address which is the value at i
	fmt.Printf("\n")

	fmt.Printf("TYPE *p: %T\n", *p)
	fmt.Printf("VALUE *p: %v\n", *p)

	fmt.Printf("BEFORE i: %d\n", i)
	*p = 21
	fmt.Printf("AFTER i: %d\n", i)

	p = &j // we can do this because p is *int
	fmt.Printf("BEFORE j: %d\n", j)
	*p = *p / 37
	fmt.Printf("AFTER j: %d\n", j)

}
