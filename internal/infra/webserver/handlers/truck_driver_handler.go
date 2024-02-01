package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/RomeroGabriel/gobrax-challenge/internal/dto"
	"github.com/RomeroGabriel/gobrax-challenge/internal/service"
)

type WebTruckDriverHandler struct {
	TruckDriverService service.TruckDriverServiceInterface
}

func NewWebTruckDriverHandler(service service.TruckDriverServiceInterface) *WebTruckDriverHandler {
	return &WebTruckDriverHandler{
		TruckDriverService: service,
	}
}

func (h *WebTruckDriverHandler) Create(w http.ResponseWriter, r *http.Request) {
	var data dto.CreateDriverDTO
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	responseData, err := h.TruckDriverService.CreateTruckDriver(data)
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
