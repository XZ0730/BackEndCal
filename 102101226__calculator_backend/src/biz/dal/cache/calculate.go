package cache

import (
	"context"
	"time"
)

type HistoryList []string

func History(ctx context.Context, key string) (HistoryList, error) {
	return RedisDB.HVals(ctx, key).Result()
}

// key ip地址 field 时间戳 val 表达式
func HistoryTrimAndSet(ctx context.Context, key string, field string) error {
	return RedisDB.HDel(ctx, key, field).Err()
}

func SetHistory(ctx context.Context, key, field, exp string) error {
	if err := RedisDB.HSet(ctx, key, field, exp).Err(); err != nil {
		return err
	}
	RedisDB.Expire(ctx, key, time.Hour*24*30)
	return nil
}
