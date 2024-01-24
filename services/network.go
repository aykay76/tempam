package services

import (
	"encoding/json"
	"fmt"

	"github.com/aykay76/tempam/models"
	"github.com/aykay76/tempam/storage"
)

type NetworkService struct {
	store storage.Storage
}

func NewNetworkService(store storage.Storage) *NetworkService {
	return &NetworkService{store: store}
}

func (s *NetworkService) CreateNetwork(network models.Network) {
	// network.CreatedAt = time.Now().String()
	// network.UpdatedAt = time.Now().String()

	bytes, err := json.Marshal(network)
	if err != nil {
		fmt.Println(err)
	}
	s.store.StoreBlob(fmt.Sprint("network-", network.ID, ".json"), bytes)
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

func (s *NetworkService) UpdateNetwork(network models.Network) {
	bytes, err := json.Marshal(network)
	if err != nil {
		fmt.Println(err)
	}
	s.store.StoreBlob(fmt.Sprint("network-", network.ID, ".json"), bytes)
}

func (s *NetworkService) DeleteNetwork(id int) {
	s.store.DeleteBlob(fmt.Sprint("network-", id, ".json"))
}
