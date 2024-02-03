package db

import (
	"github.com/RomeroGabriel/gobrax-challenge/internal/entity"
	pkg "github.com/RomeroGabriel/gobrax-challenge/pkg/entity"
)

type ITruckDriverRepository interface {
	Save(truckDriver *entity.TruckDriver) error
	FindById(id string) (*entity.TruckDriver, error)
	FindAll() ([]entity.TruckDriver, error)
	Update(truckDriver *entity.TruckDriver) error
	Delete(id string) error
	// TODO: Add Find with pagination
}

type ITruckRepository interface {
	Save(truck *entity.Truck) error
	FindById(id string) (*entity.Truck, error)
	FindAll() ([]entity.Truck, error)
	Update(truck *entity.Truck) error
	Delete(id string) error
	// TODO: Add Find with pagination
}

type IDriverTruckBindingRespository interface {
	CreateBinding(truck entity.Truck, driver entity.TruckDriver) (*pkg.ID, error)
	RemoveBinding(truck entity.Truck, driver entity.TruckDriver) error
	RemoveBindingById(id string) error
	GetCurrentTruckOfDriver(driver entity.TruckDriver) (*entity.Truck, error)
	DriverHasTruck(driver entity.TruckDriver) (bool, error)
	TruckIsAvailable(truck entity.Truck) (bool, error)
}
