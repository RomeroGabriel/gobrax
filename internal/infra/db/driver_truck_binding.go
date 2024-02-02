package db

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/RomeroGabriel/gobrax-challenge/internal/entity"
	pkg "github.com/RomeroGabriel/gobrax-challenge/pkg/entity"
)

type DriverTruckBindingRespository struct {
	Db *sql.DB
}

var schemaDriverTruckBinding = `CREATE TABLE IF NOT EXISTS DriverTruckMapping (
    Id TEXT NOT NULL,
    FkDriver TEXT NOT NULL,
    FkTruck TEXT NOT NULL,
	CreatedAt TEXT NOT NULL,
	DeletedAt TEXT,
	PRIMARY KEY (id),
	FOREIGN KEY(FkDriver) REFERENCES TruckDriver(Id),
	FOREIGN KEY(FkTruck) REFERENCES Truck(Id));`

func NewDriverTruckBindingRespository(db *sql.DB) *DriverTruckBindingRespository {
	err := db.Ping()
	if err != nil {
		log.Printf("%q\n", err)
		return nil
	}
	_, err = db.Exec(schemaDriverTruckBinding)
	if err != nil {
		log.Printf("%q: %s\n", err, schemaDriverTruckBinding)
		return nil
	}
	return &DriverTruckBindingRespository{db}
}

func (r *DriverTruckBindingRespository) CreateBinding(truck entity.Truck, driver entity.TruckDriver) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	conn, err := acquireConn(ctx, r.Db)
	if err != nil {
		return err
	}
	defer conn.Close()
	if err != nil {
		return err
	}

	var sql = `INSERT INTO DriverTruckMapping (
		Id,
		FkDriver,
		FkTruck,
		CreatedAt
	)
	VALUES (?,?,?,?)`

	_, err = conn.ExecContext(ctx, sql,
		pkg.NewID().String(),
		driver.ID.String(),
		truck.ID.String(),
		time.Now().Format("2006-01-02 15:04:05"),
	)
	if err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		log.Println("Context Canceled on FindById")
		return context.Canceled
	default:
		return nil
	}
}

func (r *DriverTruckBindingRespository) GetCurrentTruckOfDriver(driver entity.TruckDriver) (*entity.Truck, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	conn, err := acquireConn(ctx, r.Db)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	if err != nil {
		return nil, err
	}

	var sql = `SELECT
		truck.Id,
		truck.ModelType,
		truck.Year,
		truck.Manufacturer,
		truck.LicensePlate,
		truck.Fueltype
	FROM DriverTruckMapping AS map
		LEFT JOIN Truck as truck ON truck.Id = map.FkTruck AND map.FkDriver = ? and DeletedAt IS NULL;
	`
	var result entity.Truck
	err = conn.QueryRowContext(ctx, sql, driver.ID.String()).Scan(
		&result.ID,
		&result.ModelType,
		&result.Year,
		&result.Manufacturer,
		&result.LicensePlate,
		&result.FuelType,
	)
	if err != nil {
		return nil, err
	}
	select {
	case <-ctx.Done():
		log.Println("Context Canceled on FindById")
		return nil, context.Canceled
	default:
		return &result, nil
	}
}
