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

type NetworkController struct {
	networkService *services.NetworkService
}

func NewNetworkController(store storage.Storage) *NetworkController {
	networkService := services.NewNetworkService(store)
	return &NetworkController{networkService: networkService}
}

func (this *NetworkController) NetworkController(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, "NetworkController.NetworkController")

	vars := mux.Vars(r)
	fmt.Println(vars)

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

	switch r.Method {
	case "GET":
		fmt.Println(r.URL.Path)
		if path.Base(r.URL.Path) == "networks" {
			this.getAllTheNetworks(w, r)
		} else {
			this.getNetwork(w, r)
		}
	case "POST":
		this.createNetwork(w, r)
	case "PUT":
		this.updateNetwork(w, r)
	case "DELETE":
		this.deleteNetwork(w, r)
	case "OPTIONS":
		this.returnOptions(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func (this *NetworkController) getAllTheNetworks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CommentController.getAllTheNetworks")

	body, err := json.Marshal(this.networkService.GetAllNetworks())
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	}
}

func (this *NetworkController) getNetwork(w http.ResponseWriter, r *http.Request) {
	fmt.Println("NetworkController.getNetwork")

	num, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		network := this.networkService.GetNetwork(num)
		if network == nil {
			w.WriteHeader(http.StatusNotFound)
		} else {
			body, err := json.Marshal(network)
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

func (this *NetworkController) createNetwork(w http.ResponseWriter, r *http.Request) {
	fmt.Println("NetworkController.createNetwork")

	// deserialize the request body into a new network
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	var networkRequest models.NetworkRequest
	json.Unmarshal(requestBody, &networkRequest)

	// ask the network service to create the network
	network := this.networkService.CreateNetwork(networkRequest)
	if network == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// if we got this far we can report success
	w.WriteHeader(http.StatusCreated)
	w.Write(requestBody)
	fmt.Println("Created network")
}

func (controller *NetworkController) updateNetwork(w http.ResponseWriter, r *http.Request) {
	fmt.Println("NetworkController.updateNetwork")

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	var network models.Network
	json.Unmarshal(requestBody, &network)

	controller.networkService.UpdateNetwork(&network)

	w.WriteHeader(http.StatusOK)
}

func (this *NetworkController) deleteNetwork(w http.ResponseWriter, r *http.Request) {
	fmt.Println("NetworkController.deleteNetwork")

	id, _ := strconv.Atoi(path.Base(r.URL.Path))

	this.networkService.DeleteNetwork(id)

	w.WriteHeader(http.StatusOK)
}

func (this *NetworkController) returnOptions(w http.ResponseWriter, r *http.Request) {
	fmt.Println("NetworkController.returnOptions")

	w.WriteHeader(http.StatusOK)
}
