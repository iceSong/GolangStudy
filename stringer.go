package main

import (
	"fmt"
)

type IPAdddr [4]byte

func (i IPAdddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", i[0],i[1],i[2],i[3])
}

func main() {
	addrs := map[string]IPAdddr {
		"loopback":{127,0,0,1},
		"googleDNS":{8,8,8,8},
	}

	for n,a := range addrs {
		fmt.Printf("%v: %v\n", n, a)
	}
}