package main

import "fmt"

func sum(nums ...int) {
	// fmt.Print(nums ," " )
	fmt.Printf("%v ", nums)
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Printf("%d\n", total)
}

func main() {
	sum(1, 2, 3)
	sum(1, 2, 3, 4)

	nums := []int{1, 2, 3, 4, 5}

	sum(nums...)
}
