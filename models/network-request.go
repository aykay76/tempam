package models

type NetworkRequest struct {
	Name string `json:"name"`
	CIDR string `json:"cidr"`
}
