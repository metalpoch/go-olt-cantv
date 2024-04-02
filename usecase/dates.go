package usecase

import (
	"context"
	"time"

	"github.com/metalpoch/go-olt-cantv/model"
	"github.com/metalpoch/go-olt-cantv/repository"
)

type DateUsecase struct {
	Repository repository.DatesRepository
}

func NewDateUsecase(repo repository.DatesRepository) *DateUsecase {
	return &DateUsecase{
		Repository: repo,
	}
}

func (d DateUsecase) Add(date model.Date) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id, err := d.Repository.Save(ctx, date)

	if err != nil {
		return id, err
	}

	return id, nil
}
