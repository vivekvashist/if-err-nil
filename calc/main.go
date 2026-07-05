package main

import (
	"fmt"
	"strconv"
)

func add(i int, j int) int { return i + j }

func sub(i int, j int) int { return i - j }

func mul(i int, j int) int { return i * j }

func div(i int, j int) int { return i / j }

var opMap = map[string]func(int, int) int{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

func main() {
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"two", "+", "three"},
		{"5"},
	}
	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Printf("invalid expression: %v\n", expression)
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		fmt.Printf("%v\n", p1)
		if err != nil {
			fmt.Printf("%v\n", err)
			continue
		}

		op := expression[1]
		fmt.Printf("%v\n", op)
		opFunc, ok := opMap[op]
		fmt.Printf("%T\n", opFunc)

		if !ok {
			fmt.Printf("unsupported operator:%v\n", op)
			continue
		}
		p2, err := strconv.Atoi(expression[2])
		fmt.Printf("%v\n", p2)
		if err != nil {
			fmt.Printf("%v\n", err)
			continue
		}
		result := opFunc(p1, p2)
		fmt.Printf("%d\n", result)

	}
}
