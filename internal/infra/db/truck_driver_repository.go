package db

import (
	"log"

	"github.com/RomeroGabriel/gobrax-challenge/internal/entity"
	"github.com/jmoiron/sqlx"
)

type TruckDriverRepository struct {
	Db *sqlx.DB
}

var schema = `CREATE TABLE IF NOT EXISTS TruckDriver (
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
	_, err = db.Exec(schema)
	if err != nil {
		log.Printf("%q: %s\n", err, schema)
		return nil
	}
	return &TruckDriverRepository{db}
}

func (r *TruckDriverRepository) Save(order *entity.TruckDriver) error {
	_, err := r.Db.Exec("INSERT INTO TruckDriver (id, name, email, licensenumber) VALUES (?, ?, ?, ?)", order.ID.String(), order.Name, order.Email, order.LicenseNumber)
	if err != nil {
		return err
	}
	return nil
}

func (r *TruckDriverRepository) FindById(id string) (*entity.TruckDriver, error) {
	tDriver := entity.TruckDriver{}
	err := r.Db.Get(&tDriver, "SELECT id, name, email, licensenumber FROM TruckDriver WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return &tDriver, err
}
