package parsers

import (
	"github.com/RomeroGabriel/gobrax-challenge/internal/dto"
	"github.com/RomeroGabriel/gobrax-challenge/internal/entity"
)

func CreateTruckDTOToEntity(data dto.CreateTruckDTO) (*entity.Truck, error) {
	return entity.NewTruck(data.ModelType, data.Manufacturer, data.LicensePlate, data.FuelType, data.Year)
}

func EntityToTruckDTO(tdEntity entity.Truck) *dto.TruckResponseDTO {
	return &dto.TruckResponseDTO{
		Id:           tdEntity.ID.String(),
		ModelType:    tdEntity.ModelType,
		Year:         tdEntity.Year,
		Manufacturer: tdEntity.Manufacturer,
		LicensePlate: tdEntity.LicensePlate,
		FuelType:     tdEntity.FuelType,
	}
}
