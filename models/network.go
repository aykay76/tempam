package models

import (
	"fmt"
	"math"
	"time"
)

// A Network is a top level network with a collection of subnets
type Network struct {
	Name         string            `json:"name"`
	ID           int               `json:"id"`
	AddressSpace IPRange           `json:"cidr"`
	Subnets      []*Subnet         `json:"subnets"`
	Tags         map[string]string `json:"tags"`
	CreatedAt    time.Time         `json:"createdAt"`
	UpdatedAt    time.Time         `json:"updatedAt"`
}

func (this *Network) anyOverlaps(checkRange IPRange) bool {
	for _, subnet := range this.Subnets {
		// if there is an overlap
		if subnet.CIDR.StartAddress <= checkRange.EndAddress && subnet.CIDR.EndAddress >= checkRange.StartAddress {
			fmt.Printf("Found an overlap with subnet: %s\n", subnet.Name)
			return true
		}
	}

	return false
}

func (this *Network) GetNextAvailableSubnet(mask uint32) *Subnet {
	fmt.Println("> GetNextAvailableSubnet")
	fmt.Printf("mask: %d\n", mask)
	requestedSize := uint64(math.Pow(2, 32.0-float64(mask)))

	checkRange := IPRange{
		StartAddress: this.AddressSpace.StartAddress,
		EndAddress:   this.AddressSpace.StartAddress + requestedSize - 1,
	}
	fmt.Println(checkRange)

	// crude check for overlaps - could be optimised for sure!
	for this.anyOverlaps(checkRange) {
		checkRange.StartAddress += requestedSize
		checkRange.EndAddress += requestedSize
		fmt.Println("updated check range, trying again")
		fmt.Println(checkRange)
	}

	if checkRange.StartAddress >= this.AddressSpace.EndAddress {
		fmt.Println("Passed end of address space")
		fmt.Println("< GetNextAvailableSubnet")
		return nil
	}

	// to get this far we found an available subnet range
	subnet := Subnet{
		CIDR: IPRange{
			StartAddress: checkRange.StartAddress,
			EndAddress:   checkRange.EndAddress,
		},
	}

	fmt.Println("< GetNextAvailableSubnet")
	return &subnet
}
