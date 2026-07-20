package main

import (
	"fmt"
)

func main() {
	connections := make([]string, 0)
	fmt.Printf("active connections: %d\n", len(connections))

	connections = append(connections, "192.168.1.20")
	fmt.Printf("active connections: %d\n", len(connections))

	connections = append(connections, "192.168.1.22")
	fmt.Printf("active connections: %d\n", len(connections))

	fmt.Printf("first connected host: %s\n", connections[0])

	connections[0] = "192.168.1.99"
	fmt.Printf("updated connected host: %s\n", connections[0])

	dnsServers := make([]string, 2)

	fmt.Printf("dnsServers %q\n", dnsServers)
	fmt.Printf("total dnsServers: %d\n", len(dnsServers))
	fmt.Printf("slot 0 DNS (empty): %q\n", dnsServers[0])

	dnsServers = append(dnsServers, "1.1.1.1")
	fmt.Printf("dnsServers %q\n", dnsServers)             /// ["" "" "1.1.1.1"]
	fmt.Printf("total dnsServers: %d\n", len(dnsServers)) // 3

	dnsServers[0] = "8.8.8.8"
	fmt.Printf("dnsServers %q\n", dnsServers) /// ["8.8.8.8" "" "1.1.1.1"]

	dnsServers[1] = "9.9.9.9"
	fmt.Printf("dnsServers %q\n", dnsServers) /// ["8.8.8.8" "" "1.1.1.1"]

	for _, d := range dnsServers {
		fmt.Printf("%s\n", d)
	}

	v6dnsServers := []string{"2620:fe::fe", "2620:fe::9", "2001:4860:4860::8888", "2001:4860:4860::8844"}

	for i, v := range v6dnsServers {
		fmt.Printf("%d --> %v\n", i, v)
	}
	// slices are reference types
	s := v6dnsServers[1:3]
	fmt.Printf("%q\n", s)
	fmt.Printf("%q\n", v6dnsServers)

	s[1] = "1.1.1.1"
	fmt.Printf("%q\n", s)            // ["2620:fe::9" "1.1.1.1"]
	fmt.Printf("%q\n", v6dnsServers) // ["2620:fe::fe" "2620:fe::9" "1.1.1.1" "2001:4860:4860::8844"]

}
