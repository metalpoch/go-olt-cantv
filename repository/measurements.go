package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/metalpoch/go-olt-cantv/model"
)

type measurementsRepository struct {
	db *sql.DB
}

type MeasurementsRepository interface {
	Save(ctx context.Context, measurements model.SaveMeasurement) (int, error)
}

func NewMeasurementsRepository(db *sql.DB) *measurementsRepository {
	return &measurementsRepository{
		db: db,
	}
}

func (repo measurementsRepository) Save(ctx context.Context, measurement model.SaveMeasurement) (int, error) {
	res, err := repo.db.ExecContext(
		ctx,
		"INSERT INTO measurements (bytes_in, bytes_out, bandwidth, date_id, elements_id) VALUES(?, ?, ?, ?, ?)",
		measurement.ByteIn,
		measurement.ByteOut,
		measurement.Bandwidth,
		measurement.DateID,
		measurement.ElementID,
	)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error when trying to save %d measurement: %v", measurement.ElementID, err)
	}

	return int(id), nil
}
