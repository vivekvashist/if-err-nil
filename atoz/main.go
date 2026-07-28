package main

import (
	"fmt"
)

func main() {
	for ch := 'a'; ch <= 'z'; ch++ {
		fmt.Printf("%c", ch)
	}
	fmt.Printf("\n")

	for ch := 'A'; ch <= 'Z'; ch++ {
		fmt.Printf("%c", ch)
	}
	fmt.Printf("\n")
}
