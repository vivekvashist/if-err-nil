package main

import (
	"fmt"
)

func main() {

	m := map[string]int{"apples": 1}
	fmt.Printf("%v\n", m)
	fmt.Printf("%#v\n", m)
	fmt.Printf("%T\n", m)

	if v, ok := m["apples"]; ok {
		fmt.Printf("found key %q with value %d\n", "apples", v)
	} else {
		fmt.Printf("key %q not found", "apples")
	}
}
