package main

import (
	"os"

	cli "gopkg.in/urfave/cli.v1"
)

var Name = "dynamotool"
var Version = "0.1.0"
var Usage = "nifty tool to operate DynamoDB"

func main() {
	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Usage = Usage

	app.Commands = []cli.Command{
		{
			Name:    "export",
			Aliases: []string{"e"},
			Usage:   "export data from DynamoDB table",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name:    "import",
			Aliases: []string{"i"},
			Usage:   "import data to DynamoDB table",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name:  "erase",
			Usage: "erase DynamoDB table. DANGER!!!",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
	}

	// TODO: credentials
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "format, f",
			Value: "csv",
			Usage: "Sets `FORMAT` for import and export operations",
		},
		cli.StringFlag{
			Name:  "separator, s",
			Value: "|",
			Usage: "`SEPARATOR` symbol for CSV files",
		},
		cli.StringFlag{
			Name:  "readlimit, rl",
			Value: "0",
			Usage: "read throughput to avoid DynamoDB throttling",
		},
		cli.StringFlag{
			Name:  "writelimit, wl",
			Value: "0",
			Usage: "read throughput to avoid DynamoDB throttling",
		},
	}

	app.Run(os.Args)
}
