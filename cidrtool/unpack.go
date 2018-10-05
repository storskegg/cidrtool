package main

import (
	"fmt"
	"github.com/storskegg/cidrtool/fio"
	"github.com/storskegg/cidrtool/ips"
	"gopkg.in/urfave/cli.v1"
)

func actionUnpack(c *cli.Context) error {
	infile := c.Args().Get(0)
	if infile == "" {
		return fmt.Errorf("repack requires a valid path to an input file")
	}
	//outfile := c.Args().Get(1)

	cidrList, err := fio.ReadLinesFromFile(infile)
	if err != nil {
		return err
	}

	ipList, err := ips.Unpack(cidrList)
	if err != nil {
		return err
	}

	for _, ip := range ipList {
		fmt.Println(ip)
	}

	return nil
}
