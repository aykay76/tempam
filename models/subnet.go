package models

// Subnet represents the structure for a subnet.
type Subnet struct {
	ID   string      `json:"id"`
	CIDR string      `json:"cidr"`
	IPs  []IPAddress `json:"ips"`
}
