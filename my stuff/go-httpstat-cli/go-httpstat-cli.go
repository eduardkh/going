package main

import (
	"log"
	"os"

	"go-httpstat-cli/pkg/sendmessage"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "http-stats",
		Usage: "get HTTP Stats for any website",
		Authors: []*cli.Author{
			{
				Name: "Eduard Khiaev",
				// Email: "",
			},
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:        "uri",
			Usage:       "get stats for specified URI",
			Description: "Example: http-stats uri https://google.com",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "syslog-server",
					Usage: "send syslog to server",
				},
			},
			Action: func(c *cli.Context) error {
				// get flag data
				flagData := c.String("syslog-server")
				log.Println("> stats for http-stats uri https://google.com")
				sendmessage.SendMessage("tcp", "192.168.1.155", "80", "message")

				// check if flag set (does not return "")
				if flagData != "" {
					log.Println(">", flagData)
				}
				return nil
			},
		},
		{
			Name:        "uri-file",
			Usage:       "get stats for specified URI list from File",
			Description: "Example: http-stats uri https://google.com",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "syslog-server",
					Usage: "send syslog to server",
				},
			},
			Action: func(c *cli.Context) error {
				// get flag data
				flagData := c.String("syslog-server")
				log.Println("> stats for http-stats uri https://google.com")

				// check if flag set (does not return "")
				if flagData != "" {
					log.Println(">", flagData)
				}

				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
