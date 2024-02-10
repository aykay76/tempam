package models

import (
	"fmt"
	"math"
	"strconv"
)

type IPRange struct {
	StartAddress uint64
	EndAddress   uint64
	CIDR         string
}

// Does this IP range overlap with another IP range?
func (this *IPRange) Overlaps(other IPRange) bool {
	if this.StartAddress <= other.EndAddress || this.EndAddress >= other.StartAddress {
		return true
	}

	return false
}

func (this *IPRange) String() string {
	maskLength := 32 - uint32(math.Log2(float64(this.EndAddress-this.StartAddress)))

	b0 := strconv.FormatUint((this.StartAddress>>24)&0xff, 10)
	b1 := strconv.FormatUint((this.StartAddress>>16)&0xff, 10)
	b2 := strconv.FormatUint((this.StartAddress>>8)&0xff, 10)
	b3 := strconv.FormatUint((this.StartAddress & 0xff), 10)

	return fmt.Sprintf("%s.%s.%s.%s/%d", b0, b1, b2, b3, maskLength)
}
