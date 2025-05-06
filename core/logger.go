package core

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/gorm/logger"
	"os"
	"path/filepath"
	"rbac_manager/global"
	"time"
)

/*
自定义日志模块：
	1.时间格式化
	2.输出美化
	3.预设字段
	4.使用全局日志
	5.日志双写
	6.日志分片
*/

// 全局日志实例

// InitLogger 初始化zap日志配置
func InitLogger(logDir string) {
	// 创建日志目录
	if logDir == "" {
		logDir = "logs"
	}
	if err := os.MkdirAll(logDir, 0755); err != nil {
		panic("无法创建日志目录: " + err.Error())
	}

	// 1. 配置日志输出切割规则
	logFile := filepath.Join(logDir, "app-"+time.Now().Format("2006-01-02")+".log")
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logFile, // 日志文件路径
		MaxSize:    100,     // 单个文件最大尺寸(MB)
		MaxBackups: 30,      // 保留旧文件的最大个数
		MaxAge:     7,       // 保留旧文件的最大天数
		Compress:   true,    // 压缩/归档旧文件
		LocalTime:  true,    // 使用本地时间
	}

	// 2. 配置日志编码器
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// 3. 创建多端输出（文件 + 控制台）
	fileWriter := zapcore.AddSync(lumberJackLogger)
	consoleWriter := zapcore.Lock(os.Stdout)

	// 创建多输出同步器
	multiWriteSyncer := zapcore.NewMultiWriteSyncer(
		fileWriter,
		consoleWriter, // 生产环境可移除
	)

	// 4. 创建核心配置
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig), // JSON格式编码器
		multiWriteSyncer,                      // 多端输出
		zap.InfoLevel,                         // 日志级别
	)

	// 5. 构建Logger实例
	global.Log = zap.New(core,
		zap.AddCaller(),                   // 记录调用信息,即代码行号
		zap.AddCallerSkip(1),              // 包装函数调用层级
		zap.AddStacktrace(zap.ErrorLevel), // 错误级别记录堆栈
	)
	global.Log.Info("init logger success")
}

// ZapGormLogger 适配gorm日志
/*
// 实现以下接口即可
//	type Interface interface {
//		LogMode(LogLevel) Interface
//		Info(context.Context, string, ...interface{})
//		Warn(context.Context, string, ...interface{})
//		Error(context.Context, string, ...interface{})
//		Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error)
//	}
*/

// ZapGormLogger 实现 gorm logger.Interface
type ZapGormLogger struct {
	ZapLogger     *zap.Logger
	SlowThreshold time.Duration
}

// NewZapGormLogger 创建适配 zap 的 GORM 日志器
func NewZapGormLogger(slowThreshold time.Duration) *ZapGormLogger {
	return &ZapGormLogger{
		ZapLogger:     global.Log,
		SlowThreshold: slowThreshold,
	}
}

// LogMode 实现 logger.Interface 的 LogMode 方法
func (l *ZapGormLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *l
	return &newLogger
}

// Info 实现 logger.Interface 的 Info 方法
func (l *ZapGormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	l.ZapLogger.Info(msg, zap.Any("data", data))
}

// Warn 实现 logger.Interface 的 Warn 方法
func (l *ZapGormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	l.ZapLogger.Warn(msg, zap.Any("data", data))
}

// Error 实现 logger.Interface 的 Error 方法
func (l *ZapGormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	l.ZapLogger.Error(msg, zap.Any("data", data))
}

// Trace 实现 logger.Interface 的 Trace 方法
func (l *ZapGormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()

	fields := []zap.Field{
		zap.String("sql", sql),
		zap.Int64("rows", rows),
		zap.Duration("elapsed", elapsed),
	}

	switch {
	case err != nil:
		l.ZapLogger.Error("gorm_trace_error", append(fields, zap.Error(err))...)
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0:
		l.ZapLogger.Warn("gorm_slow_query", fields...)
	default:
		l.ZapLogger.Debug("gorm_trace", fields...)
	}
}

// 日志双写参考方法
// 创建文件日志核心
func createFileCore(filename string) zapcore.Core {
	fileWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   filename,
		MaxSize:    100, // MB
		MaxBackups: 30,
		MaxAge:     7, // days
		Compress:   true,
	})

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	return zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		fileWriter,
		zap.InfoLevel,
	)
}

// 创建控制台日志核心
func createConsoleCore(level zapcore.Level) zapcore.Core {
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	return zapcore.NewCore(
		consoleEncoder,
		zapcore.Lock(os.Stdout),
		level,
	)
}
