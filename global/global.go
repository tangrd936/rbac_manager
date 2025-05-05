package global

import (
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
)
