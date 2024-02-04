package service

import (
	"errors"

	"github.com/RomeroGabriel/gobrax-challenge/internal/dto"
	"github.com/RomeroGabriel/gobrax-challenge/internal/dto/parsers"
	"github.com/RomeroGabriel/gobrax-challenge/internal/infra/db"
)

type DriverService struct {
	TruckDriverDB db.IDriverRepository
}

func NewTruckDriverService(db db.IDriverRepository) *DriverService {
	return &DriverService{
		TruckDriverDB: db,
	}
}

var (
	ErrDriverNotFound = errors.New("driver not found")
)

func (t *DriverService) CreateDriver(input dto.CreateDriverDTO) (*dto.DriverResponseDTO, error) {
	tdEntity, err := parsers.CreateDriverDTOToEntity(input)
	if err != nil {
		return nil, err
	}
	err = t.TruckDriverDB.Save(tdEntity)
	if err != nil {
		return nil, err
	}
	return &dto.DriverResponseDTO{
		Id:    tdEntity.ID.String(),
		Name:  tdEntity.Name,
		Email: tdEntity.Email,
	}, err
}

func (t *DriverService) FindById(id string) (*dto.DriverResponseDTO, error) {
	tdEntity, err := t.TruckDriverDB.FindById(id)
	if err != nil {
		return nil, err
	}
	return parsers.EntityToDriverDTO(*tdEntity), err
}

func (t *DriverService) FindAll() ([]dto.DriverResponseDTO, error) {
	data, err := t.TruckDriverDB.FindAll()
	if err != nil {
		return nil, err
	}
	result := []dto.DriverResponseDTO{}
	for _, v := range data {
		result = append(result, *parsers.EntityToDriverDTO(v))
	}
	return result, err
}

func (t *DriverService) Update(input dto.UpdateDriverDTO) error {
	tdEntity, err := parsers.UpdateDriverDTOToEntity(input)
	if err != nil {
		return err
	}
	return t.TruckDriverDB.Update(tdEntity)
}

func (t *DriverService) Delete(id string) (*dto.DriverResponseDTO, error) {
	td, err := t.TruckDriverDB.FindById(id)
	if err != nil {
		return nil, ErrDriverNotFound
	}
	err = t.TruckDriverDB.Delete(id)
	if err != nil {
		return nil, err
	}
	return &dto.DriverResponseDTO{
		Id:    td.ID.String(),
		Name:  td.Name,
		Email: td.Email,
	}, err
}
