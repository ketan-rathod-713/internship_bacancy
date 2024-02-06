package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
// TODO: great
func (ip IPAddr) String() string {
	s := ""
	for i, val := range ip {
		s += fmt.Sprintf("%v", val)
		if i < len(ip)-1 {
			s += fmt.Sprintf("%v", ".")
		}
	}

	return s
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
