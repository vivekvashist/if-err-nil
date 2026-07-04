package main

import (
	"io"
	"log"
	"os"
)

func main() {
	input, err := os.Open("lab.txt")
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(os.Stdout, input)
}
