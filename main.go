package main

import (
	"fmt"
	"net/http"

	"github.com/aykay76/tempam/controllers"
	"github.com/aykay76/tempam/storage"
	"github.com/gorilla/mux"
	"go.uber.org/dig"
)

func main() {
	// add some things to dependency injection container
	container := dig.New()
	container.Provide(storage.LocalStorage)
	container.Provide(controllers.NewNetworkController)
	container.Provide(controllers.NewSubnetController)

	// create a new router
	router := mux.NewRouter()

	// add routes for networks - top level object that represents a virtual network/VPC
	container.Invoke(func(controller *controllers.NetworkController) {
		router.HandleFunc("/api/networks", controller.NetworkController).Methods("GET", "POST")
		router.HandleFunc("/api/networks/{networkId}", controller.NetworkController).Methods("GET", "DELETE")
	})

	container.Invoke(func(controller *controllers.SubnetController) {
		router.HandleFunc("/api/networks/{networkId}/subnets", controller.SubnetController).Methods("POST")
	})

	port := 8080
	fmt.Printf("Starting IPAM API on :%d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}
