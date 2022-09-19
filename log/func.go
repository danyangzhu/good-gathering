package log

import (
	"good_gathering/conf"

	"go.uber.org/zap"
)

func Debug(msg string, fields ...zap.Field) {
	logger.Zap.Debug(msg, fields...)
}

func Debugf(template string, args ...interface{}) {
	logger.ZapSugar.Debugf(template, args...)
}

func Info(msg string, fields ...zap.Field) {
	logger.Zap.Info(msg, fields...)
}

func Infof(template string, args ...interface{}) {
	logger.ZapSugar.Infof(template, args...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Zap.Warn(msg, fields...)
}

func Warnf(template string, args ...interface{}) {
	logger.ZapSugar.Warnf(template, args...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Zap.Error(msg, fields...)
}

func Errorf(template string, args ...interface{}) {
	logger.ZapSugar.Errorf(template, args...)
}

func DPanic(msg string, fields ...zap.Field) {
	logger.Zap.DPanic(msg, fields...)
}

func DPanicf(template string, args ...interface{}) {
	logger.ZapSugar.DPanicf(template, args...)
}

func Panic(msg string, fields ...zap.Field) {
	logger.Zap.Panic(msg, fields...)
}

func Panicf(template string, args ...interface{}) {
	logger.ZapSugar.Panicf(template, args...)
}

func Fatal(msg string, fields ...zap.Field) {
	logger.Zap.Fatal(msg, fields...)
}

func Fatalf(template string, args ...interface{}) {
	logger.ZapSugar.Fatalf(template, args...)
}

//注册日志接收配置变更消息，热变更日志级别
func OnConfigChange(name string, op uint32) {
	levelStr := conf.Use("log").MustString("level", "info")
	logger.SetAppLevel(levelStr)
}
