package main

import (
	"fmt"
)

func main() {
	var fruits [5]string

	fmt.Printf("%v\n", fruits) // [   ]
	fmt.Printf("%T\n", fruits)
	fmt.Printf("%q\n", fruits) // ["" "" "" "" ""]

	for i, fruit := range fruits {
		fmt.Printf("%d -> %q\n", i, fruit)
		//0 -> ""
		//1 -> ""
		//2 -> ""
		//3 -> ""
		//4 -> ""
	}

	fruits[0] = "Apple"
	fruits[1] = "Plum"
	fruits[2] = "Orange"
	fruits[3] = "Grape"
	fruits[4] = "Mango"

	//	for i, fruit := range fruits {
	//		fmt.Printf("%d -> %q\n", i, fruit)

	for i := range fruits {
		fmt.Printf("%d\n", i)
	}

}
