package main

import (
	"fmt"
)

func main() {
	a := nil
	fmt.Printf("%v\n", a)

	// $ go vet main.go
	// # command-line-arguments
	//# [command-line-arguments]
	//vet: ./main.go:8:7: use of untyped nil in assignment

	//$ go run main.go
	///# command-line-arguments
	//./main.go:8:7: use of untyped nil in assignment

}
