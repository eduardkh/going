/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	_ "go-httpstat-cobra/pkg/sendmessage"

	"github.com/spf13/cobra"
)

// uriCmd represents the uri command
var uriCmd = &cobra.Command{
	Use:   "uri",
	Short: "get stats for uri",
	Long: `provide a single uri to scan and get the time stats for it.
you can send the stats to a syslog server it the syslog-server flag is added.
Example usage:
	go-httpstat-cobra uri https://www.google.co.il/ --syslog-server 192.168.1.155`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("uri called")
		arg := args[0]
		flg, _ := cmd.Flags().GetString("syslog-server")
		if len(flg) > 0 && len(args) > 0 {
			fmt.Printf("> user argument %q - flag %q\n", arg, flg)
			// sendmessage.SendMessage("tcp", "192.168.1.155", "514", "message")
		} else if len(args) > 0 {
			fmt.Printf("> user argument %q\n", arg)
		}
	},
}

func init() {
	rootCmd.AddCommand(uriCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	uriCmd.PersistentFlags().String("syslog-server", "", "pass in a syslog-server")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uriCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
