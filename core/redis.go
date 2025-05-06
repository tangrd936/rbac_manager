package core

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"rbac_manager/global"
	"strconv"
)

func InitRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     global.Conf.Redis.Host + ":" + strconv.Itoa(global.Conf.Redis.Port), // Redis 服务器地址
		Password: global.Conf.Redis.Password,                                          // 没有密码则留空
		DB:       0,                                                                   // 使用默认数据库
	})

	// 测试连接
	pong, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		global.Log.Error(fmt.Sprintf("load redis fail: %v", err))
		return
	}
	global.Redis = rdb
	global.Log.Info(fmt.Sprintf("load redis success: %v", pong))
}
