package main

import (
	"fmt"
)

func main() {

	type user struct {
		name  string
		age   int
		email string
	}

	u := user{"tom", 20, "tom@tom.com"}

	fmt.Printf("%T\n", u)
	fmt.Printf("%v\n", u)
	fmt.Printf("%#v\n", u)
	fmt.Printf("%+v\n", u)
	fmt.Printf("%p\n", &u)

}
