package parsers

import (
	"github.com/RomeroGabriel/gobrax-challenge/internal/dto"
	"github.com/RomeroGabriel/gobrax-challenge/internal/entity"
)

func TruckDriverDTOToEntity(data dto.CreateDriverDTO) (*entity.TruckDriver, error) {
	return entity.NewTruckDriver(data.Name, data.Email, data.LicenseNumber)
}
