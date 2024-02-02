//go:build !prod
// +build !prod

package db

import (
	"database/sql"
	"testing"

	"github.com/RomeroGabriel/gobrax-challenge/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type DriverTruckBindingRespositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

func (suite *DriverTruckBindingRespositoryTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	suite.Db = db
}

func (suite *DriverTruckBindingRespositoryTestSuite) TearDownSuite() {
	suite.Db.Close()
}

func TestSuiteDriverTruckBinding(t *testing.T) {
	suite.Run(t, new(DriverTruckBindingRespositoryTestSuite))
}

func (suite *DriverTruckBindingRespositoryTestSuite) TestCreate() {
	t, err := entity.NewTruck(modelType, manufacturer, licensePlate, fuelType, year)
	suite.NoError(err)

	td, err := entity.NewTruckDriver(name, email, license)
	suite.NoError(err)

	repo := NewDriverTruckBindingRespository(suite.Db)
	_, err = repo.CreateBinding(*t, *td)
	suite.NoError(err)
}

func (suite *DriverTruckBindingRespositoryTestSuite) TestGetCurrentTruckOfDriver() {
	t, err := entity.NewTruck(modelType, manufacturer, licensePlate, fuelType, year)
	suite.NoError(err)
	tRepo := NewTruckRepository(suite.Db)
	err = tRepo.Save(t)
	suite.NoError(err)

	td, err := entity.NewTruckDriver(name, email, license)
	suite.NoError(err)
	tdRepo := NewTruckDriverRepository(suite.Db)
	err = tdRepo.Save(td)
	suite.NoError(err)

	repo := NewDriverTruckBindingRespository(suite.Db)
	_, err = repo.CreateBinding(*t, *td)
	suite.NoError(err)

	truck, err := repo.GetCurrentTruckOfDriver(*td)
	suite.NoError(err)
	suite.Equal(truck.ID.String(), t.ID.String())
}

func (suite *DriverTruckBindingRespositoryTestSuite) TestRemoveBinding() {
	t, err := entity.NewTruck(modelType, manufacturer, licensePlate, fuelType, year)
	suite.NoError(err)
	tRepo := NewTruckRepository(suite.Db)
	err = tRepo.Save(t)
	suite.NoError(err)

	td, err := entity.NewTruckDriver(name, email, license)
	suite.NoError(err)
	tdRepo := NewTruckDriverRepository(suite.Db)
	err = tdRepo.Save(td)
	suite.NoError(err)

	repo := NewDriverTruckBindingRespository(suite.Db)
	_, err = repo.CreateBinding(*t, *td)
	suite.NoError(err)

	err = repo.RemoveBinding(*t, *td)
	suite.NoError(err)
}

func (suite *DriverTruckBindingRespositoryTestSuite) TestRemoveBindingById() {
	t, err := entity.NewTruck(modelType, manufacturer, licensePlate, fuelType, year)
	suite.NoError(err)
	tRepo := NewTruckRepository(suite.Db)
	err = tRepo.Save(t)
	suite.NoError(err)

	td, err := entity.NewTruckDriver(name, email, license)
	suite.NoError(err)
	tdRepo := NewTruckDriverRepository(suite.Db)
	err = tdRepo.Save(td)
	suite.NoError(err)

	repo := NewDriverTruckBindingRespository(suite.Db)
	id, err := repo.CreateBinding(*t, *td)
	suite.NoError(err)

	err = repo.RemoveBindingById(id.String())
	suite.NoError(err)
}

func (suite *DriverTruckBindingRespositoryTestSuite) TestDriverWithoutCurrentTruck() {
	t, err := entity.NewTruck(modelType, manufacturer, licensePlate, fuelType, year)
	suite.NoError(err)
	tRepo := NewTruckRepository(suite.Db)
	err = tRepo.Save(t)
	suite.NoError(err)

	td, err := entity.NewTruckDriver(name, email, license)
	suite.NoError(err)
	tdRepo := NewTruckDriverRepository(suite.Db)
	err = tdRepo.Save(td)
	suite.NoError(err)

	repo := NewDriverTruckBindingRespository(suite.Db)

	truck, err := repo.GetCurrentTruckOfDriver(*td)
	suite.Error(err)
	suite.Nil(truck)
}

func (suite *DriverTruckBindingRespositoryTestSuite) TestDriverHasTruck() {
	t, err := entity.NewTruck(modelType, manufacturer, licensePlate, fuelType, year)
	suite.NoError(err)
	tRepo := NewTruckRepository(suite.Db)
	err = tRepo.Save(t)
	suite.NoError(err)

	td, err := entity.NewTruckDriver(name, email, license)
	suite.NoError(err)
	tdRepo := NewTruckDriverRepository(suite.Db)
	err = tdRepo.Save(td)
	suite.NoError(err)

	repo := NewDriverTruckBindingRespository(suite.Db)
	_, err = repo.CreateBinding(*t, *td)
	suite.NoError(err)

	hasValue, err := repo.DriverHasTruck(*td)
	suite.NoError(err)
	suite.Equal(true, hasValue)
}

func (suite *DriverTruckBindingRespositoryTestSuite) TestTruckIsNotAvailable() {
	t, err := entity.NewTruck(modelType, manufacturer, licensePlate, fuelType, year)
	suite.NoError(err)
	tRepo := NewTruckRepository(suite.Db)
	err = tRepo.Save(t)
	suite.NoError(err)

	td, err := entity.NewTruckDriver(name, email, license)
	suite.NoError(err)
	tdRepo := NewTruckDriverRepository(suite.Db)
	err = tdRepo.Save(td)
	suite.NoError(err)

	repo := NewDriverTruckBindingRespository(suite.Db)
	_, err = repo.CreateBinding(*t, *td)
	suite.NoError(err)

	hasValue, err := repo.TruckIsAvailable(*t)
	suite.NoError(err)
	suite.Equal(false, hasValue)
}

func (suite *DriverTruckBindingRespositoryTestSuite) TestTruckIsAvailable() {
	t, err := entity.NewTruck(modelType, manufacturer, licensePlate, fuelType, year)
	suite.NoError(err)
	tRepo := NewTruckRepository(suite.Db)
	err = tRepo.Save(t)
	suite.NoError(err)

	td, err := entity.NewTruckDriver(name, email, license)
	suite.NoError(err)
	tdRepo := NewTruckDriverRepository(suite.Db)
	err = tdRepo.Save(td)
	suite.NoError(err)

	repo := NewDriverTruckBindingRespository(suite.Db)
	hasValue, err := repo.TruckIsAvailable(*t)
	suite.NoError(err)
	suite.Equal(true, hasValue)
}
