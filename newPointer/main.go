package main

import "fmt"

func setTo10Fail(a *int) {
	fmt.Println("  inside, before reassign, a points to:", a)
	a = new(int)
	fmt.Println("  inside, after reassign,  a points to:", a)
	*a = 10
}

func main() {
	a := 20
	fmt.Println("main's a is at:", &a, "value:", a)
	setTo10Fail(&a)
	fmt.Println("main's a is at:", &a, "value:", a)
}
