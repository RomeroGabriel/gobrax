package service

import "github.com/RomeroGabriel/gobrax-challenge/internal/infra/db"

type DriverTruckBindingService struct {
	TruckDriverDB db.TruckDriverRepositoryInterface
	TruckDB       db.TruckRepositoryInterface
}

func NewDriverTruckBindingService(
	tdDB db.TruckDriverRepositoryInterface,
	tDB db.TruckRepositoryInterface,
) *DriverTruckBindingService {
	return &DriverTruckBindingService{TruckDriverDB: tdDB, TruckDB: tDB}
}
