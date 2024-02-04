package parsers

import (
	"github.com/RomeroGabriel/gobrax-challenge/internal/dto"
	"github.com/RomeroGabriel/gobrax-challenge/internal/entity"
)

func CreateDriverDTOToEntity(data dto.CreateDriverDTO) (*entity.Driver, error) {
	return entity.NewDriver(data.Name, data.Email, data.LicenseNumber)
}

func EntityToDriverDTO(tdEntity entity.Driver) *dto.DriverResponseDTO {
	return &dto.DriverResponseDTO{
		Id:    tdEntity.ID.String(),
		Name:  tdEntity.Name,
		Email: tdEntity.Email,
	}
}
