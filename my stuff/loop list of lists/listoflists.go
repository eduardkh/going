package main

import "fmt"

func main() {
	groups := [][2]string{
		[2]string{"10", "VLAN_10"},
		[2]string{"30", "VLAN_30"},
		[2]string{"40", "VLAN_40"},
		[2]string{"20", "VLAN_20"},
		[2]string{"50", "VLAN_50"},
		[2]string{"60", "VLAN_60"},
		[2]string{"70", "VLAN_70"}}
	for _, v := range groups {
		fmt.Printf("vlan %v\n", v[0])
		fmt.Printf("name %v\n", v[1])

	}
}
