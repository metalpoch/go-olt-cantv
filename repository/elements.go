package repository

import (
	"context"
	"database/sql"

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
	Save(ctx context.Context, element model.Element) (int, error)
	FindID(ctx context.Context, element model.Element) (int, error)
	FindByID(ctx context.Context, id int) (model.Element, error)
}

func (repo elementsRepository) Save(ctx context.Context, element model.Element) (int, error) {
	res, err := repo.db.ExecContext(ctx, "INSERT INTO elements (device_id, shell, card, port) VALUES(?, ?, ?, ?)",
		element.DeviceID,
		element.Shell,
		element.Card,
		element.Port,
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

func (repo elementsRepository) FindByID(ctx context.Context, id int) (model.Element, error) {
	var element model.Element
	err := repo.db.QueryRowContext(ctx, "SELECT * FROM elements WHERE id=?", id).Scan(&element.DeviceID, &element.Shell, &element.Card, &element.Port)
	if err != nil {
		return element, err
	}
	return element, nil
}

func (repo elementsRepository) FindID(ctx context.Context, element model.Element) (int, error) {
	var id int
	err := repo.db.QueryRowContext(
		ctx,
		"SELECT id FROM elements WHERE shell=? and card=? and port=? and device_id=?",
		element.Shell,
		element.Card,
		element.Port,
		element.DeviceID,
	).Scan(&id)

	if err != nil {
		return id, err
	}

	return id, nil
}
