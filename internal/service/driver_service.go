package service

import (
	"database/sql"
	"errors"

	"github.com/RomeroGabriel/gobrax-challenge/internal/dto"
	"github.com/RomeroGabriel/gobrax-challenge/internal/dto/parsers"
	"github.com/RomeroGabriel/gobrax-challenge/internal/infra/db"
)

type DriverService struct {
	driverDB           *db.DriverRepository
	bindingDriverTruck *db.DriverTruckBindingRespository
}

func NewDriverService(db *db.DriverRepository, bindingDriverTruck *db.DriverTruckBindingRespository) *DriverService {
	return &DriverService{
		driverDB:           db,
		bindingDriverTruck: bindingDriverTruck,
	}
}

var (
	ErrDriverNotFound = errors.New("driver not found")
	ErrDriverHasTruck = errors.New("driver has linked truck")
)

func (t *DriverService) CreateDriver(input dto.CreateDriverDTO) (*dto.DriverResponseDTO, error) {
	tdEntity, err := parsers.CreateDriverDTOToEntity(input)
	if err != nil {
		return nil, err
	}
	err = t.driverDB.Save(tdEntity)
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
	tdEntity, err := t.driverDB.FindById(id)
	if err != nil {
		return nil, err
	}
	return parsers.EntityToDriverDTO(*tdEntity), err
}

func (t *DriverService) FindAll() ([]dto.DriverResponseDTO, error) {
	data, err := t.driverDB.FindAll()
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
	data, err := t.driverDB.FindById(input.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrDriverNotFound
		}
		return err
	}
	data.Name = input.Name
	data.Email = input.Email
	data.LicenseNumber = input.LicenseNumber
	return t.driverDB.Update(data)
}

func (t *DriverService) Delete(id string) (*dto.DriverResponseDTO, error) {
	td, err := t.driverDB.FindById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrDriverNotFound
		}
		return nil, ErrDriverNotFound
	}

	isAvailable, err := t.bindingDriverTruck.DriverIsAvailable(*td)
	if !isAvailable {
		return nil, ErrDriverHasTruck
	}

	err = t.driverDB.Delete(id)
	if err != nil {
		return nil, err
	}
	return &dto.DriverResponseDTO{
		Id:    td.ID.String(),
		Name:  td.Name,
		Email: td.Email,
	}, err
}
