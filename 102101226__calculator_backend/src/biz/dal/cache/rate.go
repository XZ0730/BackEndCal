package cache

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
)

// type 0 存款利率
// type 1 贷款利率
func GetRateList(ctx context.Context, the_type int) (map[int]float64, error) {
	res, err := RedisDB.HGet(ctx, "rate", fmt.Sprint("v", the_type)).Result()
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	res = strings.ReplaceAll(res, "'", "\"")
	klog.Info("rate_list :", res)
	rate_map := make(map[int]float64, 0)
	err = sonic.Unmarshal([]byte(res), &rate_map)
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	return rate_map, nil
}
func SetRateList(ctx context.Context, the_type int, data string) error {
	if err := RedisDB.HSet(ctx, "rate", fmt.Sprint("v", the_type), data).Err(); err != nil {
		return err
	}
	RedisDB.Expire(ctx, "rate", time.Hour*24*30)
	return nil
}

// TODO: 历史记录设计 get set

func GetRate(ctx context.Context) ([]string, error) {
	str := make([]string, 2)
	res, err := RedisDB.HGet(ctx, "rate", "v0").Result()
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	res = strings.ReplaceAll(res, "'", "\"")
	klog.Info("rate_list :", res)
	str[0] = res
	res, err = RedisDB.HGet(ctx, "rate", "v1").Result()
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	res = strings.ReplaceAll(res, "'", "\"")
	str[1] = res
	klog.Info("rate_list :", res)
	return str, nil
}
