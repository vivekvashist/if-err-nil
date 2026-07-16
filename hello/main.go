package main

import (
	"fmt"
)

func main() {
	//fmt.Printf("Hello World!\n")

	s := "hello world🌍"

	b := s[0]
	bb := s[4]
	fmt.Printf("%d %d\n", b, bb)                 // 104 111
	fmt.Printf("%s %s\n", string(b), string(bb)) // h o

	s2 := s[0:5]

	fmt.Printf("%s\n", s2) // hello

	aa := "hello, "
	var r rune = 127757
	fmt.Printf("%T\n", r) // int32
	fmt.Printf("%v\n", r) // 127757

	aa = aa + string(r)

	fmt.Printf("%s\n", aa)

	// rr := "🌍" // string	double quotes
	rr := '🌍'              // rune single quotes
	fmt.Printf("%v\n", rr) // 127757
	fmt.Printf("%T\n", rr) // int32

	// 0 104 h int32
	// 1 101 e int32
	// 2 108 l int32
	// 3 108 l int32
	// 4 111 o int32
	// 5 32   int32
	// 6 119 w int32
	// 7 111 o int32
	// 8 114 r int32
	// 9 108 l int32
	// 10 100 d int32
	// 11 127757 🌍 int32

	// v is a rune i.e int32
	for i, v := range s {
		fmt.Printf("%d %v %v %T \n", i, v, string(v), v)
	}

}
