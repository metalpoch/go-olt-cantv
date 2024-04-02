package handler

import (
	"database/sql"

	"github.com/metalpoch/go-olt-cantv/repository"
	"github.com/metalpoch/go-olt-cantv/usecase"
)

func handlerDevice(db *sql.DB) *usecase.DevicesUsecase {
	return usecase.NewDevicesUsecase(
		repository.NewDevicesRepository(db),
	)
}

func handlerDate(db *sql.DB) *usecase.DateUsecase {
	return usecase.NewDateUsecase(
		repository.NewDatesRepository(db),
	)
}

func handlerMeasurement(db *sql.DB) *usecase.MeasuremetsUsecase {
	return usecase.NewMeasurementsUsecase(
		repository.NewMeasurementsRepository(db),
	)
}

func handlerElement(db *sql.DB) *usecase.ElementsUsecase {
	return usecase.NewElementsUsecase(
		repository.NewElementsRepository(db),
	)
}
