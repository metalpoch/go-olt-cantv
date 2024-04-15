package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/metalpoch/go-olt-cantv/model"
)

type trafficRepository struct {
	db *sql.DB
}

type TrafficRepository interface {
	Save(ctx context.Context, measurements model.Traffic) (int, error)
}

func NewMeasurementsRepository(db *sql.DB) *trafficRepository {
	return &trafficRepository{
		db: db,
	}
}

func (repo trafficRepository) Save(ctx context.Context, traffic model.Traffic) (int, error) {
	res, err := repo.db.ExecContext(
		ctx,
		"INSERT INTO traffic (element_id, date, kbps_in, kbps_out, bandwidth) VALUES(?, ?, ?, ?, ?)",
		traffic.ElementID,
		traffic.Date,
		traffic.KpbsIn,
		traffic.KpbsOut,
		traffic.Bandwidth,
	)

	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("error when trying to save %d traffic: %v", traffic.ElementID, err)
	}

	return int(id), nil
}
