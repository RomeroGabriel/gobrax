package db

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/RomeroGabriel/gobrax-challenge/internal/entity"
)

type TruckRepository struct {
	Db *sql.DB
}

var schemaTruck = `CREATE TABLE IF NOT EXISTS Truck (
    Id TEXT NOT NULL,
    ModelType TEXT NOT NULL,
    Year INTEGER NOT NULL,
	Manufacturer TEXT NOT NULL,
	LicensePlate TEXT NOT NULL,
	Fueltype TEXT NOT NULL,
	PRIMARY KEY (id));`

func NewTruckRepository(db *sql.DB) *TruckRepository {
	err := db.Ping()
	if err != nil {
		log.Printf("%q\n", err)
		return nil
	}
	_, err = db.Exec(schemaTruck)
	if err != nil {
		log.Printf("%q: %s\n", err, schemaTruck)
		return nil
	}
	return &TruckRepository{db}
}

func (r *TruckRepository) Save(truck *entity.Truck) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	var sql = `INSERT INTO Truck (
			Id,
			ModelType,
			Year,
			Manufacturer,
			LicensePlate,
			Fueltype
		)
		VALUES (?,?,?,?,?,?)`
	_, err := r.Db.ExecContext(ctx, sql,
		truck.ID.String(),
		truck.ModelType,
		truck.Year,
		truck.Manufacturer,
		truck.LicensePlate,
		truck.FuelType,
	)
	if err != nil {
		return err
	}
	select {
	case <-ctx.Done():
		log.Println("Context Canceled on Save")
		return context.Canceled
	default:
		return nil
	}
}

func (r *TruckRepository) FindById(id string) (*entity.Truck, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	var sql = `SELECT
		Id,
		ModelType,
		Year,
		Manufacturer,
		LicensePlate,
		Fueltype
		FROM Truck WHERE id = ?;`
	var result entity.Truck
	err := r.Db.QueryRowContext(ctx, sql, id).Scan(
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

func (r *TruckRepository) FindAll() ([]entity.Truck, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	var sql = `SELECT
	id,
	ModelType,
	Year,
	Manufacturer,
	LicensePlate,
	Fueltype
	FROM Truck;`
	rows, err := r.Db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	tAll := []entity.Truck{}
	for rows.Next() {
		var td entity.Truck
		if err := rows.Scan(
			&td.ID,
			&td.ModelType,
			&td.Year,
			&td.Manufacturer,
			&td.LicensePlate,
			&td.FuelType,
		); err != nil {
			return nil, err
		}
		tAll = append(tAll, td)
	}
	select {
	case <-ctx.Done():
		log.Println("Context Canceled on FindAll")
		return nil, context.Canceled
	default:
		return tAll, nil
	}
}

func (r *TruckRepository) Update(truck *entity.Truck) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	var sql = `UPDATE Truck SET ModelType = ?, LicensePlate = ? WHERE id = ?;`
	_, err := r.Db.ExecContext(ctx, sql, truck.ModelType, truck.LicensePlate, truck.ID.String())
	if err != nil {
		return err
	}
	select {
	case <-ctx.Done():
		log.Println("Context Canceled on Update")
		return context.Canceled
	default:
		return nil
	}
}

func (r *TruckRepository) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err := r.Db.ExecContext(ctx, "DELETE FROM Truck WHERE id = ?", id)
	if err != nil {
		return err
	}
	select {
	case <-ctx.Done():
		log.Println("Context Canceled on Delete")
		return nil
	default:
		return nil
	}
}
