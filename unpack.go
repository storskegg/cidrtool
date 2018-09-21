package main

import (
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"keybase/cidrtool/fio"
	ips2 "keybase/cidrtool/ips"
	"regexp"
)

func actionUnpack(c *cli.Context) error {
	infile := c.Args().Get(0)
	if infile == "" {
		return fmt.Errorf("repack requires a valid path to an input file")
	}
	//outfile := c.Args().Get(1)

	cidrs, err := fio.ReadLinesFromFile(infile)
	if err != nil {
		return err
	}

	ips, err := unpack(cidrs)
	if err != nil {
		return err
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

	return nil
}

func unpack(cidrs []string) ([]string, error) {
	var ipStrings []string
	// Don't force the user to declare every range in CIDR
	// e.g. a range of /32 for a single IPv4
	for i, cidr := range cidrs {
		if b, err := regexp.MatchString("/\\d{1,2}$", cidr); !b && err != nil {
			cidrs[i] = cidr + "/32"
		}
	}

	for _, cidr := range cidrs {
		hosts, err := ips2.Hosts(cidr)
		if err != nil {
			return nil, err
		}

		ipStrings = append(ipStrings, hosts...)
	}

	ipStrings = ips2.Uniq(ipStrings)

	return ipStrings, nil
}
