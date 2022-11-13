/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"go-httpstat-cobra/pkg/gethttpstat"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var longList string = `provide a uri-list in "uri-list.txt" file to scan and get the time stats for it.
you can send the stats to a syslog server it the syslog-server flag is added.
Example usage:
	go-httpstat-cobra uriList --syslog-server 192.168.1.155`

// uriListCmd represents the uriList command
var uriListCmd = &cobra.Command{
	Use:   "uriList",
	Short: "get stats for uri-list.txt",
	Long:  longList,
	Run: func(cmd *cobra.Command, args []string) {
		// https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
		file, err := os.Open("uri-list.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			log.Println(scanner.Text())
			gethttpstat.Gethttpstat(scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(uriListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uriListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uriListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
