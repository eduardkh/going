package main

import (
	"fmt"
	"log"

	"github.com/bi-zone/wmi"
)

type win32Process struct {
	PID             uint32 `wmi:"ProcessId"`
	Name            string
	ParentProcessId uint32
	UserField       int `wmi:"-"`
}

func main() {
	var dst []win32Process

	q := wmi.CreateQueryFrom(&dst, "Win32_Process", "")
	// https://docs.microsoft.com/en-us/windows/win32/cimwin32prov/win32-process
	// fmt.Println(q)
	fmt.Println("\tPID\t\tName\t\tPPID")

	if err := wmi.Query(q, &dst); err != nil {
		log.Fatal(err)
	}
	for _, v := range dst {
		fmt.Printf("%6d\t%s\t%v\n", v.PID, v.Name, v.ParentProcessId)
	}
}
