package main

import (
	"fmt"
	"os"
)

func main() {
	//	word := os.Args[1]
	//
	//	if word == "hello" {
	//		fmt.Printf("hi yourself\n")
	//	} else if word == "goodby" {
	//		fmt.Printf("so long\n")
	//	} else if word == "greetings" {
	//		fmt.Printf("salutations\n")
	//	} else {
	//		fmt.Printf("what did you say\n")
	//	}

	// use switch instead of if-else-if-else

	word := os.Args[1]
	switch l := len(word); word {
	case "hello":
		fmt.Printf("hi yourself\n")
	case "goodby", "bye":
		fmt.Printf("goodby\n")
	case "greetings":
		fmt.Printf("salutations\n")
	default:
		fmt.Printf("what did you say %s of lenght %d\n", word, l)
	}
}

// atoz := "the quick brown fox jumps over the lazy dog"
//
// vowels := 0
// consonants := 0
//
//	for _, r := range atoz {
//		switch r {
//		case 'a', 'e', 'i', 'o', 'u':
//			vowels += 1
//		default:
//			consonants += 1
//		}
//		fmt.Printf("Vowels: %d; Consonants: %d\n", vowels, consonants)
//
// }
