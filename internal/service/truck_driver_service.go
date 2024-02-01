package service

import (
	"errors"

	"github.com/RomeroGabriel/gobrax-challenge/internal/dto"
	"github.com/RomeroGabriel/gobrax-challenge/internal/dto/parsers"
	"github.com/RomeroGabriel/gobrax-challenge/internal/infra/db"
)

type TruckDriverService struct {
	TruckDriverDB db.TruckDriverRepositoryInterface
}

func NewTruckDriverService(db db.TruckDriverRepositoryInterface) *TruckDriverService {
	return &TruckDriverService{
		TruckDriverDB: db,
	}
}

var (
	ErrDriverNotFound = errors.New("truck driver not found")
)

func (t *TruckDriverService) CreateTruckDriver(input dto.CreateDriverDTO) (*dto.DriverResponseDTO, error) {
	tdEntity, err := parsers.CreateTruckDriverDTOToEntity(input)
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

func (t *TruckDriverService) FindByIdTruckDriver(id string) (*dto.DriverResponseDTO, error) {
	tdEntity, err := t.TruckDriverDB.FindById(id)
	if err != nil {
		return nil, err
	}
	return parsers.EntityToTruckDriverDTO(*tdEntity), err
}

func (t *TruckDriverService) FindByAll() ([]dto.DriverResponseDTO, error) {
	data, err := t.TruckDriverDB.FindAll()
	if err != nil {
		return nil, err
	}
	result := []dto.DriverResponseDTO{}
	for _, v := range data {
		result = append(result, *parsers.EntityToTruckDriverDTO(v))
	}
	return result, err
}

func (t *TruckDriverService) Update(input dto.UpdateDriverDTO) error {
	tdEntity, err := parsers.UpdateTruckDriverDTOToEntity(input)
	if err != nil {
		return err
	}
	return t.TruckDriverDB.Update(tdEntity)
}

func (t *TruckDriverService) Delete(id string) (*dto.DriverResponseDTO, error) {
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
