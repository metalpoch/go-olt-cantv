package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/metalpoch/go-olt-cantv/model"
)

type deviceRepository struct {
	db *sql.DB
}

func NewDeviceRepository(db *sql.DB) *deviceRepository {
	return &deviceRepository{
		db: db,
	}
}

type DeviceRepository interface {
	Save(ctx context.Context, device model.Device) (int, error)
	FindAll(ctx context.Context) ([]model.Device, error)
}

// return id of new device
func (repo deviceRepository) Save(ctx context.Context, device model.Device) (int, error) {
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

func (repo deviceRepository) FindAll(ctx context.Context) ([]model.Device, error) {
	var devices []model.Device

	rows, err := repo.db.QueryContext(ctx, "SELECT * FROM devices")
	if err != nil {
		return devices, nil
	}
	defer rows.Close()

	for rows.Next() {
		var device model.Device
		err = rows.Scan(&device.ID, &device.IP, &device.Community, &device.Sysname)
		if err != nil {
			return devices, err
		}
		devices = append(devices, device)

	}
	return devices, nil
}
