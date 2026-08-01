package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "10.1.1.1"
	h := sha256.New()

	h.Write([]byte(s))

	bs := h.Sum(nil)

	fmt.Printf("%s\n", s)
	fmt.Printf("%x\n", bs)

}
