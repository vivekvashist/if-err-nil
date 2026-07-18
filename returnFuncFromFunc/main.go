package main

import "fmt"

func makeAdder(b int) func(int) int {
	// creating and returning anonymous func
	return func(a int) int {
		return a + b
	}
}

// func that takes in a func and returns a func
// whats the use of doing all this
// http.HandleFunc("/hello", printHello)
// we could build factory wrapper funcs that take in business logic
// and return the business logic with the house keeping code
// http.HandleFunc("/bye", log(printBye))

func makeDoubler(f func(int) int) func(int) int {
	return func(a int) int {
		b := f(a)
		return b * 2
	}
}

func main() {
	addOne := makeAdder(1)
	//addTwo := makeAdder(2)

	doubleAaddOne := makeDoubler(addOne)

	fmt.Printf("%d\n", addOne(1))
	//fmt.Printf("%d\n", addTwo(1))
	fmt.Printf("%d\n", doubleAaddOne(1))
}
