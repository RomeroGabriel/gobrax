package service

import (
	"github.com/RomeroGabriel/gobrax-challenge/internal/dto"
)

type TruckDriverServiceInterface interface {
	CreateTruckDriver(input dto.CreateDriverDTO) (*dto.DriverResponseDTO, error)
}
