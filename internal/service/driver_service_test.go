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

type DriverServiceTestSuite struct {
	suite.Suite
	driverDB   *db.DriverRepository
	dbInstance *sql.DB
}

func (suite *DriverServiceTestSuite) SetupSuite() {
	dbInstance, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	driverDB := db.NewDriverRepository(dbInstance)
	suite.driverDB = driverDB
	suite.dbInstance = dbInstance
}

func TestSuiteDriverService(t *testing.T) {
	suite.Run(t, new(DriverServiceTestSuite))
}

func (suite *DriverServiceTestSuite) TestDeleteDriverWhenTruckIsLinked() {
	truckDB := db.NewTruckRepository(suite.dbInstance)
	driverTruckBindingDB := db.NewDriverTruckBindingRespository(suite.dbInstance)

	td, err := entity.NewDriver(name, email, license)
	suite.NoError(err)
	err = suite.driverDB.Save(td)
	suite.NoError(err)

	t, err := entity.NewTruck(modelType, manufacturer, licensePlate, fuelType, year)
	err = truckDB.Save(t)
	suite.NoError(err)

	serviceBinding := NewDriverTruckBindingService(*suite.driverDB, *truckDB, *driverTruckBindingDB)
	var input = dto.CreateBindingDTO{
		IdDriver: td.ID.String(),
		IdTruck:  t.ID.String(),
	}
	id, err := serviceBinding.BindingDriverToTruck(input)
	suite.NoError(err)
	suite.NotNil(id)

	service := NewDriverService(suite.driverDB, driverTruckBindingDB)
	_, err = service.Delete(td.ID.String())
	suite.Error(err)
	suite.Equal(ErrDriverHasTruck, err)
}
