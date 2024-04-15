package repository

import (
	"context"
	"database/sql"

	"github.com/metalpoch/go-olt-cantv/entity"
	"github.com/metalpoch/go-olt-cantv/model"
)

type countRepository struct {
	db *sql.DB
}

type CountRepository interface {
	Save(ctx context.Context, counts model.Count) (int, error)
	Remove(ctx context.Context, elementID int) error
	FindPreviouCount(ctx context.Context, elementID int) (model.Count, error)
}

func NewCountsRepository(db *sql.DB) *countRepository {
	return &countRepository{
		db: db,
	}
}

func (repo countRepository) Save(ctx context.Context, count model.Count) (int, error) {
	res, err := repo.db.ExecContext(
		ctx,
		"INSERT INTO tempcounts (element_id, date, bytes_in, bytes_out, bandwidth) VALUES(?, ?, ?, ?, ?)",
		count.ElementID,
		count.Date,
		count.BytesIn,
		count.BytesOut,
		count.Bandwidth,
	)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (repo countRepository) FindPreviouCount(ctx context.Context, elementID int) (model.Count, error) {
	count := entity.Count{}
	err := repo.db.QueryRowContext(ctx, "SELECT * FROM tempcounts WHERE element_id=?", elementID).Scan(
		&count.ElementID,
		&count.Date,
		&count.BytesIn,
		&count.BytesOut,
		&count.Bandwidth,
	)
	if err != nil {
		return model.Count{}, err
	}
	return model.Count{
		ElementID: count.ElementID,
		Date:      count.Date,
		BytesIn:   count.BytesIn,
		BytesOut:  count.BytesOut,
		Bandwidth: count.Bandwidth,
	}, nil
}

func (repo countRepository) Remove(ctx context.Context, elementID int) error {
	if _, err := repo.db.ExecContext(ctx, "DELETE FROM tempcounts WHERE element_id=?", elementID); err != nil {
		return err
	}
	return nil
}
