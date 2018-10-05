package main

import (
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"keybase/cidrtool/fio"
	ips2 "keybase/cidrtool/ips"
)

func actionPack(c *cli.Context) error {
	infile := c.Args().Get(0)
	if infile == "" {
		return fmt.Errorf("repack requires a valid path to an input file")
	}
	//outfile := c.Args().Get(1)

	ips, err := fio.ReadLinesFromFile(infile)
	if err != nil {
		return err
	}

	cidrs, err := pack(ips)
	if err != nil {
		return err
	}

	for _, cidr := range cidrs {
		fmt.Println(cidr)
	}

	return nil
}

func pack(ipStrings []string) ([]string, error) {
	ipNums := make([]uint64, 0, len(ipStrings))
	for _, ip := range ipStrings {
		n, err := ips2.IPtoN(ip)
		if err != nil {
			return nil, err
		}

		ipNums = append(ipNums, n)
	}

	bounds, err := ips2.BlockBoundaries(ipNums)
	if err != nil {
		return nil, err
	}

	cidrs2 := make([]string, 0, 10)
	for _, bound := range bounds {
		cidrs2 = append(cidrs2, ips2.RangeToCIDR(bound))
	}

	return cidrs2, nil
}
