package main

import (
	"fmt"
)

func main() {
	c := make(chan string, 1) // buffered channels
	c <- "hello"
	s := <-c
	fmt.Printf("%s\n", s)
}
