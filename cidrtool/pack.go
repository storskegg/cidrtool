package main

import (
	"fmt"
	"github.com/storskegg/cidrtool/fio"
	"github.com/storskegg/cidrtool/ips"
	"gopkg.in/urfave/cli.v1"
)

func actionPack(c *cli.Context) error {
	infile := c.Args().Get(0)
	if infile == "" {
		return fmt.Errorf("repack requires a valid path to an input file")
	}

	ipList, err := fio.ReadLinesFromFile(infile)
	if err != nil {
		return err
	}

	cidrList, err := ips.Pack(ipList)
	if err != nil {
		return err
	}

	for _, cidr := range cidrList {
		fmt.Println(cidr)
	}

	return nil
}
