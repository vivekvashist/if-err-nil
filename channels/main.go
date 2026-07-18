package main

import "fmt"

func main() {
	in := make(chan string)
	out := make(chan string)

	fmt.Printf("%T\n", in) / chan string
	fmt.Printf("%v\n", in) // 0xxxxx

	go func() {
		// <-chan is the receive operator
		name := <-in // read/receive from in channel
		// chan <- // write value to a channel
		out <- fmt.Sprintf("hello, " + name)
	}()

	in <- "tom"
	close(in)
	message := <-out
	fmt.Printf("%s\n", message)
}
