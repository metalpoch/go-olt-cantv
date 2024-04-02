package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/metalpoch/go-olt-cantv/entity"
	"github.com/metalpoch/go-olt-cantv/model"
)

type devicesRepository struct {
	db *sql.DB
}

func NewDevicesRepository(db *sql.DB) *devicesRepository {
	return &devicesRepository{
		db: db,
	}
}

type DevicesRepository interface {
	Save(ctx context.Context, device model.Device) (int, error)
	FindAll(ctx context.Context) ([]entity.Devices, error)
}

// return id of new device
func (repo devicesRepository) Save(ctx context.Context, device model.Device) (int, error) {
	res, err := repo.db.ExecContext(
		ctx,
		"INSERT INTO devices (ip, community, sysname) VALUES(?, ?, ?)",
		device.IP,
		device.Community,
		device.Sysname,
	)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error when trying to save device %s: %v", device.IP, err)
	}

	return int(id), nil
}

func (repo devicesRepository) FindAll(ctx context.Context) ([]entity.Devices, error) {
	var devices []entity.Devices

	rows, err := repo.db.QueryContext(ctx, "SELECT * FROM devices")
	if err != nil {
		return devices, nil
	}
	defer rows.Close()

	for rows.Next() {
		var device entity.Devices
		err = rows.Scan(&device.ID, &device.IP, &device.Community, &device.Sysname)
		if err != nil {
			return devices, err
		}
		devices = append(devices, device)

	}
	return devices, nil
}
