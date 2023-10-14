package service

import (
	"context"

	"github.com/XZ0730/tireCV/biz/dal/cache"
	"github.com/XZ0730/tireCV/biz/model/calculate"
)

type Calculate interface {
	BaseCalculate(ctx context.Context, exp string) (string, error)
	RateCalculate(ctx context.Context, req *calculate.RateRequest) (string, error)
	SetRate(ctx context.Context, req *calculate.SetRateRequest) error
	History(ctx context.Context, key string) (cache.HistoryList, error)
	GetRate(ctx context.Context) ([]string, error)
}

type Do_CalculateService struct {
	context context.Context
}

func NewDo_CalculateService(ctx context.Context) Calculate {
	return &Do_CalculateService{
		context: ctx,
	}
}
