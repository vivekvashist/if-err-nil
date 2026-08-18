package main

import "fmt"

func main() {
	var num int = 42

	types := []interface{}{
		true, // bool

		"hello", // string

		int(10),   // int
		int8(8),   // int8
		int16(16), // int16
		int32(32), // int32
		int64(64), // int64

		uint(10),   // uint
		uint8(8),   // uint8
		uint16(16), // uint16
		uint32(32), // uint32
		uint64(64), // unint64

		float32(3.14),    // float32
		float64(3.14159), // float64

		complex64(1 + 2i),  // complex64
		complex128(2 + 4i), // complex128

		byte('A'), // byte (alias for uint8)
		rune('⌘'), // rune (alias for int32)

		[3]int{1, 2, 3},                       // array
		[]int{1, 2, 3},                        // slice
		map[string]int{"a": 1},                // map
		make(chan int),                        // channel
		func() {},                             // function
		struct{ Name string }{Name: "Gopher"}, // struct

		&num, // pointer (*int)

	}

	for _, v := range types {
		fmt.Printf("%-20T | %v\n", v, v)
	}
}
