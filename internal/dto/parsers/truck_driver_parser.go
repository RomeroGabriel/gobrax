package parsers

import (
	"github.com/RomeroGabriel/gobrax-challenge/internal/dto"
	"github.com/RomeroGabriel/gobrax-challenge/internal/entity"
	pkg "github.com/RomeroGabriel/gobrax-challenge/pkg/entity"
)

func CreateDriverDTOToEntity(data dto.CreateDriverDTO) (*entity.Driver, error) {
	return entity.NewDriver(data.Name, data.Email, data.LicenseNumber)
}

func UpdateDriverDTOToEntity(data dto.UpdateDriverDTO) (*entity.Driver, error) {
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

func EntityToDriverDTO(tdEntity entity.Driver) *dto.DriverResponseDTO {
	return &dto.DriverResponseDTO{
		Id:    tdEntity.ID.String(),
		Name:  tdEntity.Name,
		Email: tdEntity.Email,
	}
}
