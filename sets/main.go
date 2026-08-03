package main

import "fmt"

func main() {
	openPorts := map[string]bool{
		"80":  true,
		"443": true,
		"22":  true,
	}

	port := "8080"

	if openPorts[port] {
		fmt.Printf("port %s is OPEN\n", port)
	} else {
		fmt.Printf("port %s is CLOSED\n", port)
	}

}
