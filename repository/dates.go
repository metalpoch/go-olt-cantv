package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/metalpoch/go-olt-cantv/model"
)

type datesRepository struct {
	db *sql.DB
}

func NewDatesRepository(db *sql.DB) *datesRepository {
	return &datesRepository{
		db: db,
	}
}

type DatesRepository interface {
	Save(ctx context.Context, date model.Date) (int, error)
}

// return id of new date
func (repo datesRepository) Save(ctx context.Context, date model.Date) (int, error) {
	res, err := repo.db.ExecContext(ctx, "INSERT INTO dates (date) VALUES(?)", date.Date)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error when trying to save date %d: %v", date, err)
	}

	return int(id), nil
}
