package parsers

import (
	"github.com/RomeroGabriel/gobrax-challenge/internal/dto"
	"github.com/RomeroGabriel/gobrax-challenge/internal/entity"
)

func TruckDriverDTOToEntity(data dto.CreateDriverDTO) (*entity.TruckDriver, error) {
	return entity.NewTruckDriver(data.Name, data.Email, data.LicenseNumber)
}

func EntityToTuckDriverDTO(tdEntity entity.TruckDriver) *dto.DriverResponseDTO {
	return &dto.DriverResponseDTO{
		Id:    tdEntity.ID.String(),
		Name:  tdEntity.Name,
		Email: tdEntity.Email,
	}
}
