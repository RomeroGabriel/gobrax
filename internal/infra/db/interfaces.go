package db

import (
	"github.com/RomeroGabriel/gobrax-challenge/internal/entity"
	pkg "github.com/RomeroGabriel/gobrax-challenge/pkg/entity"
)

type IDriverRepository interface {
	Save(truckDriver *entity.Driver) error
	FindById(id string) (*entity.Driver, error)
	FindAll() ([]entity.Driver, error)
	Update(truckDriver *entity.Driver) error
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
	CreateBinding(truck entity.Truck, driver entity.Driver) (*pkg.ID, error)
	RemoveBinding(truck entity.Truck, driver entity.Driver) error
	RemoveBindingById(id string) error
	GetCurrentTruckOfDriver(driver entity.Driver) (*entity.Truck, error)
	DriverIsAvailable(driver entity.Driver) (bool, error)
	TruckIsAvailable(truck entity.Truck) (bool, error)
}
