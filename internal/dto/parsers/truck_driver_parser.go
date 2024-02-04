package parsers

import (
	"github.com/RomeroGabriel/gobrax-challenge/internal/dto"
	"github.com/RomeroGabriel/gobrax-challenge/internal/entity"
	pkg "github.com/RomeroGabriel/gobrax-challenge/pkg/entity"
)

func CreateTruckDriverDTOToEntity(data dto.CreateDriverDTO) (*entity.Driver, error) {
	return entity.NewTruckDriver(data.Name, data.Email, data.LicenseNumber)
}

func UpdateTruckDriverDTOToEntity(data dto.UpdateDriverDTO) (*entity.Driver, error) {
	id, err := pkg.ParseID(data.Id)
	if err != nil {
		return nil, err
	}
	return &entity.Driver{
		ID:            id,
		Name:          data.Name,
		Email:         data.Email,
		LicenseNumber: data.LicenseNumber,
	}, nil
}

func EntityToTruckDriverDTO(tdEntity entity.Driver) *dto.DriverResponseDTO {
	return &dto.DriverResponseDTO{
		Id:    tdEntity.ID.String(),
		Name:  tdEntity.Name,
		Email: tdEntity.Email,
	}
}
