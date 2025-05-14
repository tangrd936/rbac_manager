package core

import (
	"fmt"
	"go.uber.org/zap"
	"rbac_manager/global"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitDB  初始化数据库连接
func InitDB() {
	// 构建 DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/rbac_manager?charset=utf8mb4&parseTime=True&loc=Local",
		global.Conf.Db.Username,
		global.Conf.Db.Password,
		global.Conf.Db.Host,
		global.Conf.Db.Port,
	)

	// 创建数据库连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                                   NewZapGormLogger(100 * time.Millisecond), // 配置 GORM 日志器
		PrepareStmt:                              true,                                     // 预处理
		DisableForeignKeyConstraintWhenMigrating: true,                                     // 不生成实体外键
	})
	if err != nil {
		global.Log.Error("InitDB() failed to connect database", zap.Error(err))
	}

	// 获取通用数据库对象以设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		global.Log.Error("InitDB() failed to get generic database object", zap.Error(err))
	}

	// 连接池配置
	sqlDB.SetMaxIdleConns(10)           // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)          // 最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Hour) // 连接的最大可复用时间

	// 测试连接是否成功
	if err = sqlDB.Ping(); err != nil {
		global.Log.Error("InitDB() database ping failed", zap.Error(err))
	}

	// 设置全局数据库连接
	global.Db = db
	global.Log.Info("init db success")
}
