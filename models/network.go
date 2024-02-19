package models

import (
	"fmt"
	"math"
	"time"
)

// A Network is a top level network with a collection of subnets
type Network struct {
	Name         string            `json:"name,omitempty" bson:"name,omitempty"`
	ID           int               `json:"id,omitempty" bson:"id,omitempty"`
	AddressSpace IPRange           `json:"range,omitempty" bson:"range,omitempty"`
	CIDR         string            `json:"cidr,omitempty" bson:"cidr,omitempty"`
	Subnets      []*Subnet         `json:"subnets,omitempty" bson:"subnets,omitempty"`
	Tags         map[string]string `json:"tags,omitempty" bson:"tags,omitempty"`
	CreatedAt    time.Time         `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt    time.Time         `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

func (network *Network) anyOverlaps(checkRange IPRange) bool {
	for _, subnet := range network.Subnets {
		// if there is an overlap
		if subnet.Range.StartAddress <= checkRange.EndAddress && subnet.Range.EndAddress >= checkRange.StartAddress {
			fmt.Printf("Found an overlap with subnet: %s\n", subnet.Name)
			return true
		}
	}

	return false
}

func (network *Network) GetNextAvailableSubnet(mask uint32) *Subnet {
	fmt.Println("> GetNextAvailableSubnet")
	fmt.Printf("mask: %d\n", mask)
	requestedSize := uint64(math.Pow(2, 32.0-float64(mask)))

	checkRange := IPRange{
		StartAddress: network.AddressSpace.StartAddress,
		EndAddress:   network.AddressSpace.StartAddress + requestedSize - 1,
	}
	fmt.Println(checkRange)

	// crude check for overlaps - could be optimised for sure!
	for network.anyOverlaps(checkRange) {
		checkRange.StartAddress += requestedSize
		checkRange.EndAddress += requestedSize
		fmt.Println("updated check range, trying again")
		fmt.Println(checkRange)
	}

	if checkRange.StartAddress >= network.AddressSpace.EndAddress {
		fmt.Println("Passed end of address space")
		fmt.Println("< GetNextAvailableSubnet")
		return nil
	}

	// to get this far we found an available subnet range
	subnet := Subnet{
		Range: IPRange{
			StartAddress: checkRange.StartAddress,
			EndAddress:   checkRange.EndAddress,
		},
	}
	subnet.CIDR = subnet.Range.String()

	fmt.Println("< GetNextAvailableSubnet")
	return &subnet
}
