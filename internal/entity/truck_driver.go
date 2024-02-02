package entity

import (
	"errors"

	"github.com/RomeroGabriel/gobrax-challenge/pkg/entity"
)

type TruckDriver struct {
	ID            entity.ID
	Name          string
	Email         string
	LicenseNumber string
}

var (
	ErrNameIsRequired          = errors.New("truck driver name is required")
	ErrEmailIsRequired         = errors.New("truck driver email is required")
	ErrLicenseNumberIsRequired = errors.New("truck driver license number is required")
)

func NewTruckDriver(name, email, licenseNumber string) (*TruckDriver, error) {
	// TODO: Add validations for email(eee@@@@.com), name(size), and licenseNumber(size)
	// TODO: Create entity for LicenseNumber
	id := entity.NewID()
	td := &TruckDriver{
		ID:            id,
		Name:          name,
		Email:         email,
		LicenseNumber: licenseNumber,
	}
	err := td.Validate()
	if err != nil {
		return nil, err
	}
	return td, nil
}

func (t TruckDriver) Validate() error {
	if t.ID.String() == "" {
		return entity.ErrIDIsRequired
	}
	if t.Name == "" {
		return ErrNameIsRequired
	}
	if t.Email == "" {
		return ErrEmailIsRequired
	}
	if t.LicenseNumber == "" {
		return ErrLicenseNumberIsRequired
	}
	return nil
}
