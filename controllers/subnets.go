package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path"
	"strconv"

	"github.com/aykay76/tempam/models"
	"github.com/aykay76/tempam/services"
	"github.com/aykay76/tempam/storage"
	"github.com/gorilla/mux"
)

type SubnetController struct {
	networkService *services.NetworkService
}

func NewSubnetController(store storage.Storage) *SubnetController {
	networkService := services.NewNetworkService(store)
	return &SubnetController{networkService: networkService}
}

func (controller *SubnetController) SubnetController(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, "SubnetController.SubnetController")

	fmt.Println(r.URL.Path)

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

	switch r.Method {
	case "GET":
		if path.Base(r.URL.Path) == "subnets" {
			controller.getAllTheSubnets(w, r)
		} else {
			controller.getSubnet(w, r)
		}
	case "POST":
		controller.createSubnet(w, r)
	// case "PUT":
	// 	controller.updateSubnet(w, r)
	case "DELETE":
		controller.deleteSubnet(w, r)
	case "OPTIONS":
		controller.returnOptions(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func (controller *SubnetController) getAllTheSubnets(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CommentController.getAllTheSubnets")

	// get the request variables for the network ID
	vars := mux.Vars(r)
	fmt.Println(vars)

	networkId, err := strconv.Atoi(vars["networkId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	subnets := controller.networkService.GetAllSubnets(networkId)
	body, err := json.Marshal(subnets)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	}
}

func (controller *SubnetController) getSubnet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SubnetController.getSubnet")

	// get the request variables for the network ID
	vars := mux.Vars(r)
	fmt.Println(vars)

	networkId, err := strconv.Atoi(vars["networkId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	subnetId, err := strconv.Atoi(vars["subnetId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		subnet := controller.networkService.GetSubnet(networkId, subnetId)
		if subnet == nil {
			w.WriteHeader(http.StatusNotFound)
		} else {
			body, err := json.Marshal(subnet)
			if err != nil {
				fmt.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write(body)
			}
		}
	}
}

func (controller *SubnetController) createSubnet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SubnetController.createSubnet")

	// get the request variables for the network ID
	vars := mux.Vars(r)
	fmt.Println(vars)

	// deserialize the request body into a new subnet
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	var subnetRequest models.SubnetRequest
	json.Unmarshal(requestBody, &subnetRequest)

	// ask the network service to create a subnet on our behalf
	networkId, err := strconv.Atoi(vars["networkId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(requestBody)
	}
	subnet := controller.networkService.CreateSubnet(networkId, subnetRequest)
	if subnet == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	responseBody, err := json.Marshal(subnet)
	if err != nil {
		w.Write(responseBody)
	}

	fmt.Println("Created subnet")
}

// func (controller *SubnetController) updateSubnet(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("SubnetController.updateSubnet")

// 	requestBody, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	var subnet models.Subnet
// 	json.Unmarshal(requestBody, &subnet)

// 	controller.subnetService.UpdateSubnet(subnet)

// 	w.WriteHeader(http.StatusOK)
// }

func (controller *SubnetController) deleteSubnet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SubnetController.deleteSubnet")

	// get the request variables for the network ID
	vars := mux.Vars(r)
	fmt.Println(vars)

	// ask the network service to create a subnet on our behalf
	networkId, err := strconv.Atoi(vars["networkId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	subnetId, err := strconv.Atoi(vars["subnetId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	subnet := controller.networkService.DeleteSubnet(networkId, subnetId)
	if subnet == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (controller *SubnetController) returnOptions(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SubnetController.returnOptions")

	w.WriteHeader(http.StatusOK)
}
