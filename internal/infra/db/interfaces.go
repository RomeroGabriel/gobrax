package db

import "github.com/RomeroGabriel/gobrax-challenge/internal/entity"

type TruckDriverRepositoryInterface interface {
	Save(truckDriver *entity.TruckDriver) error
	FindById(id string) (*entity.TruckDriver, error)
	FindAll() ([]entity.TruckDriver, error)
	Update(truckDriver *entity.TruckDriver) error
	Delete(id string) error
	// TODO: Add Find with pagination
}

type TruckRepositoryInterface interface {
	Save(truck *entity.Truck) error
	FindById(id string) (*entity.Truck, error)
	FindAll() ([]entity.Truck, error)
	Update(truck *entity.Truck) error
	Delete(id string) error
	// TODO: Add Find with pagination
}
