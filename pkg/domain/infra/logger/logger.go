package infra

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	defaultLog "log"
	"os"
	"strings"
	"sync"
	"ted-processor/pkg/domain/infra/errors"
)

var (
	once      sync.Once
	singleton *zap.SugaredLogger
	Logger    = Get()
)

func Get() *zap.SugaredLogger {
	once.Do(func() {
		log, err := newLogger()
		if err != nil {
			defaultLog.Panicf("Can't initialize logger: %v", err)
		}
		singleton = log.Sugar()
		singleton.Debug("Created new logger")
	})

	return singleton
}

func newLogger() (*zap.Logger, error) {
	zap.NewDevelopmentConfig()
	config := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
	if lvl, exists := os.LookupEnv("LOG_LEVEL"); exists {
		lvl = strings.ToLower(lvl)
		switch lvl {
		case "debug", "trace":
			config.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
		case "info":
			config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
		case "warn":
			config.Level = zap.NewAtomicLevelAt(zapcore.WarnLevel)
		case "error":
			config.Level = zap.NewAtomicLevelAt(zapcore.ErrorLevel)
		case "panic":
			config.Level = zap.NewAtomicLevelAt(zapcore.PanicLevel)
		case "fatal":
			config.Level = zap.NewAtomicLevelAt(zapcore.FatalLevel)
		}
	}
	log, err := config.Build()
	return log, errors.Wrap(err)
}
