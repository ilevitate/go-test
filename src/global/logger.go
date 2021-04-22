package global

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var Logger *zap.Logger

func SetLogger() (err error) {
	section := Config.Section("log")
	level := section.Key("level").String()
	logFile := section.Key("logFile").String()
	encode := section.Key("encode").String()

	config := zap.NewProductionConfig()
	switch level {
	case "debug":
		config.Level.SetLevel(zap.DebugLevel)
	case "info":
		config.Level.SetLevel(zap.InfoLevel)
	case "warn":
		config.Level.SetLevel(zap.WarnLevel)
	case "error":
		config.Level.SetLevel(zap.ErrorLevel)
	case "panic":
		config.Level.SetLevel(zap.PanicLevel)
	case "final":
		config.Level.SetLevel(zap.FatalLevel)
	}

	if logFile != "" {
		config.OutputPaths = []string{logFile} //日志文件输出路径，为空则输出到控制台
	}

	config.Encoding = encode
	config.EncoderConfig.StacktraceKey = "source" //定义源码文件项的键名
	//定义日期格式
	config.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}
	Logger, err = config.Build() //使用自定义的配置构建Logger
	return
}
