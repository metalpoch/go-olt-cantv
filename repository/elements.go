package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/metalpoch/go-olt-cantv/model"
)

type elementsRepository struct {
	db *sql.DB
}

func NewElementsRepository(db *sql.DB) *elementsRepository {
	return &elementsRepository{
		db: db,
	}
}

type ElementsRepository interface {
	Save(ctx context.Context, element model.Elements) (int, error)
	Find(ctx context.Context, element model.Elements) (model.Elements, error)
	FindByID(ctx context.Context, id int) (model.Elements, error)
}

// return id of new element
func (repo elementsRepository) Save(ctx context.Context, element model.Elements) (int, error) {
	res, err := repo.db.ExecContext(ctx, "INSERT INTO elements (port, device_id) VALUES(?, ?)",
		element.Port,
		element.DeviceID,
	)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error when trying to save element %d-%s: %v", element.DeviceID, element.Port, err)
	}

	return int(id), nil
}

func (repo elementsRepository) FindByID(ctx context.Context, id int) (model.Elements, error) {
	var element model.Elements
	err := repo.db.QueryRowContext(ctx, "SELECT * FROM elements WHERE id=?", id).Scan(&element.ID, &element.Port, &element.DeviceID)
	if err != nil {
		return element, err
	}
	return element, nil
}

func (repo elementsRepository) Find(ctx context.Context, element model.Elements) (model.Elements, error) {
	var result model.Elements
	err := repo.db.QueryRowContext(
		ctx,
		"SELECT * FROM elements WHERE port=? and device_id=?",
		element.Port,
		element.DeviceID,
	).Scan(&result.ID, &result.Port, &result.DeviceID)

	if err != nil {
		return result, err
	}

	return result, nil
}
