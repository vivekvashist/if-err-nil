package main

import (
	"fmt"
)

func main() {

	type NetworkInterface struct {
		MTU     int
		macAddr string
		IPAddr  string
	}

	eth0 := NetworkInterface{}
	fmt.Printf("%v\n", eth0)  // {0  }
	fmt.Printf("%#v\n", eth0) // main.NetworkInterface{MTU:0, macAddr:"", IPAddr:""}
	fmt.Printf("%+v\n", eth0) // {MTU:0 macAddr: IPAddr:}

	eth0.IPAddr = "192.168.1.10"
	eth0.MTU = 1500
	eth0.macAddr = "00:1a:2b:3c:4d:5e"

	fmt.Printf("configured eth0: %+v\n", eth0) // configured eth0: {MTU:1500 macAddr:00:1a:2b:3c:4d:5e IPAddr:192.168.1.10}
}
