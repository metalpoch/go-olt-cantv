package repository

import (
	"context"
	"database/sql"
	"fmt"
)

type dateRepository struct {
	db *sql.DB
}

func NewDateRepository(db *sql.DB) *dateRepository {
	return &dateRepository{
		db: db,
	}
}

type DateRepository interface {
	Save(ctx context.Context, unix int) (int, error)
	FindByID(ctx context.Context, id int) (int, error)
}

func (repo dateRepository) FindByID(ctx context.Context, id int) (int, error) {
	var date int
	err := repo.db.QueryRowContext(ctx, "SELECT unix FROM dates WHERE id=?", id).Scan(&date)
	if err != nil {
		return date, err
	}
	return date, nil
}

func (repo dateRepository) Save(ctx context.Context, unix int) (int, error) {
	res, err := repo.db.ExecContext(ctx, "INSERT INTO dates (unix) VALUES(?)", unix)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error when trying to save the date %d: %v", unix, err)
	}

	return int(id), nil
}
