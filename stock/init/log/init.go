package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"stockPicker/stock/init/config"
	"stockPicker/stock/init/helper"
	"time"
)

var (
	level   zapcore.Level
	options []zap.Option
)

func New(c *config.Config) *zap.Logger {
	// create log root directory
	createRootDir(c)

	// set log level
	setLogLevel(c)

	if c.Log.ShowLine {
		options = append(options, zap.AddCaller())
	}

	// init zap
	logger := zap.New(getZapCore(c), options...)
	zap.ReplaceGlobals(logger)
	return logger
}

func createRootDir(c *config.Config) {
	if ok, _ := helper.PathExists(c.Log.RootDir); !ok {
		_ = os.Mkdir(c.Log.RootDir, os.ModePerm)
	}
}

func setLogLevel(c *config.Config) {
	switch c.Log.Level {
	case "debug":
		level = zap.DebugLevel
		options = append(options, zap.AddStacktrace(level))
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
		options = append(options, zap.AddStacktrace(level))
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}
}

func getZapCore(c *config.Config) zapcore.Core {
	var encoder zapcore.Encoder

	// adjust encoder default settings
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format("[" + "2006-01-02 15:04:05.000" + "]"))
	}
	encoderConfig.EncodeLevel = func(l zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(c.App.Env + "." + l.String())
	}

	// set encoder
	if c.Log.Format == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	return zapcore.NewCore(encoder, getLogWriter(c), level)
}

func getLogWriter(c *config.Config) zapcore.WriteSyncer {
	file := &lumberjack.Logger{
		Filename:   c.Log.RootDir + "/" + c.Log.Filename,
		MaxSize:    c.Log.MaxSize,
		MaxBackups: c.Log.MaxBackups,
		MaxAge:     c.Log.MaxAge,
		Compress:   c.Log.Compress,
	}

	return zapcore.AddSync(file)
}
