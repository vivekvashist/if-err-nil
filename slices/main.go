package main

import (
	"fmt"
)

func main() {
	// cart := []string{"apple", "banana", "milk"}
	// fmt.Printf("%q\n", cart)
	// fruit := cart[:2]
	// fmt.Printf("%q\n", fruit)
	// cart[0] = "pear"
	// fmt.Printf("%q\n", cart) // ["pear" "banana" "milk"]
	// fmt.Printf("%q\n", fruit) // ["pear" "banana"]
	//	a := []int{1, 2, 3}
	//	b := a[:]
	//	a = append(a, 4)
	//	a[1] = 99
	//	fmt.Printf("%v\n", a) // [1 99 3 4]
	//	fmt.Printf("%v\n", b) // [1 2 3]

	// var intSlice []int = []int{3, 1, 4, 3, 9, 6}
	// stringSlice := []string{"The", "sky", "was", "the", "color"}
	//
	//	if n := len(intSlice); n > 0 {
	//		fmt.Printf("intSlice has %d elements\n", n)
	//	} else {
	//
	//		fmt.Printf("intSlice has no elements\n")
	//	}
	//
	//	if n := len(stringSlice); n > 0 {
	//		fmt.Printf("stringSlice has %d elements\n", n)
	//	} else {
	//
	//		fmt.Printf("stringSlice has no elements\n")
	//	}

	// nums := []int{1, 11, 21, 1211, 111221, 323311}
	// fmt.Printf("nums: %v\n", nums)
	// fmt.Printf("nums lenght: %d, nums capacity: %d\n", len(nums), cap(nums))
	//
	// middle := nums[1:3]
	// fmt.Printf("middle: %v\n", middle)
	// fmt.Printf("middle lenght: %d, middle capacity: %d\n", len(middle), cap(middle))
	//
	// fmt.Printf("nums: %v\n", nums)
	// fmt.Printf("nums lenght: %d, nums capacity: %d\n", len(nums), cap(nums))
	// nums[1] *= 4
	//
	// fmt.Printf("nums: %v\n", nums)
	// fmt.Printf("middle: %v\n", middle)
	// fmt.Printf("middle lenght: %d, middle capacity: %d\n", len(middle), cap(middle))
	//
	// middle = append(middle, 100)
	// middle = append(middle, 200)
	// middle = append(middle, 300)
	// middle = append(middle, 400)
	//
	// fmt.Printf("middle: %v\n", middle)
	// fmt.Printf("nums: %v\n", nums)
	// nums[1] *= 10
	//
	// fmt.Printf("nums: %v\n", nums)
	// fmt.Printf("middle: %v\n", middle)
	// fmt.Printf("\n\n")

	//	s := make([]string, 3, 2) // ./main.go:68:22: invalid argument: length and capacity swapped
	s := make([]string, 1, 2)

	fmt.Printf("len(s) %d cap(s) %d\n", len(s), cap(s))
	fmt.Printf("%T\n", s)
	fmt.Printf("%q\n", s) // ["" ""]
	s[0] = "hello"
	fmt.Printf("%q\n", s)
	//	s[1] = "world" // panic: runtime error: index out of range [1] with length 1
	s = append(s, "world")
	fmt.Printf("%q\n", s)
	fmt.Printf("len(s) %d cap(s) %d\n", len(s), cap(s))

	s = append(s, "golang")
	fmt.Printf("%q\n", s)
	fmt.Printf("len(s) %d cap(s) %d\n", len(s), cap(s))
	s = append(s, "new")
	fmt.Printf("%q\n", s)
	fmt.Printf("len(s) %d cap(s) %d\n", len(s), cap(s)) // len(s) 4 cap(s) 4
	s = append(s, "apple")
	fmt.Printf("len(s) %d cap(s) %d\n", len(s), cap(s)) // len(s) 5 cap(s) 8
	fmt.Printf("%#v\n", s)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%q\n", s)
}
