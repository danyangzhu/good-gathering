package log

import (
	"good_gathering/conf"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *Logger

// error logger
type Logger struct {
	Zap      *zap.Logger
	ZapSugar *zap.SugaredLogger
	AppLevel *AtomicLevel
	Level    string
}

func (l *Logger) Print(v ...interface{}) {
	l.ZapSugar.Info(v)
}

func (l *Logger) SetAppLevel(levelStr string) {
	if l.Level == levelStr {
		return
	}
	Info("setloglevel", zap.String("level", levelStr))
	level := getLoggerLevel(levelStr)
	if l.AppLevel != nil {
		l.AppLevel.Level = level
	}

}

type AtomicLevel struct {
	Level zapcore.Level
}

func (a *AtomicLevel) Enabled(lvl zapcore.Level) bool {
	//return lvl < zapcore.ErrorLevel && lvl >= a.Level
	return lvl >= a.Level
}

var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func getLoggerLevel(lvl string) zapcore.Level {
	if level, ok := levelMap[lvl]; ok {
		return level
	}
	return zapcore.InfoLevel
}

//添加日志调用者函数名
func CallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	funcPath := runtime.FuncForPC(caller.PC).Name()
	funcPaths := strings.Split(funcPath, "/")
	funcName := funcPaths[len(funcPaths)-1:][0]
	enc.AppendString("[" + caller.TrimmedPath() + "][" + funcName + "]")
}

//修改日志等级格式
func CapitalLevelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + l.CapitalString() + "]")
}

//修改时间格式
func ISO8601TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + t.Format("2006-01-02T15:04:05.000Z0700") + "]")
}

func FullCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + caller.String() + "][" + runtime.FuncForPC(caller.PC).Name() + "]")
}

func lumberjackHook(logfile string, maxSize int, maxBackups int, maxAge int, compress bool) lumberjack.Logger {
	return lumberjack.Logger{
		Filename:   logfile,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   compress,
	}
}

func NewConsoleCore(level zapcore.Level) zapcore.Core {
	encoderConfig := zapcore.EncoderConfig{
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    CapitalLevelEncoder,
		EncodeTime:     ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   FullCallerEncoder,
	}

	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	writer := zapcore.AddSync(os.Stdout)
	priority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= level
	})

	return zapcore.NewCore(encoder, writer, priority)
}

func NewAppCore(hook lumberjack.Logger, atomiclevel *AtomicLevel) zapcore.Core {
	encoderConfig := zapcore.EncoderConfig{
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    CapitalLevelEncoder,
		EncodeTime:     ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   CallerEncoder,
	}

	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	writer := zapcore.AddSync(&hook)

	return zapcore.NewCore(encoder, writer, atomiclevel)
}

func NewErrCore(hook lumberjack.Logger) zapcore.Core {
	encoderConfig := zapcore.EncoderConfig{
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    CapitalLevelEncoder,
		EncodeTime:     ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   FullCallerEncoder,
	}

	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	writer := zapcore.AddSync(&hook)

	priority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})

	return zapcore.NewCore(encoder, writer, priority)
}

//初始化，需要在入口函数调用，确保最先被初始化
func Init() {
	logConf := conf.NeedUse("log")
	logPath := logConf.MustString("path", "./log")
	levelStr := logConf.MustString("level", "info")
	appLogName := logConf.MustString("appLogName", "app.log")
	errLogName := logConf.MustString("errLogName", "error.log")
	maxSize := logConf.MustInt("maxSize", 100)
	maxBackups := logConf.GetInt("maxBackups")
	maxAge := logConf.MustInt("maxAge", 7)
	compress := logConf.MustBool("compress", true)
	showConsole := logConf.MustBool("showConsole", false)
	level := getLoggerLevel(levelStr)
	atomiclevel := &AtomicLevel{
		Level: level,
	}

	var core zapcore.Core

	appFile := filepath.Join(filepath.FromSlash(logPath), appLogName)
	errFile := filepath.Join(filepath.FromSlash(logPath), errLogName)

	appHook := lumberjackHook(appFile, maxSize, maxBackups, maxAge, compress)
	errHook := lumberjackHook(errFile, maxSize, maxBackups, maxAge, compress)

	appCore := NewAppCore(appHook, atomiclevel)
	errCore := NewErrCore(errHook)

	//输出到控制台
	if showConsole == true {
		core = zapcore.NewTee(
			appCore,
			errCore,
			NewConsoleCore(level),
		)
	} else {
		core = zapcore.NewTee(
			appCore,
			errCore,
		)
	}

	zapLogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	logger = &Logger{
		Zap:      zapLogger,
		ZapSugar: zapLogger.Sugar(),
		AppLevel: atomiclevel,
		Level:    levelStr,
	}

	//注册日志接收配置变更消息，日志级别热加载
	conf.Register(OnConfigChange)
}

func GetLogger() *Logger {
	return logger
}
