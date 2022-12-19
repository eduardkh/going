package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// briefCmd represents the brief command
var briefCmd = &cobra.Command{
	Use:   "brief",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(briefDummy)
	},
}

var briefDummy string = `
Interface                  IP-Address      OK? Method Status                Protocol
GigabitEthernet0/0         unassigned      YES NVRAM  administratively down down
GigabitEthernet0/1         13.0.0.3        YES NVRAM  up                    up
GigabitEthernet0/2         unassigned      YES NVRAM  administratively down down
GigabitEthernet0/3         unassigned      YES NVRAM  administratively down down
Loopback0                  3.3.3.3         YES NVRAM  up                    up

`

func init() {
	interfaceCmd.AddCommand(briefCmd)
}
