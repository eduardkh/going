package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// hostnameCmd represents the hostname command
var hostnameCmd = &cobra.Command{
	Use:   "hostname",
	Short: "Gets OS hostname",
	Long: `get the current OS hostname
example usage: cobracli hostname
`,
	Run: func(cmd *cobra.Command, args []string) {
		name, err := os.Hostname()
		if err != nil {
			panic(err)
		}

		fmt.Println("hostname:", name)
	},
}

func init() {
	rootCmd.AddCommand(hostnameCmd)
}
