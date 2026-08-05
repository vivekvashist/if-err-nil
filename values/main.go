package main

import "fmt"

func main() {
	fmt.Printf("%s\n", "go"+"lang")
	fmt.Printf("%f\n", 4.0/3.0)
	fmt.Printf("%d\n", 42+1)
	fmt.Printf("\n")
	fmt.Printf("%t\n", true && false)
	fmt.Printf("%t\n", false && true)
	fmt.Printf("\n")
	fmt.Printf("%t\n", true || false)
	fmt.Printf("%t\n", false || true)
	fmt.Printf("\n")
	fmt.Printf("%t\n", !true)
	fmt.Printf("%t\n", !false)
	fmt.Printf("\n")
}
