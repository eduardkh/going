package main

import (
	"fmt"
	"log"

	"github.com/bi-zone/wmi"
)

type Win32_NetworkAdapterConfiguration struct {
	Description string
	IPAddress   []string //`wmi:"ProcessId"`
	MACAddress  string
	// PhysicalAdapter  bool
	// UserField       int `wmi:"-"`
}

func main() {
	var dst []Win32_NetworkAdapterConfiguration

	q := wmi.CreateQueryFrom(&dst, "Win32_NetworkAdapterConfiguration", "")
	// https://learn.microsoft.com/en-us/windows/win32/cimwin32prov/win32-networkadapterconfiguration

	if err := wmi.Query(q, &dst); err != nil {
		log.Fatal(err)
	}
	for _, v := range dst {
		if len(v.IPAddress) != 0 {
			fmt.Printf("%v - %v - %v\n", v.MACAddress, v.IPAddress[0], v.Description)
		}
	}
}
