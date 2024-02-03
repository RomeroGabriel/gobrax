package service

import (
	"errors"

	"github.com/RomeroGabriel/gobrax-challenge/internal/dto"
	"github.com/RomeroGabriel/gobrax-challenge/internal/dto/parsers"
	"github.com/RomeroGabriel/gobrax-challenge/internal/infra/db"
)

type TruckService struct {
	TruckDB db.ITruckRepository
}

func NewTruckService(db db.ITruckRepository) *TruckService {
	return &TruckService{
		TruckDB: db,
	}
}

var (
	ErrTruckNotFound = errors.New("truck not found")
)

func (t *TruckService) CreateTruck(input dto.CreateTruckDTO) (*dto.TruckResponseDTO, error) {
	tdEntity, err := parsers.CreateTruckDTOToEntity(input)
	if err != nil {
		return nil, err
	}
	err = t.TruckDB.Save(tdEntity)
	if err != nil {
		return nil, err
	}
	return parsers.EntityToTruckDTO(*tdEntity), err
}

func (t *TruckService) FindByIdTruck(id string) (*dto.TruckResponseDTO, error) {
	tdEntity, err := t.TruckDB.FindById(id)
	if err != nil {
		return nil, err
	}
	return parsers.EntityToTruckDTO(*tdEntity), err
}

func (t *TruckService) FindByAll() ([]dto.TruckResponseDTO, error) {
	data, err := t.TruckDB.FindAll()
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
	tdEntity, err := t.TruckDB.FindById(input.Id)
	if err != nil {
		return ErrTruckNotFound
	}
	tdEntity.ModelType = input.ModelType
	tdEntity.LicensePlate = input.LicensePlate
	return t.TruckDB.Update(tdEntity)
}

func (t *TruckService) Delete(id string) (*dto.TruckResponseDTO, error) {
	tdEntity, err := t.TruckDB.FindById(id)
	if err != nil {
		return nil, ErrTruckNotFound
	}
	err = t.TruckDB.Delete(id)
	if err != nil {
		return nil, err
	}
	return parsers.EntityToTruckDTO(*tdEntity), err
}
