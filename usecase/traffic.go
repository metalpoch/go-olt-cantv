package usecase

import (
	"context"
	"time"

	"github.com/metalpoch/go-olt-cantv/model"
	helper "github.com/metalpoch/go-olt-cantv/pkg"
	"github.com/metalpoch/go-olt-cantv/repository"
)

type TrafficUsecase struct {
	Repository repository.TrafficRepository
}

func NewTrafficUsecase(repo repository.TrafficRepository) *TrafficUsecase {
	return &TrafficUsecase{
		Repository: repo,
	}
}

func (m TrafficUsecase) Add(count model.CountDiff) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	diffDate := count.CurrDate - count.PrevDate
	traffic := model.Traffic{
		ElementID: count.ElementID,
		Date:      count.CurrDate,
		KpbsIn:    helper.BytesToKbps(count.PrevBytesIn, count.CurrBytesIn, diffDate),
		KpbsOut:   helper.BytesToKbps(count.PrevBytesOut, count.CurrBytesOut, diffDate),
		Bandwidth: count.CurrBandwidth,
	}

	id, err := m.Repository.Save(ctx, traffic)
	if err != nil {
		return id, err
	}

	return id, nil
}
