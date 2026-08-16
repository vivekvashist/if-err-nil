package main

import "fmt"

func main() {
	var num int = 42

	types := []interface{}{
		true,      // bool
		"hello",   // string
		int(10),   // int
		int8(8),   // int8
		int16(16), // int16
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
