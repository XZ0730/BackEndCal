package service

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/XZ0730/tireCV/biz/dal/cache"
	"github.com/XZ0730/tireCV/biz/model/calculate"
	"github.com/XZ0730/tireCV/utils"
	"github.com/cloudwego/kitex/pkg/klog"
)

type RateType int

const (
	StoreMoneyThreeMonth int = iota
	StoreMoneySixMonth
	StoreMoneyOneYear
	StoreMoneyTwoYear
	StoreMoneyThreeYear
	StoreMoneyFiveYear
)

const (
	ProvideMoneySixMonth int = iota
	ProvideMoneyOneYear
	ProvideMoneyOne2Three
	ProvideMoneyThree2Five
	ProvideMoneyFiveYear
)

func (c *Do_CalculateService) RateCalculate(ctx context.Context, req *calculate.RateRequest) (string, error) {

	switch req.Type {
	case 0, 1:
		// TODO:存款利息
		// 		查询数据库得到利息表
		//		获取利息表后得到利息的信息比对后根据存款计算利息
		//
		rate_list, err := cache.GetRateList(ctx, int(req.Type))
		if err != nil {
			klog.Error("[store ratelist] ratelist get error:", err.Error())
			return "", err
		}
		rate, ok := rate_list[int(req.Duration)]
		if !ok {
			klog.Error("[ratelist] store:", rate_list)
			return "", errors.New("[ratelist] no such type")
		}
		money, err := strconv.ParseFloat(req.Money, 64)
		if err != nil {
			klog.Error("[store parse] money is error，money:", req.Money)
			return "", errors.New("[store parse] money is error")
		}
		interest, _ := utils.StoreRateToMoney(money, rate)
		return fmt.Sprint(interest), nil

	default:
		return "", errors.New("[type] no such type select")
	}
}

func (c *Do_CalculateService) SetRate(ctx context.Context, req *calculate.SetRateRequest) error {
	return cache.SetRateList(ctx, int(req.TheType), req.Data)
}

func (c *Do_CalculateService) GetRate(ctx context.Context) ([]string, error) {

	return cache.GetRate(ctx)
}
