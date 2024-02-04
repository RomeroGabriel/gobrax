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

type TruckServiceTestSuite struct {
	suite.Suite
	truckDb    *db.TruckRepository
	dbInstance *sql.DB
}

func (suite *TruckServiceTestSuite) SetupSuite() {
	dbInstance, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	truckDb := db.NewTruckRepository(dbInstance)
	suite.truckDb = truckDb
	suite.dbInstance = dbInstance
}

func TestSuiteTruckervice(t *testing.T) {
	suite.Run(t, new(TruckServiceTestSuite))
}

func (suite *TruckServiceTestSuite) TestDeleteDriverWhenTruckIsLinked() {
	driverDB := db.NewDriverRepository(suite.dbInstance)
	td, err := entity.NewDriver(name, email, license)
	suite.NoError(err)
	err = driverDB.Save(td)
	suite.NoError(err)

	t, err := entity.NewTruck(modelType, manufacturer, licensePlate, fuelType, year)
	err = suite.truckDb.Save(t)
	suite.NoError(err)

	driverTruckBindingDB := db.NewDriverTruckBindingRespository(suite.dbInstance)
	serviceBinding := NewDriverTruckBindingService(driverDB, suite.truckDb, driverTruckBindingDB)
	var input = dto.CreateBindingDTO{
		IdDriver: td.ID.String(),
		IdTruck:  t.ID.String(),
	}
	id, err := serviceBinding.BindingDriverToTruck(input)
	suite.NoError(err)
	suite.NotNil(id)

	service := NewTruckService(suite.truckDb, driverTruckBindingDB)
	_, err = service.Delete(t.ID.String())
	suite.Error(err)
	suite.Equal(ErrTruckHasDriver, err)
}
