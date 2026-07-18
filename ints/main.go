package main

import (
	"fmt"
)

// var j int = 20
var j = 20

// j := 20
// ./main.go:9:1: syntax error: non-declaration statement outside function body

var z int // zero value of int

func main() {
	// var i int = 10
	// var i = 10
	i := 10
	var y int // zero value of int
	// byte is uint8
	// rune is int32

	fmt.Printf("%T\n", i)
	fmt.Printf("%v\n", i)
	fmt.Printf("%d\n", i)
	fmt.Printf("%d\n", j)
	fmt.Printf("%d\n", z)
	fmt.Printf("%d\n", y)
}
