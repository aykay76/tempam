package models

import "math"

// A Network is a top level network with a collection of subnets
type Network struct {
	Name         string            `json:"name"`
	ID           int               `json:"id"`
	AddressSpace IPRange           `json:"cidr"`
	Subnets      []Subnet          `json:"subnets"`
	Tags         map[string]string `json:"tags"`
}

func (this *Network) anyOverlaps(checkRange IPRange) bool {
	for _, subnet := range this.Subnets {
		// if there is an overlap
		if subnet.CIDR.StartAddress <= checkRange.EndAddress || subnet.CIDR.EndAddress >= checkRange.StartAddress {
			return true
		}
	}

	return false
}

func (this *Network) GetNextAvailableSubnet(mask uint32) Subnet {
	requestedSize := uint64(math.Pow(2, 32.0-float64(mask)))

	checkRange := IPRange{
		StartAddress: this.AddressSpace.StartAddress,
		EndAddress:   this.AddressSpace.StartAddress + requestedSize,
	}

	// crude check for overlaps - could be optimised for sure!
	for this.anyOverlaps(checkRange) {
		checkRange.StartAddress += requestedSize
		checkRange.EndAddress += requestedSize
	}

	// to get this far we either found a free space or we have got to the end of our address space and there is no free subnet
	subnet := Subnet{
		CIDR: IPRange{
			StartAddress: checkRange.StartAddress,
			EndAddress:   checkRange.EndAddress,
		},
	}

	// add the new subnet to the list of subnets
	this.Subnets = append(this.Subnets, subnet)

	return subnet
}
