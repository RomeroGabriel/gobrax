package db

import (
	"context"
	"database/sql"
	"log"

	"github.com/RomeroGabriel/gobrax-challenge/internal/entity"
)

type TruckDriverRepository struct {
	Db *sql.DB
}

func NewTruckDriverRepository(db *sql.DB) *TruckDriverRepository {
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS TruckDriver 
		(id TEXT NOT NULL, name TEXT NOT NULL, email TEXT NOT NULL, licenseNumber TEXT NOT NULL);
	`
	ctx := context.Background()
	conn, err := db.Conn(ctx)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return nil
	}
	defer conn.Close()
	_, err = conn.ExecContext(ctx, sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return nil
	}
	return &TruckDriverRepository{db}
}

func (r *TruckDriverRepository) Save(order *entity.TruckDriver) error {
	stmt, err := r.Db.Prepare("INSERT INTO TruckDriver (id, name, email, licenseNumber) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID.String(), order.Name, order.Email, order.LicenseNumber)
	if err != nil {
		return err
	}
	return nil
}
