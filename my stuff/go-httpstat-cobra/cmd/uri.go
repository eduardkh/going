/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"go-httpstat-cobra/pkg/gethttpstat"
	"go-httpstat-cobra/pkg/sendmessage"
	"time"

	"github.com/spf13/cobra"
)

var long string = `
provide a single uri to scan and get the time stats for it.
you can send the stats to a syslog server if the syslog-server flag is added.

Example usage:
	go-httpstat-cobra uri https://www.google.co.il/ --syslog-server 192.168.1.155
`

// uriCmd represents the uri command
var uriCmd = &cobra.Command{
	Use:   "uri",
	Short: "get stats for uri",
	Long:  long,
	Run: func(cmd *cobra.Command, args []string) {
		flg, _ := cmd.Flags().GetString("syslog-server")
		if len(flg) > 0 && len(args) > 0 {
			// if an arg and a flag given - send results to syslog
			arg := args[0]
			result := gethttpstat.Gethttpstat(arg)
			t := time.Now()
			tf := t.Format("Jan 02 3:04:01")
			message := fmt.Sprintf("%v CEF:0|go-httpstat-cobra|go-httpstat-cobra|GHC0.1|method=GET request=%v NameLookup=%v Connect=%v PreTransfer=%v StartTransfer=%v Total=%v", tf, arg, result.NameLookup, result.Connect, result.Pretransfer, result.StartTransfer, result.Totall)
			sendmessage.SendMessage("tcp", flg, "514", message)
		} else if len(args) > 0 {
			// if only an arg is given - print results to screen
			arg := args[0]
			result := gethttpstat.Gethttpstat(arg)
			fmt.Printf("Name Lookup :\t%v\nConnect:\t%v\nPre-Transfer:\t%v\nStart Transfer:\t%v\nTotal:\t\t%v\n", result.NameLookup, result.Connect, result.Pretransfer, result.StartTransfer, result.Totall)
		} else {
			// if no args or flags given - print help
			fmt.Println(long)
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
