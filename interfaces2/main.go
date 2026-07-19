//https://medium.com/@Sushil_Kumar/function-types-as-bridge-to-interfaces-bad0f2d5e1f4

package main

import "fmt"

type AbstractLogic interface {
	process(value int) int
}

type MapperFunc func(int) int

func (mf MapperFunc) process(value int) int {
	return mf(value) // as mf variable can be used to call the function.
}

func MapValues(values []int, l AbstractLogic) []int {
	res := make([]int, len(values))
	for i, v := range values {
		res[i] = l.process(v)
	}
	return res
}

func quadrupleTheValue(value int) int {
	return value * 4
}

func doubleTheValue(value int) int {
	return value * 2
}

func quinTupleTheValue(value int) int {
	return value * 5
}

func double(n int) int {
	return n * 2
}

func main() {
	values := []int{1, 2, 3}
	newValues := MapValues(values, MapperFunc(quadrupleTheValue))
	fmt.Println(newValues) // prints 4, 8, 12

	doubleValues := MapValues(values, MapperFunc(doubleTheValue))
	fmt.Println(doubleValues) // prints 2, 4, 6

	quinTupleValue := MapValues(values, MapperFunc(quinTupleTheValue))
	fmt.Printf("%v\n", quinTupleValue)

	var adapter MapperFunc = double
	result := adapter.process(100)
	fmt.Printf("%d\n", result)
}
