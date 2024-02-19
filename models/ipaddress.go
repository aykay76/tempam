package models

import (
	"fmt"
	"strconv"
)

// IPAddress represents the structure for an IP address.
type IPAddress struct {
	Address uint64 `json:"address"`
}

func (ip *IPAddress) String() string {
	b0 := strconv.FormatUint((ip.Address>>24)&0xff, 10)
	b1 := strconv.FormatUint((ip.Address>>16)&0xff, 10)
	b2 := strconv.FormatUint((ip.Address>>8)&0xff, 10)
	b3 := strconv.FormatUint((ip.Address & 0xff), 10)

	return fmt.Sprintf("%s.%s.%s.%s", b0, b1, b2, b3)
}
