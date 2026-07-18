package main

import "fmt"

// very useful
// http.HandleFunc("/hello", printHello)

func addOne(a int) int {
	return a + 1
}

func addTwo(a int) int {
	return a + 2
}

func printOperation(a int, f func(int) int) {
	fmt.Printf("%d\n", f(a))
}

func main() {
	printOperation(1, addOne)
	printOperation(1, addTwo)
}
