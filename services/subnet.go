package services

import (
	"encoding/json"
	"fmt"

	"github.com/aykay76/tempam/models"
	"github.com/aykay76/tempam/storage"
)

type SubnetService struct {
	store storage.Storage
}

func NewSubnetService(store storage.Storage) *SubnetService {
	return &SubnetService{store: store}
}

func (s *SubnetService) CreateSubnet(subnet models.Subnet) models.Subnet {
	// given the network ID search for space for the requested subnet

	bytes, err := json.Marshal(subnet)
	if err != nil {
		fmt.Println(err)
	}
	s.store.StoreBlob(fmt.Sprint("subnet-", subnet.ID, ".json"), bytes)

	return models.Subnet{}
}

func (s *SubnetService) ListSubnets() []string {
	names, err := s.store.ListBlobs("subnet-*.json")
	if err != nil {
		fmt.Println(err)
	}
	return names
}

func (s *SubnetService) GetAllSubnets() []models.Subnet {
	blobs, err := s.store.GetAllBlobs("subnet-*.json")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var subnets []models.Subnet

	for _, blob := range blobs {
		var subnet models.Subnet
		json.Unmarshal(blob, &subnet)
		subnets = append(subnets, subnet)
	}

	return subnets
}

func (s *SubnetService) GetSubnet(id int) *models.Subnet {
	fmt.Println("SubnetService.GetSubnet ", id)
	var subnet models.Subnet
	blob, err := s.store.GetBlob(fmt.Sprint("subnet-", id, ".json"))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	json.Unmarshal(blob, &subnet)
	return &subnet
}

func (s *SubnetService) UpdateSubnet(subnet models.Subnet) {
	bytes, err := json.Marshal(subnet)
	if err != nil {
		fmt.Println(err)
	}
	s.store.StoreBlob(fmt.Sprint("subnet-", subnet.ID, ".json"), bytes)
}

func (s *SubnetService) DeleteSubnet(id int) {
	s.store.DeleteBlob(fmt.Sprint("subnet-", id, ".json"))
}
