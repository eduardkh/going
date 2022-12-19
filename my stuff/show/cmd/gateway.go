package cmd

import (
	"fmt"
	"log"

	"github.com/bi-zone/wmi"
	"github.com/spf13/cobra"
)

type Win32_NetworkAdapterConfiguration struct {
	Description      string
	DefaultIPGateway []string
	IPAddress        []string
	IPSubnet         []string
}

// gatewayCmd represents the gateway command
var gatewayCmd = &cobra.Command{
	Use:   "gateway",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("gateway called")
		var dst []Win32_NetworkAdapterConfiguration

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
