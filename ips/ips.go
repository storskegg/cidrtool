package ips

import (
	"fmt"
	strings2 "github.com/storskegg/cidrtool/strings"
	"net"
	"sort"
	"strconv"
	"strings"
)

// IPRange is a type representing a block of consecutive IPv4 addresses
type IPRange struct {
	Start, End string
}

// BlockBoundaries determines from a list of ip addresses the start and end of consecutive blocks
func BlockBoundaries(ipns []uint64) ([]IPRange, error) {
	blockBoundaries := make([]IPRange, 0, 50)
	var startSet, endSet bool
	var blockStart, blockEnd uint64
	sort.Slice(ipns, func(i, j int) bool { return ipns[i] < ipns[j] })
	ln := len(ipns)
	for i, n := range ipns {
		if i == 0 {
			blockStart = n
			startSet = true
		} else if ipns[i-1] != n-1 {
			blockStart = n
			startSet = true
		}

		if i+1 == ln {
			blockEnd = n
			endSet = true
		} else if ipns[i+1] != n+1 {
			blockEnd = n
			endSet = true
		}

		if startSet && endSet {
			start, err := NtoIP(blockStart)
			if err != nil {
				return nil, err
			}

			end, err := NtoIP(blockEnd)
			if err != nil {
				return nil, err
			}
			blockBoundaries = append(blockBoundaries, IPRange{start, end})
			startSet = false
			endSet = false
		}
	}
	return blockBoundaries, nil
}

// Hosts unpacks a CIDR address, and returns a string slice of IPv4's
func Hosts(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ipX := ip.Mask(ipnet.Mask); ipnet.Contains(ipX); inc(ipX) {
		ips = append(ips, ipX.String())
	}
	if len(ips) == 1 {
		return ips, nil
	}
	// remove network address and broadcast address
	return ips[1 : len(ips)-1], nil
}

// IPtoN accepts a valid IPv4 address, and returns its equivalent int64
func IPtoN(str string) (uint64, error) {
	octets := strings.Split(str, ".")
	if len(octets) != 4 {
		return 0, fmt.Errorf("str must be a valid ipv4 address")
	}
	hexStr := ""
	for _, ns := range octets {
		n, err := strconv.ParseInt(ns, 10, 64)
		if err != nil {
			return 0, err
		}
		hexStr = hexStr + fmt.Sprintf("%02x", n)
	}
	return strconv.ParseUint(hexStr, 16, 64)
}

// NtoIP accepts an int64, and returns an IPv4 string
func NtoIP(n uint64) (string, error) {
	hexStr := fmt.Sprintf("%08x", n)
	octetsH := strings2.SplitByN(hexStr, 2)
	octetsI := make([]string, 0, 4)
	for _, h := range octetsH {
		i, err := strconv.ParseInt(h, 16, 64)
		if err != nil {
			return "", err
		}
		octetsI = append(octetsI, fmt.Sprintf("%v", i))
	}

	return strings.Join(octetsI, "."), nil
}

// RangeToCIDR accepts an IPRange, and returns a CIDR notation IP range
func RangeToCIDR(ipRange IPRange) (cidr string) {
	bits := 32

	ip1 := net.ParseIP(ipRange.Start)
	ip2 := net.ParseIP(ipRange.End)

	for b := bits; b >= 0; b-- {
		mask := net.CIDRMask(b, bits)
		na := ip1.Mask(mask)
		n := net.IPNet{IP: na, Mask: mask}

		if n.Contains(ip2) {
			cidr = fmt.Sprintf("%v/%v", na, b)
			break
		}
	}

	return cidr
}

// Uniq returns a deduped slice of strings
func Uniq(ips []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	for _, ip := range ips {
		if _, val := keys[ip]; !val {
			keys[ip] = true
			list = append(list, ip)
		}
	}

	return list
}

//  http://play.golang.org/p/m8TNTtygK0
func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
