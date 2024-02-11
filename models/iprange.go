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

func (this *IPRange) FromCIDR(cidr string) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err == nil {
		ip.To4()
		ones, _ := ipnet.Mask.Size()
		this.StartAddress = ip2int(ip.String())
		this.EndAddress = this.StartAddress + uint64(math.Pow(2, float64(32.0-ones))) - 1
	}
}

// Does this IP range overlap with another IP range?
func (this *IPRange) Overlaps(other IPRange) bool {
	if this.StartAddress <= other.EndAddress && this.EndAddress >= other.StartAddress {
		return true
	}

	return false
}

func (this *IPRange) String() string {
	maskLength := 32 - uint32(math.Log2(float64(this.EndAddress-this.StartAddress+1)))

	b0 := strconv.FormatUint((this.StartAddress>>24)&0xff, 10)
	b1 := strconv.FormatUint((this.StartAddress>>16)&0xff, 10)
	b2 := strconv.FormatUint((this.StartAddress>>8)&0xff, 10)
	b3 := strconv.FormatUint((this.StartAddress & 0xff), 10)

	return fmt.Sprintf("%s.%s.%s.%s/%d", b0, b1, b2, b3, maskLength)
}
