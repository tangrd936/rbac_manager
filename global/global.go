package global

import (
	"go.uber.org/zap"
	"rbac_manager/config"
)

var (
	// Conf 全局配置
	Conf *config.Config
	// Log 全局日志配置
	Log *zap.Logger
)
