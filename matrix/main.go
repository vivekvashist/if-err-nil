package main

import (
	"fmt"
)

func main() {
	dnsAddr := [2][2]string{
		{"1.1.1.1", "8.8.8.8"},
		{"10.0.0.53", "10.0.0.54"},
	}

	fmt.Printf("%-18s %-18s\n", "primaryDns", "secondaryDns")
	for i := 0; i < len(dnsAddr); i++ {
		primaryDns := dnsAddr[i][0]
		secondaryDns := dnsAddr[i][1]
		fmt.Printf("%-18s %-18s\n", primaryDns, secondaryDns)
	}
}
