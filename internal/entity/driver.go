package entity

import (
	"errors"

	"github.com/RomeroGabriel/gobrax-challenge/pkg/entity"
)

type Driver struct {
	ID            entity.ID
	Name          string
	Email         string
	LicenseNumber string
}

var (
	ErrNameIsRequired          = errors.New("driver name is required")
	ErrEmailIsRequired         = errors.New("driver email is required")
	ErrLicenseNumberIsRequired = errors.New("driver license number is required")
)

func NewDriver(name, email, licenseNumber string) (*Driver, error) {
	id := entity.NewID()
	td := &Driver{
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

func (t Driver) Validate() error {
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
