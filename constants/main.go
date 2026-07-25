package main

import (
	"fmt"
	"time"
)

// const sleepSeconds = 2 // Untyped constant

func main() {
	sleepSeconds := 2

	fmt.Printf("sleeping for %d seconds...\n", sleepSeconds)

	//invalid operation: sleepSeconds * time.Second (mismatched types int and time.Duration)

	// Convert sleepSeconds (int) to time.Duration
	time.Sleep(time.Duration(sleepSeconds) * time.Second)

	fmt.Println("done")
}
