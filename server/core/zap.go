package core

import (
	"fmt"
	zaprotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"myblog/global"
	"myblog/utils"
	"os"
	"path"
	"time"
)

var level zapcore.Level

// Zap 创建日志实例
func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.GCONFIG.Zap.Director); !ok {
		fmt.Printf("create %v directory\n", global.GCONFIG.Zap.Director)
		_ = os.Mkdir(global.GCONFIG.Zap.Director, os.ModePerm)
	}
	switch global.GCONFIG.Zap.Level {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}
	if level == zap.DebugLevel || level == zap.ErrorLevel {
		logger = zap.New(getEncoderCore(), zap.AddStacktrace(level))
	} else {
		logger = zap.New(getEncoderCore())
	}
	if global.GCONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

// getEncoderConfig 配置日志
func getEncoderConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  global.GCONFIG.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch {
	case global.GCONFIG.Zap.EncodeLevel == "LowercaseLevelEncoder":
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case global.GCONFIG.Zap.EncodeLevel == "LowercaseColorLevelEncoder":
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case global.GCONFIG.Zap.EncodeLevel == "CapitalLevelEncoder":
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case global.GCONFIG.Zap.EncodeLevel == "CapitalColorLevelEncoder":
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

// getEncoder 日志编码方式
func getEncoder() zapcore.Encoder {
	if global.GCONFIG.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

// getEncoderCore 获取zapcore.Core
func getEncoderCore() (core zapcore.Core) {
	writer, err := getWriteSyncer()
	if err != nil {
		fmt.Printf("get write syncer failed: %v\n", err)
	}
	return zapcore.NewCore(getEncoder(), writer, level)
}

// CustomTimeEncoder 日志格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(global.GCONFIG.Zap.Prefix + "2006-01-02 15:04:05"))
}

// getWriteSyncer 获取日志输出位置
func getWriteSyncer() (zapcore.WriteSyncer, error) {
	fileWriter, err := zaprotatelogs.New(
		// 日志文件
		path.Join(global.GCONFIG.Zap.Director, "%Y-%m-%d.log"),
		// 日志软连接
		zaprotatelogs.WithLinkName(global.GCONFIG.Zap.LinkName),
		// 日志保存时间
		zaprotatelogs.WithMaxAge(global.GCONFIG.Zap.MaxAge*time.Hour),
		// 日志切割时间
		zaprotatelogs.WithRotationTime(global.GCONFIG.Zap.RotationTime*time.Hour),
	)
	if global.GCONFIG.Zap.LogInConsole {
		// 同时输出文件和控制台
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	// 只输出文件
	return zapcore.AddSync(fileWriter), err
}
