package main

import (
	"fmt"
)


//count:	 Value Of [ 10 ]	 Addr Of [ 0x21c44a7800a8 ]
//inc:	 Value Of inc [ 0x21c44a7800a8]	 Addr Of &inc [ 0x21c44a770058 ] Value *inc Points to [ 11 ]
//&count:	 Value Of [ 11 ]	 Addr Of [ 0x21c44a7800a8 ]

func increment(inc *int) {
	*inc++
	fmt.Printf("inc:\t Value Of inc [ %p]\t Addr Of &inc [ %p ] Value *inc Points to [ %d ]\n", inc, &inc, *inc)
}

func main() {
	count := 10
	fmt.Printf("count:\t Value Of [ %d ]\t Addr Of [ %p ]\n", count, &count)

	increment(&count)
	fmt.Printf("&count:\t Value Of [ %d ]\t Addr Of [ %p ]\n", count, &count)

}
