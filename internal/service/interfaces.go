package service

import (
	"github.com/RomeroGabriel/gobrax-challenge/internal/dto"
)

type TruckDriverServiceInterface interface {
	CreateTruckDriver(input dto.CreateDriverDTO) (*dto.DriverResponseDTO, error)
	FindByIdTruckDriver(id string) (*dto.DriverResponseDTO, error)
	FindByAll() ([]dto.DriverResponseDTO, error)
	Update(input dto.UpdateDriverDTO) error
	Delete(id string) (*dto.DriverResponseDTO, error)
}
