//go:build !prod
// +build !prod

package db

import (
	"testing"

	"github.com/RomeroGabriel/gobrax-challenge/internal/entity"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type OrderRepositoryTestSuite struct {
	suite.Suite
	Db *sqlx.DB
}

func (suite *OrderRepositoryTestSuite) SetupSuite() {
	db := sqlx.MustOpen("sqlite3", ":memory:")
	suite.Db = db
}

func (suite *OrderRepositoryTestSuite) TearDownSuite() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(OrderRepositoryTestSuite))
}

var name = "Trucker 1"
var email = "email1@gg.com"
var license = "ABC12345"

func (suite *OrderRepositoryTestSuite) TestSaveTruckDriver() {
	tDriver, err := entity.NewTruckDriver(name, email, license)
	suite.NoError(err)
	repo := NewTruckDriverRepository(suite.Db)
	err = repo.Save(tDriver)
	suite.NoError(err)
}

func (suite *OrderRepositoryTestSuite) TestFindByIdTruckDriver() {
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
