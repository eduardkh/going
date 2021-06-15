package main

import (
	"fmt"
	"net"
)

func main() {
	list, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	for _, iface := range list {
		// fmt.Printf("%d name=%s %v\n", i, iface.Name, iface)
		addrs, err := iface.Addrs()
		if err != nil {
			panic(err)
		}
		for j, addr := range addrs {
			if j == 0 {
				continue
			}
			fmt.Printf(" %d %v\n", j, addr)
		}
	}
}
