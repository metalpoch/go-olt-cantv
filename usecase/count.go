package usecase

import (
	"context"
	"database/sql"
	"time"

	"github.com/metalpoch/go-olt-cantv/model"
	"github.com/metalpoch/go-olt-cantv/repository"
)

type CountUsecase struct {
	Repository repository.CountRepository
}

func NewCountUsecase(repo repository.CountRepository) *CountUsecase {
	return &CountUsecase{
		Repository: repo,
	}
}

func (d CountUsecase) Add(currCount model.Count) (model.CountDiff, error) {
	countDiff := model.CountDiff{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	prevCount, err := d.Repository.FindPreviouCount(ctx, currCount.ElementID)

	if err != sql.ErrNoRows && err != nil {
		return countDiff, err
	}

	if err != sql.ErrNoRows {
		countDiff = model.CountDiff{
			ElementID:     currCount.ElementID,
			PrevDate:      prevCount.Date,
			PrevBytesIn:   prevCount.BytesIn,
			PrevBytesOut:  prevCount.BytesOut,
			CurrDate:      currCount.Date,
			CurrBytesIn:   currCount.BytesIn,
			CurrBytesOut:  currCount.BytesOut,
			CurrBandwidth: currCount.Bandwidth,
		}

		if err = d.Repository.Remove(ctx, currCount.ElementID); err != nil {
			return countDiff, err
		}

	}

	_, err = d.Repository.Save(ctx, currCount)

	if err != nil {
		return countDiff, err
	}

	return countDiff, nil
}
