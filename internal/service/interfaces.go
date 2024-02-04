package service

import (
	"github.com/RomeroGabriel/gobrax-challenge/internal/dto"
	pkg "github.com/RomeroGabriel/gobrax-challenge/pkg/entity"
)

type IDriverService interface {
	CreateTruckDriver(input dto.CreateDriverDTO) (*dto.DriverResponseDTO, error)
	FindByIdTruckDriver(id string) (*dto.DriverResponseDTO, error)
	FindByAll() ([]dto.DriverResponseDTO, error)
	Update(input dto.UpdateDriverDTO) error
	Delete(id string) (*dto.DriverResponseDTO, error)
}

type ITruckService interface {
	CreateTruck(input dto.CreateTruckDTO) (*dto.TruckResponseDTO, error)
	FindByIdTruck(id string) (*dto.TruckResponseDTO, error)
	FindByAll() ([]dto.TruckResponseDTO, error)
	Update(input dto.UpdateTruckDTO) error
	Delete(id string) (*dto.TruckResponseDTO, error)
}

type IDriverTruckBindingService interface {
	BindingDriverToTruck(input dto.CreateBindingDTO) (*pkg.ID, error)
}
