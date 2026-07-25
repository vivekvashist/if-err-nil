package main

import (
	"fmt"
	"unsafe"
)

// var j int = 20
var j = 20

// j := 20
// ./main.go:9:1: syntax error: non-declaration statement outside function body

var z int // zero value of int

func main() {
	// var i int = 10
	// var i = 10
	ii := 10
	var y int // zero value of int
	// byte is uint8
	// rune is int32

	fmt.Printf("%T\n", ii)
	fmt.Printf("%v\n", ii)
	fmt.Printf("%d\n", ii)
	fmt.Printf("%d\n", j)
	fmt.Printf("%d\n", z)
	fmt.Printf("%d\n", y)

	fmt.Printf("\n\n\n")

	var i int
	fmt.Printf("%dbits\n", unsafe.Sizeof(i)*8) // 64bits
	var ui uint
	fmt.Printf("%dbits\n", unsafe.Sizeof(ui)*8) // 64bits

	var i8 int8
	fmt.Printf("%dbits\n", unsafe.Sizeof(i8)*8)
	var i16 int16
	fmt.Printf("%dbits\n", unsafe.Sizeof(i16)*8)
	var i32 int32
	fmt.Printf("%dbits\n", unsafe.Sizeof(i32)*8)
	var i64 int64
	fmt.Printf("%dbits\n", unsafe.Sizeof(i64)*8)
}
