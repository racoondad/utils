package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logPrefix = ""
)

func InitDefaultZapLogger() (logger *zap.Logger) {
	return InitZapLogger("json", "log", "logs", true, "LowercaseLevelEncoder", "stacktrace", false, 5, 5000, 365)
}

func InitZapLogger(Format string, Prefix string, Director string, Showline bool, Encodelevel string, Stacktracekey string, Loginconsole bool, maxSize int, maxBackup int, MaxDay int) (logger *zap.Logger) {
	logPrefix = Prefix
	if !IsDirExist(Director) { // 判断是否有Director文件夹
		_ = os.Mkdir(Director, os.ModePerm)
		_ = os.Mkdir(Director+"/debug", os.ModePerm)
		_ = os.Mkdir(Director+"/info", os.ModePerm)
		_ = os.Mkdir(Director+"/warn", os.ModePerm)
		_ = os.Mkdir(Director+"/error", os.ModePerm)
	}
	// 调试级别
	debugPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.DebugLevel
	})
	// 日志级别
	infoPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.InfoLevel
	})
	// 警告级别
	warnPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.WarnLevel
	})
	// 错误级别
	errorPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})

	cores := [...]zapcore.Core{
		// func getEncoderCore(fileLocate string, format string, stacktraceKey string, encodeLevel string, level zapcore.LevelEnabler) (core zapcore.Core)
		getEncoderCore(Loginconsole, fmt.Sprintf("./%s/debug/server_debug.log", Director), Format, Stacktracekey, Encodelevel, debugPriority, maxSize, maxBackup, MaxDay),
		getEncoderCore(Loginconsole, fmt.Sprintf("./%s/info/server_info.log", Director), Format, Stacktracekey, Encodelevel, infoPriority, maxSize, maxBackup, MaxDay),
		getEncoderCore(Loginconsole, fmt.Sprintf("./%s/warn/server_warn.log", Director), Format, Stacktracekey, Encodelevel, warnPriority, maxSize, maxBackup, MaxDay),
		getEncoderCore(Loginconsole, fmt.Sprintf("./%s/error/server_error.log", Director), Format, Stacktracekey, Encodelevel, errorPriority, maxSize, maxBackup, MaxDay),
	}
	logger = zap.New(zapcore.NewTee(cores[:]...), zap.AddCaller())

	if Showline {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

// getEncoderConfig 获取zapcore.EncoderConfig
func getEncoderConfig(stacktraceKey string, encodeLevel string) (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  stacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch encodeLevel {
	case "LowercaseLevelEncoder": // 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case "LowercaseColorLevelEncoder": // 小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case "CapitalLevelEncoder": // 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case "CapitalColorLevelEncoder": // 大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

// getEncoder 获取zapcore.Encoder
func getEncoder(format string, stacktraceKey string, encodeLevel string) zapcore.Encoder {
	if format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig(stacktraceKey, encodeLevel))
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig(stacktraceKey, encodeLevel))
}

// getEncoderCore 获取Encoder的zapcore.Core
func getEncoderCore(InConsole bool, fileRotatelogs, format, stacktraceKey, encodeLevel string, level zapcore.LevelEnabler, maxSize int, maxBackup int, MaxDay int) (core zapcore.Core) {
	writer := getWriteSyncer(fileRotatelogs, InConsole, maxSize, maxBackup, MaxDay) // 使用file-rotatelogs进行日志分割
	return zapcore.NewCore(getEncoder(format, stacktraceKey, encodeLevel), writer, level)
}

// 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(logPrefix + "2006-01-02 15:04:05.000"))
}

func getWriteSyncer(fileRotatelogs string, InConsole bool, maxSize int, maxBackup int, MaxDay int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fileRotatelogs, //日志文件的位置
		MaxSize:    maxSize,        //在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxAge:     MaxDay,         //保留旧文件的最大天数
		MaxBackups: maxBackup,      //保留旧文件的最大个数
		LocalTime:  true,
		Compress:   true, //是否压缩/归档旧文件
	}

	if InConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	}
	return zapcore.AddSync(lumberJackLogger)
}
