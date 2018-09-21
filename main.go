package main

import (
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()

	app.Name = "cidrtool"
	app.Usage = "A small application to make manipulating CIDR notation easier"
	app.Description = "cidrtool is a small application to make manipulating CIDR notation a little bit\neasier. For now, it takes a list of IPv4 addresses and CIDR ranges, removes collisions,\nand returns the smallest possible set of CIDRs."
	app.Version = "0.2.0"
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
			Name:  "db",
			Usage: "Push to PostgresDB",
		},
		cli.BoolTFlag{
			Name:  "nostats",
			Usage: "Do not display stats",
		},
	}

	app.Before = checkDBEnv

	err := app.Run(os.Args)
	checkErr(err)
}

func checkDBEnv(c *cli.Context) error {
	if c.GlobalBool("db") {
		pgURI := os.Getenv("RA_PG_URI")
		if pgURI == "" {
			return fmt.Errorf("global flag --db requires a valid postgres uri set in env var RA_PG_URI")
		}
	}
	return nil
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
