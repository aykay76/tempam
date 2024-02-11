package services

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/aykay76/tempam/models"
	"github.com/aykay76/tempam/storage"
)

type NetworkService struct {
	store storage.Storage
}

func NewNetworkService(store storage.Storage) *NetworkService {
	return &NetworkService{store: store}
}

func (this *NetworkService) CreateNetwork(networkRequest models.NetworkRequest) *models.Network {
	// TODO: check for any overlaps and reject

	// find the last network
	names := this.ListNetworks()
	lastId := len(names)

	network := models.Network{
		ID:      lastId + 1,
		Name:    networkRequest.Name,
		Subnets: make([]*models.Subnet, 0),
	}

	network.CreatedAt = time.Now()
	network.UpdatedAt = time.Now()

	// TODO: put this into a database, not a local file
	bytes, err := json.Marshal(network)
	if err != nil {
		fmt.Println(err)
	}
	this.store.StoreBlob(fmt.Sprint("network-", network.ID, ".json"), bytes)

	return &network
}

func (this *NetworkService) CreateSubnet(networkId int, subnetRequest models.SubnetRequest) *models.Subnet {
	fmt.Printf("Creating subnet in network %d, with %d\n", networkId, subnetRequest.MaskLength)

	network := this.GetNetwork(networkId)
	if network == nil {
		return nil
	}

	subnet := network.GetNextAvailableSubnet(uint32(subnetRequest.MaskLength))
	if subnet == nil {
		fmt.Println("Could not find available address space")
		return nil
	}

	// created a valid subnet, update the rest of the fields and return to the controller
	subnet.Name = subnetRequest.Name
	subnet.CreatedAt = time.Now()
	subnet.UpdatedAt = time.Now()
	subnet.ID = len(network.Subnets) + 1
	network.Subnets = append(network.Subnets, subnet)

	this.UpdateNetwork(network)

	return subnet
}

func (s *NetworkService) ListNetworks() []string {
	names, err := s.store.ListBlobs("network-*.json")
	if err != nil {
		fmt.Println(err)
	}
	return names
}

func (s *NetworkService) GetAllNetworks() []models.Network {
	blobs, err := s.store.GetAllBlobs("network-*.json")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var networks []models.Network

	for _, blob := range blobs {
		var network models.Network
		json.Unmarshal(blob, &network)
		networks = append(networks, network)
	}

	return networks
}

func (s *NetworkService) GetNetwork(id int) *models.Network {
	fmt.Println("NetworkService.GetNetwork ", id)
	var network models.Network
	blob, err := s.store.GetBlob(fmt.Sprint("network-", id, ".json"))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	json.Unmarshal(blob, &network)
	return &network
}

func (s *NetworkService) UpdateNetwork(network *models.Network) {
	bytes, err := json.Marshal(network)
	if err != nil {
		fmt.Println(err)
	}
	s.store.StoreBlob(fmt.Sprint("network-", network.ID, ".json"), bytes)
}

func (s *NetworkService) DeleteNetwork(id int) {
	s.store.DeleteBlob(fmt.Sprint("network-", id, ".json"))
}
