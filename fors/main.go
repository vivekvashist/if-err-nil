package main

import "fmt"

func main() {
	i := 1
	for i <= 5 {
		fmt.Printf("%d\n", i)
		i++
	}
	fmt.Printf("\n")

	for j := 1; j <= 5; j++ {
		fmt.Printf("%d\n", j)
	}

	fmt.Printf("\n")

	for k := range 5 {
		fmt.Printf("%d\n", k)
	}

	fmt.Printf("\n")

	for {
		fmt.Printf("loop\n")
		break
	}

	fmt.Printf("\n")
	for n := range 10 {
		if n%2 == 0 {
			continue
		}
		fmt.Printf("%d\n", n)
	}
}
