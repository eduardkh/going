package cmd

import (
	"fmt"
	"log"

	"github.com/bi-zone/wmi"
	"github.com/spf13/cobra"
)

type gateway_win32_NetworkAdapterConfiguration struct {
	Description      string
	DefaultIPGateway []string
	IPAddress        []string
	IPSubnet         []string
}

// gatewayCmd represents the gateway command
var gatewayCmd = &cobra.Command{
	Use:   "gateway",
	Short: "get default gateway",
	Long:  `get default gateway`,
	Run: func(cmd *cobra.Command, args []string) {
		var dst []gateway_win32_NetworkAdapterConfiguration

		q := wmi.CreateQueryFrom(&dst, "Win32_NetworkAdapterConfiguration", "")
		if err := wmi.Query(q, &dst); err != nil {
			log.Fatal(err)
		}
		for _, v := range dst {
			if len(v.DefaultIPGateway) != 0 {
				fmt.Println("IP Address\tSubnet Mask\tDefault Gateway\tInterface Description")
				fmt.Printf("%v\t%v\t%v\t%v\n", v.IPAddress[0], v.IPSubnet[0], v.DefaultIPGateway[0], v.Description)
			}
		}
	},
}

func init() {
	ipCmd.AddCommand(gatewayCmd)
}
