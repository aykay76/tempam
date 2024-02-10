package models

// Subnet represents the structure for a subnet.
type Subnet struct {
	ID      string      `json:"id"`
	CIDR    IPRange     `json:"cidr"`
	Subnets []Subnet    `json:"subnets"`
	IPs     []IPAddress `json:"ips"`
}
