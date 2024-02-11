package models

import "net"

type NetworkRequest struct {
	Name string    `json:"name"`
	CIDR net.IPNet `json:"cidr"`
}
