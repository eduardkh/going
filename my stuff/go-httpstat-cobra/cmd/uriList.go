/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var longList string = `
provide a uri-list in "uri-list.txt" file to scan and get the time stats for it.
you can send the stats to a syslog server it the syslog-server flag is added.

Example usage:
	go-httpstat-cobra uriList uri-list.txt --syslog-server 192.168.1.155`

// uriListCmd represents the uriList command
var uriListCmd = &cobra.Command{
	Use:   "uriList",
	Short: "get stats for uri-list.txt",
	Long:  longList,
	Run: func(cmd *cobra.Command, args []string) {
		flg, _ := cmd.Flags().GetString("syslog-server")
		if len(flg) > 0 && len(args) > 0 {
			// if an arg and a flag given - send results to syslog
			arg := args[0]
			fmt.Println(arg, flg)
		} else if len(args) > 0 {
			// if only an arg is given - print results to screen
			arg := args[0]
			fmt.Println(arg)
		} else {
			// if no args or flags given - print help
			fmt.Println(longList)
		}
	},
}

func init() {
	rootCmd.AddCommand(uriListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	uriListCmd.PersistentFlags().String("syslog-server", "", "pass in a syslog-server")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uriListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
