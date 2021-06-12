package main

import "fmt"

func main() {
	groups := [][]string{
		[]string{"10", "VLAN_10"},
		[]string{"20", "VLAN_20"},
		[]string{"30", "VLAN_30"},
		[]string{"40", "VLAN_40"},
		[]string{"50", "VLAN_50"},
		[]string{"60", "VLAN_60"},
		[]string{"70", "VLAN_70"}}
	for _, v := range groups {
		fmt.Println("vlan ", v[0], "\nname ", v[1])

	}
}
