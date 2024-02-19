package services

import (
	"encoding/json"
	"fmt"
	"slices"
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

func (service *NetworkService) CreateNetwork(networkRequest models.NetworkRequest) *models.Network {
	// TODO: check for any overlaps and reject

	// find the last network
	names := service.ListNetworks()
	lastId := len(names)

	network := models.Network{
		ID:      lastId + 1,
		Name:    networkRequest.Name,
		CIDR:    networkRequest.CIDR,
		Subnets: make([]*models.Subnet, 0),
	}

	network.AddressSpace.FromCIDR(networkRequest.CIDR)
	network.CreatedAt = time.Now()
	network.UpdatedAt = time.Now()

	fmt.Println(network)
	// bytes, err := json.Marshal(network)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	err := service.store.StoreBlob("networks", fmt.Sprint(network.ID, ".json"), network)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &network
}

func (service *NetworkService) CreateSubnet(networkId int, subnetRequest models.SubnetRequest) *models.Subnet {
	fmt.Printf("Creating subnet in network %d, with %d\n", networkId, subnetRequest.MaskLength)

	network := service.GetNetwork(networkId)
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
	// TODO: add CIDR to output for easy reading

	network.Subnets = append(network.Subnets, subnet)
	service.UpdateNetwork(network)

	return subnet
}

func (service *NetworkService) ListNetworks() []string {
	names, err := service.store.ListBlobs("networks", "network-*.json")
	if err != nil {
		fmt.Println(err)
	}
	return names
}

func (service *NetworkService) GetAllNetworks() []models.Network {
	var filter models.Network
	var networks []models.Network

	err := service.store.GetAllBlobs("networks", filter, &networks)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return networks
}

func (service *NetworkService) GetAllSubnets(networkId int) []*models.Subnet {
	network := service.GetNetwork(networkId)
	if network == nil {
		return nil
	}

	return network.Subnets
}

func (service *NetworkService) GetSubnet(networkId int, subnetId int) *models.Subnet {
	network := service.GetNetwork(networkId)
	if network == nil {
		return nil
	}

	for _, subnet := range network.Subnets {
		if subnet.ID == subnetId {
			return subnet
		}
	}

	return nil
}

func (service *NetworkService) DeleteSubnet(networkId int, subnetId int) error {
	network := service.GetNetwork(networkId)

	todelete := -1

	for k, subnet := range network.Subnets {
		if subnet.ID == subnetId {
			todelete = k
			break
		}
	}

	if todelete != -1 {
		network.Subnets = slices.Delete(network.Subnets, todelete, todelete)
	}

	service.UpdateNetwork(network)
	return nil
}

func (service *NetworkService) GetNetwork(id int) *models.Network {
	fmt.Println("NetworkService.GetNetwork ", id)
	var network models.Network
	blob, err := service.store.GetBlob("networks", fmt.Sprint(id, ".json"))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	json.Unmarshal(blob, &network)
	return &network
}

func (service *NetworkService) UpdateNetwork(network *models.Network) {
	bytes, err := json.Marshal(network)
	if err != nil {
		fmt.Println(err)
	}
	service.store.StoreBlob("networks", fmt.Sprint(network.ID, ".json"), bytes)
}

func (service *NetworkService) DeleteNetwork(id int) {
	service.store.DeleteBlob("networks", fmt.Sprint(id, ".json"))
}
