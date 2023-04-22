package log

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 默认的一些配置
func DefaultEncoderConfig() zapcore.EncoderConfig {

	var encodingConfig = zap.NewProductionEncoderConfig()
	encodingConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encodingConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	return encodingConfig
}

// 统一用 json
func DefaultEncoder() zapcore.Encoder {

	return zapcore.NewJSONEncoder(DefaultEncoderConfig())
}

func DefaultOption() []zap.Option {

	var stackTraceLevel zap.LevelEnablerFunc = func(l zapcore.Level) bool {
		return l >= zapcore.DPanicLevel
	}

	return []zap.Option{
		zap.AddCaller(),
		zap.AddStacktrace(stackTraceLevel),
	}
}

// 1. 不会自动清理 backup
// 2. 每 200 mb 压缩一次，不按时间 rotate
func DefaultLumberjackLogger() *lumberjack.Logger {

	return &lumberjack.Logger{
		MaxSize:   200,
		LocalTime: true,
		Compress:  true,
	}
}
