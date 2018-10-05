package main

import (
	"fmt"
	"github.com/storskegg/cidrtool/fio"
	"github.com/storskegg/cidrtool/ips"
	"gopkg.in/urfave/cli.v1"
)

func actionRepack(c *cli.Context) error {
	infile := c.Args().Get(0)
	if infile == "" {
		return fmt.Errorf("repack requires a valid path to an input file")
	}

	cidrList, err := fio.ReadLinesFromFile(infile)
	if err != nil {
		return err
	}

	ipList, err := ips.Unpack(cidrList)
	if err != nil {
		return err
	}

	cidrListPacked, err := ips.Pack(ipList)
	if err != nil {
		return err
	}

	for _, c := range cidrListPacked {
		fmt.Println(c)
	}

	if c.BoolT("nostats") {
		fmt.Println("Num CIDRs Analyzed:", len(cidrList))
		fmt.Println("Num CIDRs Resulted:", len(cidrListPacked))
	}

	return nil
}
