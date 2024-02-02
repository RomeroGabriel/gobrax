package db

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/RomeroGabriel/gobrax-challenge/internal/entity"
)

type TruckDriverRepository struct {
	Db *sql.DB
}

var schemaTruckDriver = `CREATE TABLE IF NOT EXISTS TruckDriver (
    Id TEXT NOT NULL,
    Name TEXT NOT NULL,
    Email TEXT NOT NULL,
	LicenseNumber TEXT NOT NULL,
	PRIMARY KEY (id));`

func NewTruckDriverRepository(db *sql.DB) *TruckDriverRepository {
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
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	conn, err := acquireConn(ctx, r.Db)
	if err != nil {
		return err
	}
	defer conn.Close()
	_, err = conn.ExecContext(ctx, "INSERT INTO TruckDriver (Id, Name, Email, LicenseNumber) VALUES (?, ?, ?, ?)", truckDriver.ID.String(), truckDriver.Name, truckDriver.Email, truckDriver.LicenseNumber)
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

func (r *TruckDriverRepository) FindById(id string) (*entity.TruckDriver, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	conn, err := acquireConn(ctx, r.Db)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	var result entity.TruckDriver
	var sql = "SELECT Id, Name, Email, LicenseNumber FROM TruckDriver WHERE id = ?"
	err = conn.QueryRowContext(ctx, sql, id).Scan(&result.ID, &result.Name, &result.Email, &result.LicenseNumber)
	err = result.Validate()
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

func (r *TruckDriverRepository) FindAll() ([]entity.TruckDriver, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	conn, err := acquireConn(ctx, r.Db)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	rows, err := conn.QueryContext(ctx, "SELECT Id, Name, Email, LicenseNumber FROM TruckDriver")
	if err != nil {
		return nil, err
	}
	tAll := []entity.TruckDriver{}
	for rows.Next() {
		var td entity.TruckDriver
		if err := rows.Scan(&td.ID, &td.Name, &td.Email, &td.LicenseNumber); err != nil {
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

func (r *TruckDriverRepository) Update(truckDriver *entity.TruckDriver) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	conn, err := acquireConn(ctx, r.Db)
	if err != nil {
		return err
	}
	defer conn.Close()
	var sql = "UPDATE TruckDriver SET Name = ?, Email = ?, LicenseNumber = ?  WHERE Id = ?"
	_, err = conn.ExecContext(ctx, sql, truckDriver.Name, truckDriver.Email, truckDriver.LicenseNumber, truckDriver.ID.String())
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

func (r *TruckDriverRepository) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	conn, err := acquireConn(ctx, r.Db)
	if err != nil {
		return err
	}
	defer conn.Close()
	// TODO: Add logical delete
	_, err = conn.ExecContext(ctx, "DELETE FROM TruckDriver WHERE id = ?", id)
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
