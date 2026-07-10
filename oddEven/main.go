package main

import (
	"fmt"
)

func main() {
	var (
		even int
		odd  int
	)

	numbers := []int{1, 2, 3, 5, 7, 9, 10}

	for _, n := range numbers {
		if n%2 == 0 {
			even += 1
		} else {
			odd += 1
		}
	}

	fmt.Printf("Even %d, odd %d \n", even, odd)
}
