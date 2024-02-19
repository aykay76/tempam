package models

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"net"
	"strconv"
)

type IPRange struct {
	StartAddress uint64
	EndAddress   uint64
}

func ip2int(ip string) uint64 {
	var long uint64
	binary.Read(bytes.NewBuffer(net.ParseIP(ip).To4()), binary.BigEndian, &long)
	return long
}

func (rnge *IPRange) FromCIDR(cidr string) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ip)
	ones, _ := ipnet.Mask.Size()
	var temp uint32
	binary.Read(bytes.NewBuffer(ip.To4()), binary.BigEndian, &temp)
	rnge.StartAddress = uint64(temp)
	rnge.EndAddress = rnge.StartAddress + uint64(math.Pow(2, float64(32.0-ones))) - 1
}

// Does this IP range overlap with another IP range?
func (rnge *IPRange) Overlaps(other IPRange) bool {
	if rnge.StartAddress <= other.EndAddress && rnge.EndAddress >= other.StartAddress {
		return true
	}

	return false
}

func (rnge *IPRange) String() string {
	maskLength := 32 - uint32(math.Log2(float64(rnge.EndAddress-rnge.StartAddress+1)))

	b0 := strconv.FormatUint((rnge.StartAddress>>24)&0xff, 10)
	b1 := strconv.FormatUint((rnge.StartAddress>>16)&0xff, 10)
	b2 := strconv.FormatUint((rnge.StartAddress>>8)&0xff, 10)
	b3 := strconv.FormatUint((rnge.StartAddress & 0xff), 10)

	return fmt.Sprintf("%s.%s.%s.%s/%d", b0, b1, b2, b3, maskLength)
}
