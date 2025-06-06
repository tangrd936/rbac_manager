package global

import (
	"github.com/casbin/casbin/v2"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"rbac_manager/config"
)

var (
	// Conf 全局配置
	Conf *config.Config
	// Log 全局日志配置
	Log *zap.Logger
	// Db 数据库连接
	Db *gorm.DB
	// Redis redis连接
	Redis *redis.Client
	// Casbin 配置对象
	Casbin *casbin.CachedEnforcer
)
