package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	var action string

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "act",
				Value:       "decode",
				Usage:       "encode or decode string to/from base64",
				Destination: &action,
			},
		},
		Action: func(c *cli.Context) error {
			name := "ZGVjb2RlIHNvbWV0aGluZw=="
			if c.NArg() > 0 {
				name = c.Args().Get(0)
			}
			if action == "encode" {
				fmt.Println(base64.StdEncoding.EncodeToString([]byte(name)))
			} else {
				data, err := base64.StdEncoding.DecodeString(name)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(string(data))
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
