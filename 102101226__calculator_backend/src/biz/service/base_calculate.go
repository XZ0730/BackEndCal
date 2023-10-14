package service

import (
	"context"

	"github.com/XZ0730/tireCV/biz/dal/cache"
	"github.com/XZ0730/tireCV/utils"
)

func (c *Do_CalculateService) BaseCalculate(ctx context.Context, exp string) (string, error) {
	return utils.Calculate(exp)
}

func (c *Do_CalculateService) History(ctx context.Context, key string) (cache.HistoryList, error) {
	return cache.History(ctx, key)
}
