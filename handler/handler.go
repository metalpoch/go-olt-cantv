package handler

import (
	"database/sql"

	"github.com/metalpoch/go-olt-cantv/repository"
	"github.com/metalpoch/go-olt-cantv/usecase"
)

func handlerDate(db *sql.DB) *usecase.DateUsecase {
	return usecase.NewDateUsecase(
		repository.NewDateRepository(db),
	)
}

func handlerDevice(db *sql.DB) *usecase.DeviceUsecase {
	return usecase.NewDevicesUsecase(
		repository.NewDeviceRepository(db),
	)
}

func handlerElement(db *sql.DB) *usecase.ElementsUsecase {
	return usecase.NewElementsUsecase(
		repository.NewElementsRepository(db),
	)
}

func handlerCount(db *sql.DB) *usecase.CountUsecase {
	return usecase.NewCountUsecase(
		repository.NewCountsRepository(db),
	)
}

func handlerTraffic(db *sql.DB) *usecase.TrafficUsecase {
	return usecase.NewTrafficUsecase(
		repository.NewMeasurementsRepository(db),
	)
}
