package dal

import "github.com/XZ0730/tireCV/biz/dal/cache"

func Init() {
	cache.InitRedis()
}
