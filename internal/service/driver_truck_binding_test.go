//go:build !prod
// +build !prod

package service

import (
	"database/sql"
	"testing"

	"github.com/RomeroGabriel/gobrax-challenge/internal/dto"
	"github.com/RomeroGabriel/gobrax-challenge/internal/entity"
	"github.com/RomeroGabriel/gobrax-challenge/internal/infra/db"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type DriverTruckBindingServiceTestSuite struct {
	suite.Suite
	dbInstance *sql.DB
	// TruckDriverDB        db.IDriverRepository
	// TruckDB              db.ITruckRepository
	// DriverTruckBindingDB db.IDriverTruckBindingRespository
}

func (suite *DriverTruckBindingServiceTestSuite) SetupSuite() {
	dbInstance, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	suite.dbInstance = dbInstance
}

func TestSuiteTruckDriver(t *testing.T) {
	suite.Run(t, new(DriverTruckBindingServiceTestSuite))
}

var modelType = "Model1"
var year uint16 = 1900
var manufacturer = "Volvo"
var licensePlate = "12354"
var fuelType = "Diesel"
var name = "Trucke Driver 1"
var email = "email1@gg.com"
var license = "ABC12345"

func (suite *DriverTruckBindingServiceTestSuite) TestBindingDriverToTruck() {
	truckDriverDB := db.NewTruckDriverRepository(suite.dbInstance)
	truckDB := db.NewTruckRepository(suite.dbInstance)
	driverTruckBindingDB := db.NewDriverTruckBindingRespository(suite.dbInstance)

	td, err := entity.NewDriver(name, email, license)
	suite.NoError(err)
	err = truckDriverDB.Save(td)
	suite.NoError(err)

	t, err := entity.NewTruck(modelType, manufacturer, licensePlate, fuelType, year)
	err = truckDB.Save(t)
	suite.NoError(err)

	service := NewDriverTruckBindingService(truckDriverDB, truckDB, driverTruckBindingDB)
	var input = dto.CreateBindingDTO{
		IdDriver: td.ID.String(),
		IdTruck:  t.ID.String(),
	}
	id, err := service.BindingDriverToTruck(input)
	suite.NoError(err)
	suite.NotNil(id)
}

func (suite *DriverTruckBindingServiceTestSuite) TestBindingTruckAndDriverToTruckNotAvailable() {
	truckDriverDB := db.NewTruckDriverRepository(suite.dbInstance)
	truckDB := db.NewTruckRepository(suite.dbInstance)
	driverTruckBindingDB := db.NewDriverTruckBindingRespository(suite.dbInstance)

	td, err := entity.NewDriver(name, email, license)
	suite.NoError(err)
	err = truckDriverDB.Save(td)
	suite.NoError(err)

	t, err := entity.NewTruck(modelType, manufacturer, licensePlate, fuelType, year)
	err = truckDB.Save(t)
	suite.NoError(err)

	service := NewDriverTruckBindingService(truckDriverDB, truckDB, driverTruckBindingDB)
	var input = dto.CreateBindingDTO{
		IdDriver: td.ID.String(),
		IdTruck:  t.ID.String(),
	}
	_, err = service.BindingDriverToTruck(input)
	suite.NoError(err)
	_, err = service.BindingDriverToTruck(input)
	suite.Error(err)
	suite.Equal(err, ErrTruckAndDriverIsNotAvailable)
}

func (suite *DriverTruckBindingServiceTestSuite) TestBindingDriverToTruckNotAvailable() {
	truckDriverDB := db.NewTruckDriverRepository(suite.dbInstance)
	truckDB := db.NewTruckRepository(suite.dbInstance)
	driverTruckBindingDB := db.NewDriverTruckBindingRespository(suite.dbInstance)

	td, err := entity.NewDriver(name, email, license)
	suite.NoError(err)
	err = truckDriverDB.Save(td)
	suite.NoError(err)

	t, err := entity.NewTruck(modelType, manufacturer, licensePlate, fuelType, year)
	err = truckDB.Save(t)
	suite.NoError(err)

	service := NewDriverTruckBindingService(truckDriverDB, truckDB, driverTruckBindingDB)
	var input = dto.CreateBindingDTO{
		IdDriver: td.ID.String(),
		IdTruck:  t.ID.String(),
	}
	_, err = service.BindingDriverToTruck(input)
	suite.NoError(err)

	t, err = entity.NewTruck(modelType, manufacturer, licensePlate, fuelType, year)
	err = truckDB.Save(t)
	suite.NoError(err)
	input = dto.CreateBindingDTO{
		IdDriver: td.ID.String(),
		IdTruck:  t.ID.String(),
	}

	_, err = service.BindingDriverToTruck(input)
	suite.Error(err)
	suite.Equal(err, ErrDriverIsNotAvailable)
}

func (suite *DriverTruckBindingServiceTestSuite) TestTruckIsNotAvailable() {
	truckDriverDB := db.NewTruckDriverRepository(suite.dbInstance)
	truckDB := db.NewTruckRepository(suite.dbInstance)
	driverTruckBindingDB := db.NewDriverTruckBindingRespository(suite.dbInstance)

	td, err := entity.NewDriver(name, email, license)
	suite.NoError(err)
	err = truckDriverDB.Save(td)
	suite.NoError(err)

	t, err := entity.NewTruck(modelType, manufacturer, licensePlate, fuelType, year)
	err = truckDB.Save(t)
	suite.NoError(err)

	service := NewDriverTruckBindingService(truckDriverDB, truckDB, driverTruckBindingDB)
	var input = dto.CreateBindingDTO{
		IdDriver: td.ID.String(),
		IdTruck:  t.ID.String(),
	}
	_, err = service.BindingDriverToTruck(input)
	suite.NoError(err)

	td, err = entity.NewDriver(name, email, license)
	err = truckDriverDB.Save(td)
	suite.NoError(err)
	input = dto.CreateBindingDTO{
		IdDriver: td.ID.String(),
		IdTruck:  t.ID.String(),
	}

	_, err = service.BindingDriverToTruck(input)
	suite.Error(err)
	suite.Equal(err, ErrTruckIsNotAvailable)
}
