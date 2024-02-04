package db

import (
	"context"
	"database/sql"
	"log"

	"github.com/RomeroGabriel/gobrax-challenge/internal/entity"
)

type DriverRepository struct {
	Db *sql.DB
}

var schemaTruckDriver = `CREATE TABLE IF NOT EXISTS Driver (
    Id TEXT NOT NULL,
    Name TEXT NOT NULL,
    Email TEXT NOT NULL,
	LicenseNumber TEXT NOT NULL,
	PRIMARY KEY (id));`

func NewDriverRepository(db *sql.DB) *DriverRepository {
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
	return &DriverRepository{db}
}

func (r *DriverRepository) Save(truckDriver *entity.Driver) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultSecondTimeout)
	defer cancel()
	_, err := r.Db.ExecContext(ctx, "INSERT INTO Driver (Id, Name, Email, LicenseNumber) VALUES (?, ?, ?, ?)", truckDriver.ID.String(), truckDriver.Name, truckDriver.Email, truckDriver.LicenseNumber)
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

func (r *DriverRepository) FindById(id string) (*entity.Driver, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultSecondTimeout)
	defer cancel()
	var result entity.Driver
	var sql = "SELECT Id, Name, Email, LicenseNumber FROM Driver WHERE id = ?"
	err := r.Db.QueryRowContext(ctx, sql, id).Scan(&result.ID, &result.Name, &result.Email, &result.LicenseNumber)
	if err != nil {
		return nil, err
	}
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

func (r *DriverRepository) FindAll() ([]entity.Driver, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultSecondTimeout)
	defer cancel()
	rows, err := r.Db.QueryContext(ctx, "SELECT Id, Name, Email, LicenseNumber FROM Driver")
	if err != nil {
		return nil, err
	}
	tAll := []entity.Driver{}
	for rows.Next() {
		var td entity.Driver
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

func (r *DriverRepository) Update(truckDriver *entity.Driver) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultSecondTimeout)
	defer cancel()
	var sql = "UPDATE Driver SET Name = ?, Email = ?, LicenseNumber = ?  WHERE Id = ?"
	_, err := r.Db.ExecContext(ctx, sql, truckDriver.Name, truckDriver.Email, truckDriver.LicenseNumber, truckDriver.ID.String())
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

func (r *DriverRepository) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultSecondTimeout)
	defer cancel()
	_, err := r.Db.ExecContext(ctx, "DELETE FROM Driver WHERE id = ?", id)
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
