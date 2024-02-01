//go:build !prod
// +build !prod

package db

import (
	"fmt"
	"testing"

	"github.com/RomeroGabriel/gobrax-challenge/internal/entity"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type TruckRepositoryTestSuite struct {
	suite.Suite
	Db *sqlx.DB
}

func (suite *TruckRepositoryTestSuite) SetupSuite() {
	db := sqlx.MustOpen("sqlite3", ":memory:")
	suite.Db = db
}

func (suite *TruckRepositoryTestSuite) TearDownSuite() {
	suite.Db.Close()
}

func TestSuiteTruck(t *testing.T) {
	suite.Run(t, new(TruckRepositoryTestSuite))
}

var modelType = "Model1"
var year uint16 = 1900
var manufacturer = "Volvo"
var licensePlate = "12354"
var fuelType = "Diesel"

func (suite *TruckRepositoryTestSuite) TestSaveTruck() {
	t, err := entity.NewTruck(modelType, manufacturer, licensePlate, fuelType, year)
	suite.NoError(err)
	repo := NewTruckRepository(suite.Db)
	err = repo.Save(t)
	suite.NoError(err)
}

func (suite *TruckRepositoryTestSuite) TestSaveDuplicateTruck() {
	t, err := entity.NewTruck(modelType, manufacturer, licensePlate, fuelType, year)
	suite.NoError(err)
	repo := NewTruckRepository(suite.Db)
	err = repo.Save(t)
	suite.NoError(err)

	err = repo.Save(t)
	suite.Error(err)
}

func (suite *TruckRepositoryTestSuite) TestFindByIdTruck() {
	t, err := entity.NewTruck(modelType, manufacturer, licensePlate, fuelType, year)
	suite.NoError(err)
	repo := NewTruckRepository(suite.Db)
	err = repo.Save(t)
	suite.NoError(err)

	tDriverFind, err := repo.FindById(t.ID.String())
	suite.NoError(err)
	suite.Equal(modelType, tDriverFind.ModelType)
	suite.Equal(year, tDriverFind.Year)
	suite.Equal(manufacturer, tDriverFind.Manufacturer)
	suite.Equal(licensePlate, tDriverFind.LicensePlate)
	suite.Equal(fuelType, tDriverFind.FuelType)
}

func (suite *TruckRepositoryTestSuite) TestFindByIdTruckNotExist() {
	t, err := entity.NewTruck(modelType, manufacturer, licensePlate, fuelType, year)
	suite.NoError(err)
	repo := NewTruckRepository(suite.Db)

	tDriverFind, err := repo.FindById(t.ID.String())
	suite.Error(err)
	suite.Equal("", tDriverFind.ModelType)
	suite.Equal(uint16(0), tDriverFind.Year)
	suite.Equal("", tDriverFind.Manufacturer)
	suite.Equal("", tDriverFind.LicensePlate)
	suite.Equal("", tDriverFind.FuelType)
}

func (suite *TruckRepositoryTestSuite) TestFindAllTruck() {
	repo := NewTruckRepository(suite.Db)

	for i := 0; i < 10; i++ {
		td, err := entity.NewTruck(fmt.Sprintf("Truck %d", i), manufacturer, licensePlate, fuelType, year)
		suite.NoError(err)
		err = repo.Save(td)
		suite.NoError(err)
	}

	tds, err := repo.FindAll()
	suite.NoError(err)
	suite.Len(tds, 10)
	suite.Equal("Truck 0", tds[0].ModelType)
	suite.Equal("Truck 9", tds[9].ModelType)
}

func (suite *TruckRepositoryTestSuite) TestUpdateTruck() {
	repo := NewTruckRepository(suite.Db)

	t, err := entity.NewTruck(modelType, manufacturer, licensePlate, fuelType, year)
	suite.NoError(err)
	err = repo.Save(t)
	suite.NoError(err)

	t.ModelType = "Update Model"
	t.LicensePlate = "Update LicensePlate"
	err = repo.Update(t)

	tDriverFind, err := repo.FindById(t.ID.String())
	suite.NoError(err)
	suite.Equal("Update Model", tDriverFind.ModelType)
	suite.Equal("Update LicensePlate", tDriverFind.LicensePlate)
}

func (suite *TruckRepositoryTestSuite) TestDeleteTruck() {
	repo := NewTruckRepository(suite.Db)
	t, err := entity.NewTruck(modelType, manufacturer, licensePlate, fuelType, year)

	suite.NoError(err)
	err = repo.Save(t)
	suite.NoError(err)

	err = repo.Delete(t.ID.String())

	tFind, err := repo.FindById(t.ID.String())
	suite.Error(err)
	suite.Equal("", tFind.ModelType)
	suite.Equal(uint16(0), tFind.Year)
	suite.Equal("", tFind.Manufacturer)
	suite.Equal("", tFind.LicensePlate)
	suite.Equal("", tFind.FuelType)
}
