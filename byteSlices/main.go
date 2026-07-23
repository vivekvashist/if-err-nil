package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	defer f.Close()

	// fmt.Printf("%t\n", f)  // &{%!t(*os.file=&{{{0 0 0} 4 {<nil>} {0} 0 1 true true true} test.txt {[] {} <nil>} false false false})}
	//fmt.Printf("%v\n", f)  // &{0x231830dba180}
	//fmt.Printf("%T\n", f)  // *os.File
	//fmt.Printf("%s\n", f)  // same as %t
	//fmt.Printf("%#v\n", f) // &os.File{file:(*os.file)(0x27283a7fe180)}
	//fmt.Printf("%+v\n", f) // │&{file:0x3f53ba9d6180}

	b := make([]byte, 50)
	n, err := f.Read(b)
	// fmt.Printf("%d: %x\n", n, b)
	// fmt.Printf("c: %c\n", f)

	stringVersion := string(b)
	fmt.Printf("stringify: %d: %s\n", n, stringVersion)
	someString := "foo bar baz\n"
	fmt.Printf("writing string to test2.txt\n")
	os.WriteFile("test2.txt", []byte(someString), 0644)

}
