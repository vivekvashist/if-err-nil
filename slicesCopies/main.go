package main

import (
	"fmt"
)

func printer(w []string) {
	for _, word := range w {
		fmt.Printf("%s\n", word)
	}
	fmt.Printf("\n")
	// [4]string array does not change its copy
	// []string slice is a reference type it changes
	w[2] = "blue"
}

func main() {
	words := []string{"the", "quick", "brown", "fox"}
	printer(words)
	printer(words)

}
