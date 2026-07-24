package main

import (
	"fmt"
	"reflect"
)


type Person struct {
	name string
	age int
}

func main() {
	i := 1
	f := 1.1
	s := "hello"
	b := true

	p := Person{name: "tom", age: 10}
	pPtr := &p

	fmt.Printf("%v\n", reflect.TypeOf(i))
	fmt.Printf("%v\n", reflect.TypeOf(f))
	fmt.Printf("%v\n", reflect.TypeOf(s))
	fmt.Printf("%v\n", reflect.TypeOf(b))
	fmt.Printf("%v\n", reflect.TypeOf(p))
	fmt.Printf("%v\n", reflect.TypeOf(pPtr))
}
