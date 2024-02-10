package models

import (
	"fmt"
	"strconv"
)

// IPAddress represents the structure for an IP address.
type IPAddress struct {
	Address uint64 `json:"address"`
}

func (this *IPAddress) String() string {
	b0 := strconv.FormatUint((this.Address>>24)&0xff, 10)
	b1 := strconv.FormatUint((this.Address>>16)&0xff, 10)
	b2 := strconv.FormatUint((this.Address>>8)&0xff, 10)
	b3 := strconv.FormatUint((this.Address & 0xff), 10)

	return fmt.Sprintf("%s.%s.%s.%s", b0, b1, b2, b3)
}
