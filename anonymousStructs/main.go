package main

import (
	"fmt"
)

func main() {
	st := struct {
		Host string
		Port int
	}{
		Host: "localhost",
		Port: 8080,
	}
	fmt.Printf("%s %d\n", st.Host, st.Port)

	servers := []struct {
		Host string
		Port int
	}{
		{"localhost", 9090},
		{"192.168.1.1", 443},
		{"example.com", 80},
	}

	for _, s := range servers {
		fmt.Printf("%s %d\n", s.Host, s.Port)
	}
}
