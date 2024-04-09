package usecase

import (
	"context"
	"time"

	"github.com/metalpoch/go-olt-cantv/model"
	"github.com/metalpoch/go-olt-cantv/repository"
)

type DeviceUsecase struct {
	Repository repository.DeviceRepository
}

func NewDevicesUsecase(repo repository.DeviceRepository) *DeviceUsecase {
	return &DeviceUsecase{
		Repository: repo,
	}
}

func (d DeviceUsecase) Add(device model.Device) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := d.Repository.Save(ctx, device)

	if err != nil {
		return err
	}

	return nil
}

func (d DeviceUsecase) FindAll() ([]model.Device, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	devices, err := d.Repository.FindAll(ctx)

	if err != nil {
		return nil, err
	}

	return devices, nil
}
