package cmd

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"
)

// interfaceCmd represents the interface command
var interfaceCmd = &cobra.Command{
	Use:   "interface",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		interfaces, err := net.Interfaces()
		// handle err
		if err != nil {
			fmt.Println(err)
		}
		var nameInterf, nameAddress, nameMAC, nameMTU string = "Interface", "IP-Addres", "MAC-Address", "MTU"
		fmt.Printf("%-34v %-20v %-17v %-v\n", nameInterf, nameAddress, nameMAC, nameMTU)
		for _, iface := range interfaces {
			if iface.Flags&net.FlagUp == 0 {
				continue // interface down
			}
			if iface.Flags&net.FlagLoopback != 0 {
				continue // loopback interface
			}
			addrs, err := iface.Addrs()
			if err != nil {
				panic(err)
			}
			fmt.Printf("%-34v %-20v %-v %v\n", iface.Name, addrs[1], iface.HardwareAddr, iface.MTU)
		}
	},
}

func init() {
	ipCmd.AddCommand(interfaceCmd)
}
