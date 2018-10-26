package main

import (
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"log"
	"os"
	"time"
)

func main() {
	app := cli.NewApp()

	app.Name = "cidrtool"
	app.Usage = "A small application to make manipulating CIDR notation easier"
	app.Description = "cidrtool is a small application to make manipulating CIDR notation a little bit\neasier. For now, it takes a list of IPv4 addresses and CIDR ranges, removes collisions,\nand returns the smallest possible set of CIDRs."
	app.Version = "0.2.1"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		{
			Name:  "William Conrad",
			Email: "liam@storskegg.org",
		},
	}
	app.Copyright = "(c) 2018 William Conrad"
	app.Commands = []cli.Command{
		{
			Name:    "pack",
			Aliases: []string{"p"},
			Usage:   "pack a list of ips into narrowest cidr notations",
			Action:  actionPack,
		},
		{
			Name:    "repack",
			Aliases: []string{"r"},
			Usage:   "repack a list of ips and cidrs into narrowest cidr notations",
			Action:  actionRepack,
		},
		{
			Name:    "unpack",
			Aliases: []string{"u"},
			Usage:   "unpack a list of ips and cidrs into sorted, deduped list of IPv4s",
			Action:  actionUnpack,
		},
		{
			Name:   "srv",
			Usage:  "start a server providing a web interface to all commands. ignores --db flag",
			Action: actionServe,
		},
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "nostats",
			Usage: "Do not display stats",
		},
	}

	app.CommandNotFound = func(c *cli.Context, command string) {
		_, err := fmt.Fprintf(c.App.Writer, "Unrecognized command: %q\n", command)
		checkErr(err)
	}

	err := app.Run(os.Args)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
