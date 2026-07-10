package main

import (
	"fmt"
)

const (
	msg     = "%d %d %d\n"
	answer1 = iota
	answer2
	answer3
)

func main() {
	fmt.Printf(msg, answer1, answer2, answer3)
}
