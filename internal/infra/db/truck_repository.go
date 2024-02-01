package db

import (
	"log"

	"github.com/RomeroGabriel/gobrax-challenge/internal/entity"
	"github.com/jmoiron/sqlx"
)

type TruckRepository struct {
	Db *sqlx.DB
}

var schemaTruck = `CREATE TABLE IF NOT EXISTS Truck (
    id TEXT NOT NULL,
    modeltype TEXT NOT NULL,
    year INTEGER NOT NULL,
	manufacturer TEXT NOT NULL,
	licenseplate TEXT NOT NULL,
	fueltype TEXT NOT NULL,
	PRIMARY KEY (id));`

func NewTruckRepository(db *sqlx.DB) *TruckRepository {
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
	var sql = `INSERT INTO Truck (
			id,
			modeltype,
			year,
			manufacturer,
			licenseplate,
			fueltype
		)
		VALUES (?,?,?,?,?,?)`
	_, err := r.Db.Exec(sql,
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
	return nil
}

func (r *TruckRepository) FindById(id string) (*entity.Truck, error) {
	tDriver := entity.Truck{}
	var sql = `SELECT
			id,
			modeltype,
			year,
			manufacturer,
			licenseplate,
			fueltype
		FROM Truck WHERE id = ?;`
	err := r.Db.Get(&tDriver, sql, id)
	return &tDriver, err
}

func (r *TruckRepository) FindAll() ([]entity.Truck, error) {
	tAll := []entity.Truck{}
	var sql = `SELECT
			id,
			modeltype,
			year,
			manufacturer,
			licenseplate,
			fueltype
		FROM Truck;`
	err := r.Db.Select(&tAll, sql)
	if err != nil {
		return nil, err
	}
	return tAll, err
}

func (r *TruckRepository) Update(truck *entity.Truck) error {
	var sql = `UPDATE Truck SET modeltype = ?, licenseplate = ? WHERE id = ?;`
	_, err := r.Db.Exec(sql, truck.ModelType, truck.LicensePlate, truck.ID.String())
	if err != nil {
		return err
	}
	return nil
}

func (r *TruckRepository) Delete(id string) error {
	_, err := r.FindById(id)
	if err != nil {
		return err
	}
	// TODO: Add logical delete
	_, err = r.Db.Exec("DELETE FROM Truck WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
