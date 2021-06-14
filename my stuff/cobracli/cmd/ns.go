package cmd

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"
)

// https://tutorialedge.net/golang/building-a-cli-in-go/
// nsCmd represents the ns command
var nsCmd = &cobra.Command{
	Use:   "ns",
	Short: "Looks Up the NameServers for a Particular Host",
	Long: `Looks Up the NameServers for a Particular Host
example usage: cobracli ns -n google.com`,
	Run: func(cmd *cobra.Command, args []string) {
		ns, err := net.LookupNS(hostname)
		if err != nil {
			fmt.Println(err)
		}
		for i := 0; i < len(ns); i++ {
			fmt.Println(ns[i].Host)
		}
	},
}

// https://stackoverflow.com/questions/57181589/failed-to-get-flag-value
var hostname string

func init() {
	rootCmd.AddCommand(nsCmd)
	nsCmd.Flags().StringVarP(&hostname, "hostname", "n", "google.co.il", "usage")

}
