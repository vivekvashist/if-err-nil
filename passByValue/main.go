package main

import (
	"fmt"
)

func increment(inc int) {
	inc++
	fmt.Printf("inc:\t Value Of [ %d ]\t Addr Of [ %p ]\n", inc, &inc)
}

//count:	 Value Of [ 10 ]	 Addr Of [ 0x2a1c6771a0a8 ]
//inc:	 	 Value Of [ 11 ]	 Addr Of [ 0x2a1c6771a0c0 ]
//count:	 Value Of [ 10 ]	 Addr Of [ 0x2a1c6771a0a8 ]

func main() {
	count := 10
	fmt.Printf("count:\t Value Of [ %d ]\t Addr Of [ %p ]\n", count, &count)

	increment(count)
	fmt.Printf("count:\t Value Of [ %d ]\t Addr Of [ %p ]\n", count, &count)

}
