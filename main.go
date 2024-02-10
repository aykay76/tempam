package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"sync"

	"github.com/aykay76/tempam/controllers"
	"github.com/aykay76/tempam/models"
	"github.com/aykay76/tempam/storage"
	"github.com/gorilla/mux"
	"go.uber.org/dig"
)

func ip2int(ip string) uint32 {
	var long uint32
	binary.Read(bytes.NewBuffer(net.ParseIP(ip).To4()), binary.BigEndian, &long)
	return long
}

func int2ip(ipInt uint64) string {
	// need to do two bit shifting and “0xff” masking
	b0 := strconv.FormatUint((ipInt>>24)&0xff, 10)
	b1 := strconv.FormatUint((ipInt>>16)&0xff, 10)
	b2 := strconv.FormatUint((ipInt>>8)&0xff, 10)
	b3 := strconv.FormatUint((ipInt & 0xff), 10)
	return b0 + "." + b1 + "." + b2 + "." + b3
}

// IPAM is a simple IP Address Management structure.
type IPAM struct {
	Networks []models.Network

	mutex sync.Mutex
}

func main() {
	// add some things to dependency injection container
	container := dig.New()
	container.Provide(storage.LocalStorage)
	container.Provide(controllers.NewNetworkController)

	// create a new router
	router := mux.NewRouter()

	// add routes for networks - top level object that represents a virtual network/VPC
	container.Invoke(func(controller *controllers.NetworkController) {
		router.HandleFunc("/api/networks", controller.NetworkController).Methods("GET")
		router.HandleFunc("/api/networks/{id}", controller.NetworkController).Methods("GET")
		router.HandleFunc("/api/networks", controller.NetworkController).Methods("POST")
		router.HandleFunc("/api/networks/{id}", controller.NetworkController).Methods("DELETE")
	})

	// TODO: add routes for subnets

	// TODO: add routes for individual devices?

	port := 8080
	fmt.Printf("Starting IPAM API on :%d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}
