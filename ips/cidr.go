package ips

import "regexp"

// Pack accepts a list of IPv4's, and returns a list of the smallest possible CIDR ranges.
func Pack(ipStrings []string) ([]string, error) {
	ipNums := make([]uint64, 0, len(ipStrings))
	for _, ip := range ipStrings {
		n, err := IPtoN(ip)
		if err != nil {
			return nil, err
		}

		ipNums = append(ipNums, n)
	}

	bounds, err := BlockBoundaries(ipNums)
	if err != nil {
		return nil, err
	}

	cidrs2 := make([]string, 0, 10)
	for _, bound := range bounds {
		cidrs2 = append(cidrs2, RangeToCIDR(bound))
	}

	return cidrs2, nil
}

// Unpack accepts a list of IPv4's and CIDR ranges, and returns a list of IPv4's and error.
func Unpack(cidrs []string) ([]string, error) {
	var ipStrings []string
	// Don't force the user to declare every range in CIDR
	// e.g. a range of /32 for a single IPv4
	for i, cidr := range cidrs {
		if b, err := regexp.MatchString("/\\d{1,2}$", cidr); !b && err == nil {
			cidrs[i] = cidr + "/32"
		} else if err != nil {
			return nil, err
		}
	}

	for _, cidr := range cidrs {
		hosts, err := Hosts(cidr)
		if err != nil {
			return nil, err
		}

		ipStrings = append(ipStrings, hosts...)
	}

	ipStrings = Uniq(ipStrings)

	return ipStrings, nil
}
