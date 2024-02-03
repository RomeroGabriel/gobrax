package service

import (
	"errors"

	"github.com/RomeroGabriel/gobrax-challenge/internal/dto"
	"github.com/RomeroGabriel/gobrax-challenge/internal/infra/db"
	pkg "github.com/RomeroGabriel/gobrax-challenge/pkg/entity"
)

type DriverTruckBindingService struct {
	TruckDriverDB        db.ITruckDriverRepository
	TruckDB              db.ITruckRepository
	DriverTruckBindingDB db.IDriverTruckBindingRespository
}

func NewDriverTruckBindingService(
	tdDB db.ITruckDriverRepository,
	tDB db.ITruckRepository,
	tddDB db.IDriverTruckBindingRespository,
) *DriverTruckBindingService {
	return &DriverTruckBindingService{TruckDriverDB: tdDB, TruckDB: tDB, DriverTruckBindingDB: tddDB}
}

var (
	ErrDriverIsNotAvailable         = errors.New("driver is not available")
	ErrTruckIsNotAvailable          = errors.New("truck is not available")
	ErrTruckAndDriverIsNotAvailable = errors.New("truck and driver are not available")
)

func (s *DriverTruckBindingService) BindingDriverToTruck(input dto.CreateBindingDTO) (*pkg.ID, error) {
	idTruck, _ := pkg.ParseID(input.IdTruck)
	truck, err := s.TruckDB.FindById(idTruck.String())
	if err != nil {
		return nil, err
	}
	idDriver, _ := pkg.ParseID(input.IdDriver)
	driver, err := s.TruckDriverDB.FindById(idDriver.String())
	if err != nil {
		return nil, err
	}
	driverIsAvailable, err := s.DriverTruckBindingDB.DriverIsAvailable(*driver)
	truckIsAvailable, err := s.DriverTruckBindingDB.TruckIsAvailable(*truck)

	if !driverIsAvailable && !truckIsAvailable {
		return nil, ErrTruckAndDriverIsNotAvailable
	}
	if !driverIsAvailable {
		return nil, ErrDriverIsNotAvailable
	}
	if !truckIsAvailable {
		return nil, ErrTruckIsNotAvailable
	}
	return s.DriverTruckBindingDB.CreateBinding(*truck, *driver)
}
