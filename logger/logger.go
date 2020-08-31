package logger

import (
	"go.uber.org/zap" // 日志包
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2" // 日志包
)

var l *zap.Logger

/*
我们先定义一个包级的全局变量l，
类型是*zap.Logger，
并创建了InitLogger和GetLogger这两个函数。因为，
zap不支持日志归档
所以在InitLogger中定义了一个lumberjack的hook，
用来归档日志。
我们可以看到InitLogger这个方法有两个入参：logPath和logLevel。
一般来讲，这些参数应该是放在配置文件里的，
接下来我们来写配置。
*/


func InitLogger(logPath, logLevel string) error  {
	hook := lumberjack.Logger{

		Filename: logPath,
		MaxSize: 1024,
		MaxBackups: 3,
		MaxAge: 7,
		Compress: true,
	}
	w := zapcore.AddSync(&hook) // 添加同步

	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.DebugLevel
	}

	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encodeConfig),
		w,
		level,
		)
	l = zap.New(core)
	return nil
}

func GetLogger() *zap.Logger  {
	return l
}