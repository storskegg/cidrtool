package main

import (
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"keybase/cidrtool/fio"
)

func actionRepack(c *cli.Context) error {
	infile := c.Args().Get(0)
	if infile == "" {
		return fmt.Errorf("repack requires a valid path to an input file")
	}
	//outfile := c.Args().Get(1)

	cidrs, err := fio.ReadLinesFromFile(infile)
	if err != nil {
		return err
	}

	ipStrings, err := unpack(cidrs)
	if err != nil {
		return err
	}

	cidrs2, err := pack(ipStrings)
	if err != nil {
		return err
	}

	for _, c := range cidrs2 {
		fmt.Println(c)
	}

	if c.BoolT("nostats") {
		fmt.Println("Num CIDRs Analyzed:", len(cidrs))
		fmt.Println("Num CIDRs Resulted:", len(cidrs2))
	}

	if c.GlobalBool("db") {
		err = pushToDB(cidrs2)
		if err != nil {
			return nil
		}
	}

	return nil
}
