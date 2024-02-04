package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/RomeroGabriel/gobrax-challenge/internal/dto"
	"github.com/RomeroGabriel/gobrax-challenge/internal/service"
	"github.com/go-chi/chi"
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
	idTruck := chi.URLParam(r, "idtruck")
	if idTruck == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	idDriver := chi.URLParam(r, "iddriver")
	if idDriver == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var data dto.CreateBindingDTO
	data.IdDriver = idDriver
	data.IdTruck = idTruck

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
