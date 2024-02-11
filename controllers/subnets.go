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
	subnetService  *services.SubnetService
}

func NewSubnetController(store storage.Storage) *SubnetController {
	networkService := services.NewNetworkService(store)
	subnetService := services.NewSubnetService(store)
	return &SubnetController{networkService: networkService, subnetService: subnetService}
}

func (this *SubnetController) SubnetController(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, "SubnetController.SubnetController")

	fmt.Println(r.URL.Path)

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

	switch r.Method {
	case "GET":
		if path.Base(r.URL.Path) == "subnets" {
			this.getAllTheSubnets(w, r)
		} else {
			this.getSubnet(w, r)
		}
	case "POST":
		this.createSubnet(w, r)
	case "PUT":
		this.updateSubnet(w, r)
	case "DELETE":
		this.deleteSubnet(w, r)
	case "OPTIONS":
		this.returnOptions(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func (this *SubnetController) getAllTheSubnets(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CommentController.getAllTheSubnets")

	body, err := json.Marshal(this.subnetService.GetAllSubnets())
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	}
}

func (this *SubnetController) getSubnet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SubnetController.getSubnet")

	num, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		subnet := this.subnetService.GetSubnet(num)
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

func (this *SubnetController) createSubnet(w http.ResponseWriter, r *http.Request) {
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
	subnet := this.networkService.CreateSubnet(networkId, subnetRequest)
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

func (controller *SubnetController) updateSubnet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SubnetController.updateSubnet")

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	var subnet models.Subnet
	json.Unmarshal(requestBody, &subnet)

	controller.subnetService.UpdateSubnet(subnet)

	w.WriteHeader(http.StatusOK)
}

func (this *SubnetController) deleteSubnet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SubnetController.deleteSubnet")

	id, _ := strconv.Atoi(path.Base(r.URL.Path))

	this.subnetService.DeleteSubnet(id)

	w.WriteHeader(http.StatusOK)
}

func (this *SubnetController) returnOptions(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SubnetController.returnOptions")

	w.WriteHeader(http.StatusOK)
}
