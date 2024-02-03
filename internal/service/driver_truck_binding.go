package service

import "github.com/RomeroGabriel/gobrax-challenge/internal/infra/db"

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
	return &DriverTruckBindingService{TruckDriverDB: tdDB, TruckDB: tDB}
}
