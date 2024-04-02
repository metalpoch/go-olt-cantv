package usecase

import (
	"context"
	"time"

	"github.com/metalpoch/go-olt-cantv/model"
	"github.com/metalpoch/go-olt-cantv/repository"
)

type MeasuremetsUsecase struct {
	Repository repository.MeasurementsRepository
}

func NewMeasurementsUsecase(repo repository.MeasurementsRepository) *MeasuremetsUsecase {
	return &MeasuremetsUsecase{
		Repository: repo,
	}
}

func (m MeasuremetsUsecase) Save(measurement model.SaveMeasurement) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := m.Repository.Save(ctx, measurement)
	if err != nil {
		return err
	}

	return nil
}
