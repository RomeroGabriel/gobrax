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

func (r *DriverTruckBindingRespository) CreateBinding(truck entity.Truck, driver entity.Driver) (*pkg.ID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultSecondTimeout)
	defer cancel()

	var sql = `INSERT INTO DriverTruckMapping (
		Id,
		FkDriver,
		FkTruck,
		CreatedAt
	)
	VALUES (?,?,?,?)`

	id := pkg.NewID()
	_, err := r.Db.ExecContext(ctx, sql,
		id.String(),
		driver.ID.String(),
		truck.ID.String(),
		time.Now().Format(stringFormat),
	)
	if err != nil {
		return nil, err
	}

	select {
	case <-ctx.Done():
		log.Println("Context Canceled on FindById")
		return nil, context.Canceled
	default:
		return &id, nil
	}
}

func (r *DriverTruckBindingRespository) RemoveBinding(truck entity.Truck, driver entity.Driver) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultSecondTimeout)
	defer cancel()
	var sql = `UPDATE DriverTruckMapping SET DeletedAt = ? WHERE FkDriver = ? AND FkTruck = ? AND DeletedAt IS NULL`
	var deletedAt = time.Now().Format(stringFormat)
	_, err := r.Db.ExecContext(ctx, sql, deletedAt, driver.ID.String(), truck.ID.String())
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

func (r *DriverTruckBindingRespository) RemoveBindingById(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultSecondTimeout)
	defer cancel()
	var sql = `UPDATE DriverTruckMapping SET DeletedAt = ? WHERE Id = ?`
	var deletedAt = time.Now().Format(stringFormat)
	_, err := r.Db.ExecContext(ctx, sql, deletedAt, id)
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

func (r *DriverTruckBindingRespository) GetCurrentTruckOfDriver(driver entity.Driver) (*entity.Truck, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultSecondTimeout)
	defer cancel()
	var sql = `SELECT
		truck.Id,
		truck.ModelType,
		truck.Year,
		truck.Manufacturer,
		truck.LicensePlate,
		truck.Fueltype
	FROM DriverTruckMapping AS map
		INNER JOIN Truck as truck ON truck.Id = map.FkTruck AND map.FkDriver = ? and DeletedAt IS NULL;
	`
	var result entity.Truck
	row := r.Db.QueryRowContext(ctx, sql, driver.ID.String())
	err := row.Scan(
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

func (r *DriverTruckBindingRespository) DriverIsAvailable(driver entity.Driver) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultSecondTimeout)
	defer cancel()
	var count int
	var sql = "SELECT COUNT(*) FROM DriverTruckMapping WHERE FkDriver = ? AND DeletedAt IS NULL;"
	err := r.Db.QueryRowContext(ctx, sql, driver.ID.String()).Scan(&count)
	if err != nil {
		return false, err
	}
	return !(count == 1), err
}

func (r *DriverTruckBindingRespository) TruckIsAvailable(truck entity.Truck) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultSecondTimeout)
	defer cancel()
	var count int
	var sql = "SELECT COUNT(*) FROM DriverTruckMapping WHERE FkTruck = ? AND DeletedAt IS NULL;"
	err := r.Db.QueryRowContext(ctx, sql, truck.ID.String()).Scan(&count)
	if err != nil {
		return false, err
	}
	return !(count == 1), err
}
