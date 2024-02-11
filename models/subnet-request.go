package models

type SubnetRequest struct {
	MaskLength int    `json:"maskLength"`
	Name       string `json:"name"`
}
