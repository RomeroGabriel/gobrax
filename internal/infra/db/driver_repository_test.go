//go:build !prod
// +build !prod

package db

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/RomeroGabriel/gobrax-challenge/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type DriverRepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

func (suite *DriverRepositoryTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	suite.Db = db
}

func (suite *DriverRepositoryTestSuite) TearDownSuite() {
	suite.Db.Close()
}

func TestSuiteTruckDriver(t *testing.T) {
	suite.Run(t, new(DriverRepositoryTestSuite))
}

var name = "Trucker 1"
var email = "email1@gg.com"
var license = "ABC12345"

func (suite *DriverRepositoryTestSuite) TestSaveTruckDriver() {
	tDriver, err := entity.NewDriver(name, email, license)
	suite.NoError(err)
	repo := NewTruckDriverRepository(suite.Db)
	err = repo.Save(tDriver)
	suite.NoError(err)
}

func (suite *DriverRepositoryTestSuite) TestSaveDuplicateTruckDriver() {
	tDriver, err := entity.NewDriver(name, email, license)
	suite.NoError(err)
	repo := NewTruckDriverRepository(suite.Db)
	err = repo.Save(tDriver)
	suite.NoError(err)

	err = repo.Save(tDriver)
	suite.Error(err)
}

func (suite *DriverRepositoryTestSuite) TestFindByIdTruckDriver() {
	tDriver, err := entity.NewDriver(name, email, license)
	suite.NoError(err)
	repo := NewTruckDriverRepository(suite.Db)
	err = repo.Save(tDriver)
	suite.NoError(err)

	tDriverFind, err := repo.FindById(tDriver.ID.String())
	suite.NoError(err)
	suite.Equal(name, tDriverFind.Name)
	suite.Equal(email, tDriverFind.Email)
	suite.Equal(license, tDriverFind.LicenseNumber)
}

func (suite *DriverRepositoryTestSuite) TestFindByIdTruckDriverNotExist() {
	tDriver, err := entity.NewDriver(name, email, license)
	suite.NoError(err)
	repo := NewTruckDriverRepository(suite.Db)

	tDriverFind, err := repo.FindById(tDriver.ID.String())
	suite.Error(err)
	suite.Nil(tDriverFind)
}

func (suite *DriverRepositoryTestSuite) TestFindAllTruckDriver() {
	repo := NewTruckDriverRepository(suite.Db)

	for i := 0; i < 10; i++ {
		td, err := entity.NewDriver(fmt.Sprintf("TruckDriver %d", i), email, license)
		suite.NoError(err)
		err = repo.Save(td)
		suite.NoError(err)
	}

	tds, err := repo.FindAll()
	suite.NoError(err)
	suite.Len(tds, 10)
	suite.Equal("TruckDriver 0", tds[0].Name)
	suite.Equal("TruckDriver 9", tds[9].Name)
}

func (suite *DriverRepositoryTestSuite) TestUpdateTruckDriver() {
	repo := NewTruckDriverRepository(suite.Db)

	tDriver, err := entity.NewDriver(name, email, license)
	suite.NoError(err)
	err = repo.Save(tDriver)
	suite.NoError(err)

	tDriver.Name = "Update Name"
	tDriver.Email = "Update email"
	tDriver.LicenseNumber = "Update License"
	err = repo.Update(tDriver)

	tDriverFind, err := repo.FindById(tDriver.ID.String())
	suite.NoError(err)
	suite.Equal("Update Name", tDriverFind.Name)
	suite.Equal("Update email", tDriverFind.Email)
	suite.Equal("Update License", tDriverFind.LicenseNumber)
}

func (suite *DriverRepositoryTestSuite) TestDeleteTruckDriver() {
	repo := NewTruckDriverRepository(suite.Db)
	tDriver, err := entity.NewDriver(name, email, license)

	suite.NoError(err)
	err = repo.Save(tDriver)
	suite.NoError(err)

	err = repo.Delete(tDriver.ID.String())

	tDriverFind, err := repo.FindById(tDriver.ID.String())
	suite.Error(err)
	suite.Nil(tDriverFind)
}
