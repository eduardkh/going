package cmd

import (
	"fmt"
	"log"

	"github.com/bi-zone/wmi"
	"github.com/spf13/cobra"
)

type interface_win32_NetworkAdapterConfiguration struct {
	Description      string
	DefaultIPGateway []string
	IPAddress        []string
	IPSubnet         []string
	MACAddress       string
	TcpWindowSize    int
	IPEnabled        bool
}

// interfaceCmd represents the interface command
var interfaceCmd = &cobra.Command{
	Use:   "interface",
	Short: "get a detailed interface information",
	Long:  `get a detailed interface information`,
	Run: func(cmd *cobra.Command, args []string) {
		var dst []interface_win32_NetworkAdapterConfiguration

		q := wmi.CreateQueryFrom(&dst, "Win32_NetworkAdapterConfiguration", "")
		if err := wmi.Query(q, &dst); err != nil {
			log.Fatal(err)
		}
		for _, v := range dst {
			if v.IPEnabled {
				fmt.Printf("Interface Description: %v\n", v.Description)
				fmt.Printf("IPv4 Address:          %v\n", v.IPAddress[0])
				fmt.Printf("IPv4 Subnet Mask: %18v\n", v.IPSubnet[0])
				fmt.Printf("MAC Address: %27v\n", v.MACAddress)
				fmt.Printf("TCP Window Size:%12v\n\n", v.TcpWindowSize)
			}
		}
	},
}

func init() {
	ipCmd.AddCommand(interfaceCmd)
}
