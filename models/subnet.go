package models

import (
	"time"
)

// Subnet represents the structure for a subnet.
type Subnet struct {
	ID        int         `json:"id"`
	Name      string      `json:"name"`
	Range     IPRange     `json:"range"`
	CIDR      string      `json:"cidr"`
	IPs       []IPAddress `json:"ips"`
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt"`
}
