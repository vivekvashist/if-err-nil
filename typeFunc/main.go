package main

import (
	"fmt"
)

type op func(int, int) int

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func calculate(a, b int, operation op) int {
	return operation(a, b)
}

func main() {
	sum := calculate(10, 5, add)
	fmt.Printf("%d\n", sum)

	product := calculate(10, 5, multiply)
	fmt.Printf("%d\n", product)

	difference := calculate(10, 5, func(x, y int) int {
		return x - y
	})

	fmt.Printf("%d\n", difference)
}
