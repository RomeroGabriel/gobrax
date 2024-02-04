package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/RomeroGabriel/gobrax-challenge/internal/dto"
	"github.com/RomeroGabriel/gobrax-challenge/internal/service"
	"github.com/go-chi/chi"
)

type WebDriverHandler struct {
	DriverService service.IDriverService
}

func NewWebDriverHandler(service service.IDriverService) *WebDriverHandler {
	return &WebDriverHandler{
		DriverService: service,
	}
}

func (h *WebDriverHandler) Create(w http.ResponseWriter, r *http.Request) {
	var data dto.CreateDriverDTO
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	responseData, err := h.DriverService.CreateDriver(data)
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

func (h *WebDriverHandler) FindById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	responseData, err := h.DriverService.FindById(id)
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

func (h *WebDriverHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	responseData, err := h.DriverService.FindAll()
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

func (h *WebDriverHandler) Update(w http.ResponseWriter, r *http.Request) {
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
	err = h.DriverService.Update(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *WebDriverHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	data, err := h.DriverService.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(data)
}
