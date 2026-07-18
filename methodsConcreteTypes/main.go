package main

import "fmt"

type Port int

// value receiver
func (p Port) IsPrivileged() bool {
	return p > 0 && p < 1024
}

// pointer receiver
func (p *Port) Next() {
	*p = *p + 1

}

func main() {
	p := Port(1023)
	fmt.Printf("current port: %d\n", p)
	fmt.Printf("%d: is priviliged %t\n", p, p.IsPrivileged()) // false
	p.Next()
	fmt.Printf("new port: %d\n", p)                           // 1024
	fmt.Printf("%d: is priviliged %t\n", p, p.IsPrivileged()) // true

}
