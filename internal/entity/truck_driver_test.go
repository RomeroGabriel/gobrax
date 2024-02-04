package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var name = "Trucker 1"
var email = "email1@gg.com"
var licenseNumber = "ABC12345"

func TestNewTruckDriver(t *testing.T) {
	td, err := NewDriver(name, email, licenseNumber)
	assert.Nil(t, err)
	assert.NotNil(t, td)
	assert.Equal(t, name, td.Name)
	assert.Equal(t, email, td.Email)
	assert.Equal(t, licenseNumber, td.LicenseNumber)
}

func TestTruckDriverWhenNameIsRequired(t *testing.T) {
	p, err := NewDriver("", email, licenseNumber)
	assert.Nil(t, p)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestTruckDriverWhenEmailIsRequired(t *testing.T) {
	p, err := NewDriver(name, "", licenseNumber)
	assert.Nil(t, p)
	assert.Equal(t, ErrEmailIsRequired, err)
}

func TestTruckDriverWhenLicenseIsRequired(t *testing.T) {
	p, err := NewDriver(name, email, "")
	assert.Nil(t, p)
	assert.Equal(t, ErrLicenseNumberIsRequired, err)
}
