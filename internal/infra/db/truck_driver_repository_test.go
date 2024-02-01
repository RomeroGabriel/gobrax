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

type TruckDriverRepositoryTestSuite struct {
	suite.Suite
	Db *sqlx.DB
}

func (suite *TruckDriverRepositoryTestSuite) SetupSuite() {
	db := sqlx.MustOpen("sqlite3", ":memory:")
	suite.Db = db
}

func (suite *TruckDriverRepositoryTestSuite) TearDownSuite() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(TruckDriverRepositoryTestSuite))
}

var name = "Trucker 1"
var email = "email1@gg.com"
var license = "ABC12345"

func (suite *TruckDriverRepositoryTestSuite) TestSaveTruckDriver() {
	tDriver, err := entity.NewTruckDriver(name, email, license)
	suite.NoError(err)
	repo := NewTruckDriverRepository(suite.Db)
	err = repo.Save(tDriver)
	suite.NoError(err)
}

func (suite *TruckDriverRepositoryTestSuite) TestSaveDuplicateTruckDriver() {
	tDriver, err := entity.NewTruckDriver(name, email, license)
	suite.NoError(err)
	repo := NewTruckDriverRepository(suite.Db)
	err = repo.Save(tDriver)
	suite.NoError(err)

	err = repo.Save(tDriver)
	suite.Error(err)
}

func (suite *TruckDriverRepositoryTestSuite) TestFindByIdTruckDriver() {
	tDriver, err := entity.NewTruckDriver(name, email, license)
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

func (suite *TruckDriverRepositoryTestSuite) TestFindByIdTruckDriverNotExist() {
	tDriver, err := entity.NewTruckDriver(name, email, license)
	suite.NoError(err)
	repo := NewTruckDriverRepository(suite.Db)

	tDriverFind, err := repo.FindById(tDriver.ID.String())
	suite.Error(err)
	suite.Equal("", tDriverFind.Name)
	suite.Equal("", tDriverFind.Email)
	suite.Equal("", tDriverFind.LicenseNumber)
}

func (suite *TruckDriverRepositoryTestSuite) TestFindAllTruckDriver() {
	repo := NewTruckDriverRepository(suite.Db)

	for i := 0; i < 10; i++ {
		td, err := entity.NewTruckDriver(fmt.Sprintf("TruckDriver %d", i), email, license)
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

func (suite *TruckDriverRepositoryTestSuite) TestUpdateTruckDriver() {
	repo := NewTruckDriverRepository(suite.Db)

	tDriver, err := entity.NewTruckDriver(name, email, license)
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

func (suite *TruckDriverRepositoryTestSuite) TestDeleteTruckDriver() {
	repo := NewTruckDriverRepository(suite.Db)
	tDriver, err := entity.NewTruckDriver(name, email, license)

	suite.NoError(err)
	err = repo.Save(tDriver)
	suite.NoError(err)

	err = repo.Delete(tDriver.ID.String())

	tDriverFind, err := repo.FindById(tDriver.ID.String())
	suite.Error(err)
	suite.Equal("", tDriverFind.Name)
	suite.Equal("", tDriverFind.Email)
	suite.Equal("", tDriverFind.LicenseNumber)
}
