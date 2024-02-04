package service

import (
	"database/sql"
	"errors"

	"github.com/RomeroGabriel/gobrax-challenge/internal/dto"
	"github.com/RomeroGabriel/gobrax-challenge/internal/dto/parsers"
	"github.com/RomeroGabriel/gobrax-challenge/internal/infra/db"
)

type TruckService struct {
	truckDb            *db.TruckRepository
	bindingDriverTruck *db.DriverTruckBindingRespository
}

func NewTruckService(db *db.TruckRepository, bindingDriverTruck *db.DriverTruckBindingRespository) *TruckService {
	return &TruckService{
		truckDb:            db,
		bindingDriverTruck: bindingDriverTruck,
	}
}

var (
	ErrTruckNotFound  = errors.New("truck not found")
	ErrTruckHasDriver = errors.New("truck has linked driver")
)

func (t *TruckService) CreateTruck(input dto.CreateTruckDTO) (*dto.TruckResponseDTO, error) {
	tdEntity, err := parsers.CreateTruckDTOToEntity(input)
	if err != nil {
		return nil, err
	}
	err = t.truckDb.Save(tdEntity)
	if err != nil {
		return nil, err
	}
	return parsers.EntityToTruckDTO(*tdEntity), err
}

func (t *TruckService) FindByIdTruck(id string) (*dto.TruckResponseDTO, error) {
	tdEntity, err := t.truckDb.FindById(id)
	if err != nil {
		return nil, err
	}
	return parsers.EntityToTruckDTO(*tdEntity), err
}

func (t *TruckService) FindByAll() ([]dto.TruckResponseDTO, error) {
	data, err := t.truckDb.FindAll()
	if err != nil {
		return nil, err
	}
	result := []dto.TruckResponseDTO{}
	for _, v := range data {
		result = append(result, *parsers.EntityToTruckDTO(v))
	}
	return result, err
}

func (t *TruckService) Update(input dto.UpdateTruckDTO) error {
	tdEntity, err := t.truckDb.FindById(input.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrTruckNotFound
		}
		return ErrTruckNotFound
	}
	tdEntity.ModelType = input.ModelType
	tdEntity.LicensePlate = input.LicensePlate
	return t.truckDb.Update(tdEntity)
}

func (t *TruckService) Delete(id string) (*dto.TruckResponseDTO, error) {
	tdEntity, err := t.truckDb.FindById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrTruckNotFound
		}
		return nil, err
	}

	isAvailable, err := t.bindingDriverTruck.TruckIsAvailable(*tdEntity)
	if !isAvailable {
		return nil, ErrTruckHasDriver
	}

	err = t.truckDb.Delete(id)
	if err != nil {
		return nil, err
	}
	return parsers.EntityToTruckDTO(*tdEntity), err
}
