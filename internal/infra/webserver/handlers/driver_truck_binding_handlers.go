package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/RomeroGabriel/gobrax-challenge/internal/dto"
	"github.com/RomeroGabriel/gobrax-challenge/internal/service"
)

type DriverTruckBindingHandler struct {
	bindingService service.IDriverTruckBindingService
}

func NewDriverTruckBindingHandler(service service.IDriverTruckBindingService) *DriverTruckBindingHandler {
	return &DriverTruckBindingHandler{
		bindingService: service,
	}
}

func (h *DriverTruckBindingHandler) CreateBinding(w http.ResponseWriter, r *http.Request) {
	var data dto.CreateBindingDTO
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	responseData, err := h.bindingService.BindingDriverToTruck(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(responseData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
