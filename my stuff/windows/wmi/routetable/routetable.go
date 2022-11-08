package main

import (
	"fmt"
	"log"

	"github.com/bi-zone/wmi"
)

type Win32_IP4RouteTable struct {
	Name    string
	NextHop string //`wmi:"ProcessId"`
	Mask    string
	// PhysicalAdapter  bool
	// UserField       int `wmi:"-"`
}

func main() {
	var dst []Win32_IP4RouteTable

	q := wmi.CreateQueryFrom(&dst, "Win32_IP4RouteTable", "")
	// https://learn.microsoft.com/en-us/previous-versions/windows/desktop/wmiiprouteprov/win32-ip4routetable
	if err := wmi.Query(q, &dst); err != nil {
		log.Fatal(err)
	}
	for _, v := range dst {
		fmt.Printf("%v/%v via %v\n", v.Name, v.Mask, v.NextHop)
	}
}
