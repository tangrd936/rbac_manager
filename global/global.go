package global

import (
	"go.uber.org/zap"
	"rbac_manager/config"
)

var (
	// 全局配置
	Conf *config.Config
	// 全局日志配置
	Log *zap.Logger
)
