package usecase

import (
	"context"
	"time"

	"github.com/metalpoch/go-olt-cantv/repository"
)

type DateUsecase struct {
	Repository repository.DateRepository
}

func NewDateUsecase(repo repository.DateRepository) *DateUsecase {
	return &DateUsecase{
		Repository: repo,
	}
}

func (d DateUsecase) Add(unix int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id, err := d.Repository.Save(ctx, unix)

	if err != nil {
		return id, err
	}

	return id, nil
}

func (d DateUsecase) Get(id int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	date, err := d.Repository.FindByID(ctx, id)

	if err != nil {
		return date, err
	}

	return date, nil
}
