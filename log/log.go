package log

import (
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Plugin = zapcore.Core

// NOTE: 一些 option 选项是无法覆盖的
func NewLogger(plugin zapcore.Core, options ...zap.Option) *zap.Logger {

	return zap.New(plugin, append(DefaultOption(), options...)...)
}

func NewPlugin(writer zapcore.WriteSyncer, enabler zapcore.LevelEnabler) Plugin {

	return zapcore.NewCore(DefaultEncoder(), writer, enabler)
}

func NewStdoutPlugin(enabler zapcore.LevelEnabler) Plugin {

	return NewPlugin(zapcore.Lock(zapcore.AddSync(os.Stdout)), enabler)
}

func NewStderrPlugin(enabler zapcore.LevelEnabler) Plugin {

	return NewPlugin(zapcore.Lock(zapcore.AddSync(os.Stderr)), enabler)
}

// Lumberjack logger 虽然持有 File 但没有暴露 sync 方法, 所以没办法利用 zap 的 sync 特性
// 所以额外返回一个 closer, 需要保证在进程退出前 close 以保证写入的内容可以全部刷到磁盘
func NewFilePlugin(filePath string, enabler zapcore.LevelEnabler) (Plugin, io.Closer) {

	var writer = DefaultLumberjackLogger()
	writer.Filename = filePath

	return NewPlugin(zapcore.AddSync(writer), enabler), writer
}
