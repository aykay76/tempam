package models

// A Network is a top level network with a collection of subnets
type Network struct {
	Name    string            `json:"name"`
	ID      int               `json:"id"`
	CIDR    string            `json:"cidr"`
	Subnets []Subnet          `json:"subnets"`
	Tags    map[string]string `json:"tags"`
}
