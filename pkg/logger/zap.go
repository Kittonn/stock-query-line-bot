package logger

import (
	"os"

	"github.com/Kittonn/stock-query-line-bot/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Debug(args ...any)
	Info(args ...any)
	Warn(args ...any)
	Error(args ...any)
	Fatal(args ...any)
}

var loggerLevelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func getZapLogLevel(level string) zapcore.Level {
	if zapLevel, ok := loggerLevelMap[level]; ok {
		return zapLevel
	}

	return zapcore.InfoLevel
}

type zapLogger struct {
	sugar *zap.SugaredLogger
}

var _ Logger = (*zapLogger)(nil)

func NewZapLogger(config *config.Config) Logger {
	logLevel := getZapLogLevel(config.LogLevel)

	logWriter := zapcore.AddSync(os.Stderr)
	var encoderCfg zapcore.EncoderConfig
	if config.AppEnv == "development" {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderCfg = zap.NewProductionEncoderConfig()
	}

	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.LevelKey = "level"
	encoderCfg.CallerKey = "caller"
	encoderCfg.MessageKey = "message"
	encoderCfg.StacktraceKey = "stacktrace"
	encoderCfg.NameKey = "name"
	encoderCfg.TimeKey = "time"
	encoderCfg.EncodeDuration = zapcore.MillisDurationEncoder

	var encoder zapcore.Encoder
	if config.LogEncoder == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	}

	core := zapcore.NewCore(encoder, logWriter, zap.NewAtomicLevelAt(logLevel))

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	return &zapLogger{sugar: logger.Sugar()}
}

// Debug implements Logger.
func (z *zapLogger) Debug(args ...any) {
	z.sugar.Debug(args...)
}

// Error implements Logger.
func (z *zapLogger) Error(args ...any) {
	z.sugar.Error(args...)
}

// Fatal implements Logger.
func (z *zapLogger) Fatal(args ...any) {
	z.sugar.Fatal(args...)
}

// Info implements Logger.
func (z *zapLogger) Info(args ...any) {
	z.sugar.Info(args...)
}

// Warn implements Logger.
func (z *zapLogger) Warn(args ...any) {
	z.sugar.Warn(args...)
}
