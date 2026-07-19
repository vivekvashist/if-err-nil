package main

import "fmt"

type tester interface {
	test(int) bool
}

func runTests(i int, tests []tester) bool {
	result := true
	for _, test := range tests {
		fmt.Printf("inside runTests range\n")
		result = result && test.test(i)
	}
	return result
}

type testerFunc func(int) bool

func (tf testerFunc) test(i int) bool {
	return tf(i)
}

func main() {
	result := runTests(10, []tester{
		testerFunc(func(i int) bool {
			fmt.Printf("inside first func\n")
			return i%2 == 0
		}),
		testerFunc(func(i int) bool {
			fmt.Printf("inside second func\n")
			return i < 20
		}),
		testerFunc(func(i int) bool {
			fmt.Printf("inside third func\n")
			return i > 5
		}),
	})
	fmt.Printf("%t\n", result)
	fmt.Printf("\n")
}
