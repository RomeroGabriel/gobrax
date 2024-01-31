package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var modelType = "Model1"
var year uint16 = 1900
var manufacturer = "Volvo"
var licensePlate = "12354"
var fuelType = "Diesel"

func TestNewTruck(t *testing.T) {
	// modelType, manufacturer, licensePlate, fuelType string, year uint16
	td, err := NewTruck(modelType, manufacturer, licensePlate, fuelType, year)
	assert.Nil(t, err)
	assert.NotNil(t, td)
	assert.Equal(t, modelType, td.ModelType)
	assert.Equal(t, year, td.Year)
	assert.Equal(t, manufacturer, td.Manufacturer)
	assert.Equal(t, licensePlate, td.LicensePlate)
	assert.Equal(t, fuelType, td.FuelType)
}

func TestTruckWhenModelIsRequired(t *testing.T) {
	td, err := NewTruck("", manufacturer, licensePlate, fuelType, year)
	assert.Nil(t, td)
	assert.Equal(t, ErrModelTypeIsRequired, err)
}

func TestTruckWhenManufacturerIsRequired(t *testing.T) {
	td, err := NewTruck(modelType, "", licensePlate, fuelType, year)
	assert.Nil(t, td)
	assert.Equal(t, ErrManufacturerIsRequired, err)
}

func TestTruckWhenLicensePlateIsRequired(t *testing.T) {
	td, err := NewTruck(modelType, manufacturer, "", fuelType, year)
	assert.Nil(t, td)
	assert.Equal(t, ErrLicensePlateIsRequired, err)
}

func TestTruckWhenFuelTypeIsRequired(t *testing.T) {
	td, err := NewTruck(modelType, manufacturer, licensePlate, "", year)
	assert.Nil(t, td)
	assert.Equal(t, ErrFuelTypeIsRequired, err)
}

func TestTruckWhenYearIsRequired(t *testing.T) {
	td, err := NewTruck(modelType, manufacturer, licensePlate, fuelType, 0)
	assert.Nil(t, td)
	assert.Equal(t, ErrYearIsRequired, err)
}
