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

func (e ElementsUsecase) Find(element model.Element) (model.Element, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := e.Repository.Find(ctx, element)
	if err != nil {
		return model.Element{}, err
	}

	return result, nil
}

func (e ElementsUsecase) Save(newElement model.Element) (model.Element, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id, err := e.Repository.Save(ctx, newElement)

	if err != nil {
		return newElement, err
	}

	newElement.ID = uint(id)
	return newElement, nil
}
