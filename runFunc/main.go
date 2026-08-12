package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Printf("%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	f, err := os.Create("hello.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	n, err := f.WriteString("hello world")
	if err != nil {
		return err
	}

	fmt.Printf("wrote %d bytes.\n", n)
	return nil
}
