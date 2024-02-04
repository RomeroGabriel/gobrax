package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/RomeroGabriel/gobrax-challenge/internal/dto"
	"github.com/RomeroGabriel/gobrax-challenge/internal/service"
	"github.com/go-chi/chi"
)

type WebTruckDriverHandler struct {
	TruckDriverService service.ITruckDriverService
}

func NewWebTruckDriverHandler(service service.ITruckDriverService) *WebTruckDriverHandler {
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

func (h *WebTruckDriverHandler) FindById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	responseData, err := h.TruckDriverService.FindByIdTruckDriver(id)
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

func (h *WebTruckDriverHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	responseData, err := h.TruckDriverService.FindByAll()
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

func (h *WebTruckDriverHandler) Update(w http.ResponseWriter, r *http.Request) {
	var data dto.UpdateDriverDTO
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	data.Id = id
	err = h.TruckDriverService.Update(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *WebTruckDriverHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	data, err := h.TruckDriverService.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(data)
}
