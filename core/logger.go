package core

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
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

	// ----------------------------
	// 5. 构建Logger实例
	// ----------------------------
	global.Log = zap.New(core,
		zap.AddCaller(),                   // 记录调用信息,即代码行号
		zap.AddCallerSkip(1),              // 包装函数调用层级
		zap.AddStacktrace(zap.ErrorLevel), // 错误级别记录堆栈
	)
}
