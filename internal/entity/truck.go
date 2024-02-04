package entity

import (
	"errors"

	"github.com/RomeroGabriel/gobrax-challenge/pkg/entity"
)

type Truck struct {
	ID           entity.ID
	ModelType    string
	Year         uint16
	Manufacturer string
	LicensePlate string
	FuelType     string
}

var (
	ErrModelTypeIsRequired    = errors.New("model type is required")
	ErrYearIsRequired         = errors.New("year is required")
	ErrManufacturerIsRequired = errors.New("manufacturer is required")
	ErrLicensePlateIsRequired = errors.New("license plate is required")
	ErrFuelTypeIsRequired     = errors.New("fuel type is required")
)

func NewTruck(modelType, manufacturer, licensePlate, fuelType string, year uint16) (*Truck, error) {
	id := entity.NewID()
	td := &Truck{
		ID:           id,
		ModelType:    modelType,
		Year:         year,
		Manufacturer: manufacturer,
		LicensePlate: licensePlate,
		FuelType:     fuelType,
	}
	err := td.Validate()
	if err != nil {
		return nil, err
	}
	return td, nil
}

func (t Truck) Validate() error {
	if t.ID.String() == "" {
		return entity.ErrIDIsRequired
	}
	if t.ModelType == "" {
		return ErrModelTypeIsRequired
	}
	if t.Year == 0 {
		return ErrYearIsRequired
	}
	if t.Manufacturer == "" {
		return ErrManufacturerIsRequired
	}
	if t.LicensePlate == "" {
		return ErrLicensePlateIsRequired
	}
	if t.FuelType == "" {
		return ErrFuelTypeIsRequired
	}
	return nil
}
