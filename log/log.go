package log

import (
	"good_gathering/util"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type YGinLog interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
}

type Log struct {
	RequestId string
}

func NewLog(rid string) *Log {
	return &Log{
		RequestId: "[" + rid + "] ",
	}
}

func New(c *gin.Context) *Log {
	return Get(c)
}

func Get(c *gin.Context) *Log {
	if l, exist := c.Get("Logger"); exist {
		return l.(*Log)
	}

	return NewLog(util.GetRequestId(c))
}

func (l *Log) Print(v ...interface{}) {
	GetLogger().ZapSugar.Info(l.RequestId, v)
}

func (l *Log) Printf(format string, v ...interface{}) {
	GetLogger().ZapSugar.Info(l.RequestId+format, v)
}

func (l *Log) Debug(fields ...zap.Field) {
	GetLogger().Zap.Debug(l.RequestId, fields...)
}

func (l *Log) Debugf(template string, args ...interface{}) {
	GetLogger().ZapSugar.Debugf(l.RequestId+template, args...)
}

func (l *Log) Info(fields ...zap.Field) {
	GetLogger().Zap.Info(l.RequestId, fields...)
}

func (l *Log) Infof(template string, args ...interface{}) {
	GetLogger().ZapSugar.Infof(l.RequestId+template, args...)
}

func (l *Log) Warn(fields ...zap.Field) {
	GetLogger().Zap.Warn(l.RequestId, fields...)
}

func (l *Log) Warnf(template string, args ...interface{}) {
	GetLogger().ZapSugar.Warnf(l.RequestId+template, args...)
}

func (l *Log) Error(fields ...zap.Field) {
	GetLogger().Zap.Error(l.RequestId, fields...)
}

func (l *Log) Errorf(template string, args ...interface{}) {
	GetLogger().ZapSugar.Errorf(l.RequestId+template, args...)
}

func (l *Log) DPanic(fields ...zap.Field) {
	GetLogger().Zap.DPanic(l.RequestId, fields...)
}

func (l *Log) DPanicf(template string, args ...interface{}) {
	GetLogger().ZapSugar.DPanicf(l.RequestId+template, args...)
}

func (l *Log) Panic(fields ...zap.Field) {
	GetLogger().Zap.Panic(l.RequestId, fields...)
}

func (l *Log) Panicf(template string, args ...interface{}) {
	GetLogger().ZapSugar.Panicf(l.RequestId+template, args...)
}

func (l *Log) Fatal(fields ...zap.Field) {
	GetLogger().Zap.Fatal(l.RequestId, fields...)
}

func (l *Log) Fatalf(template string, args ...interface{}) {
	GetLogger().ZapSugar.Fatalf(l.RequestId+template, args...)
}
