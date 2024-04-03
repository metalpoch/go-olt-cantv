package repository

import (
	"context"
	"database/sql"

	"github.com/metalpoch/go-olt-cantv/entity"
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
	Find(ctx context.Context, element model.Element) (model.Element, error)
	FindByID(ctx context.Context, id int) (model.Element, error)
}

func (repo elementsRepository) Save(ctx context.Context, element model.Element) (int, error) {
	newElement := entity.Elements{
		Shell:    element.Shell,
		Card:     element.Card,
		Port:     element.Port,
		DeviceID: element.DeviceID,
	}

	res, err := repo.db.ExecContext(ctx, "INSERT INTO elements (shell, card, port, device_id) VALUES(?, ?, ?, ?)",
		newElement.Shell,
		newElement.Card,
		newElement.Port,
		newElement.DeviceID,
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
	err := repo.db.QueryRowContext(ctx, "SELECT * FROM elements WHERE id=?", id).Scan(&element.ID, &element.Shell, &element.Card, &element.Port, &element.DeviceID)
	if err != nil {
		return element, err
	}
	return element, nil
}

func (repo elementsRepository) Find(ctx context.Context, element model.Element) (model.Element, error) {
	var result model.Element
	err := repo.db.QueryRowContext(
		ctx,
		"SELECT * FROM elements WHERE shell=? and card=? and port=? and device_id=?",
		element.Shell,
		element.Card,
		element.Port,
		element.DeviceID,
	).Scan(&result.ID, &result.Shell, &result.Card, &result.Port, &result.DeviceID)

	if err != nil {
		return result, err
	}

	return result, nil
}
