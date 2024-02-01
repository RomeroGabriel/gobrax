package db

import (
	"log"

	"github.com/RomeroGabriel/gobrax-challenge/internal/entity"
	"github.com/jmoiron/sqlx"
)

type TruckDriverRepository struct {
	Db *sqlx.DB
}

var schemaTruckDriver = `CREATE TABLE IF NOT EXISTS TruckDriver (
    id TEXT NOT NULL,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
	licensenumber TEXT NOT NULL,
	PRIMARY KEY (id));`

func NewTruckDriverRepository(db *sqlx.DB) *TruckDriverRepository {
	err := db.Ping()
	if err != nil {
		log.Printf("%q\n", err)
		return nil
	}
	_, err = db.Exec(schemaTruckDriver)
	if err != nil {
		log.Printf("%q: %s\n", err, schemaTruckDriver)
		return nil
	}
	return &TruckDriverRepository{db}
}

func (r *TruckDriverRepository) Save(truckDriver *entity.TruckDriver) error {
	_, err := r.Db.Exec("INSERT INTO TruckDriver (id, name, email, licensenumber) VALUES (?, ?, ?, ?)", truckDriver.ID.String(), truckDriver.Name, truckDriver.Email, truckDriver.LicenseNumber)
	if err != nil {
		return err
	}
	return nil
}

func (r *TruckDriverRepository) FindById(id string) (*entity.TruckDriver, error) {
	tDriver := entity.TruckDriver{}
	err := r.Db.Get(&tDriver, "SELECT id, name, email, licensenumber FROM TruckDriver WHERE id = ?", id)
	return &tDriver, err
}

func (r *TruckDriverRepository) FindAll() ([]entity.TruckDriver, error) {
	tAll := []entity.TruckDriver{}
	err := r.Db.Select(&tAll, "SELECT id, name, email, licensenumber FROM TruckDriver")
	if err != nil {
		return nil, err
	}
	return tAll, err
}

func (r *TruckDriverRepository) Update(truckDriver *entity.TruckDriver) error {
	_, err := r.Db.Exec("UPDATE TruckDriver SET name = ?, email = ?, licensenumber = ?  WHERE id = ?", truckDriver.Name, truckDriver.Email, truckDriver.LicenseNumber, truckDriver.ID.String())
	if err != nil {
		return err
	}
	return nil
}

func (r *TruckDriverRepository) Delete(id string) error {
	_, err := r.FindById(id)
	if err != nil {
		return err
	}
	// TODO: Add logical delete
	_, err = r.Db.Exec("DELETE FROM TruckDriver WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
