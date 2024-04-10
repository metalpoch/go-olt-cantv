package usecase

import (
	"context"
	"time"

	"github.com/metalpoch/go-olt-cantv/model"
	"github.com/metalpoch/go-olt-cantv/repository"
)

type ElementsUsecase struct {
	Repository repository.ElementsRepository
}

func NewElementsUsecase(repo repository.ElementsRepository) *ElementsUsecase {
	return &ElementsUsecase{
		Repository: repo,
	}
}

func (e ElementsUsecase) FindID(element model.Element) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id, err := e.Repository.FindID(ctx, element)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (e ElementsUsecase) Save(newElement model.Element) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id, err := e.Repository.Save(ctx, newElement)

	if err != nil {
		return id, err
	}

	return id, nil
}
