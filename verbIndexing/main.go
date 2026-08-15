package main

import "fmt"

func main() {
	fmt.Printf("normal order: %d, %s\n", 42, "hello")
	fmt.Printf("reversed: %[2]s, %[1]d\n", 42, "hello")
	fmt.Printf("hello %[1]s, hello again %[1]s\n", "gopher")
	fmt.Printf("hello %[1]s, hello again %[1]s and again %[1]s\n", "gopher")
}
