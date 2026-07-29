package main

import (
	"fmt"
)

func main() {
	// nil slice - does not have an underlying array
	var nilSlice []int

	// non-nil, empty slice - which has an underlying array but no elements
	emptySlice := []int{}

	// non-nil, empty slice
	populatedSlice := make([]int, 0)

	fmt.Printf("nilSlice:%#v len:%d cap:%d nilSlice==nil:%t\n", nilSlice, len(nilSlice), cap(nilSlice), nilSlice == nil)
	fmt.Printf("emptySlice:%#v len:%d cap:%d emptySlice==nil:%t\n", emptySlice, len(emptySlice), cap(emptySlice), emptySlice == nil)
	fmt.Printf("populatedSlice:%#v len:%d cap:%d populatedSlice==nil:%t\n", populatedSlice, len(populatedSlice), cap(populatedSlice), populatedSlice == nil)
}
