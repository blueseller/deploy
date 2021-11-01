package logger

import (
	"context"
	"fmt"
	"runtime"
	"sync"

	"github.com/sirupsen/logrus"
)

type loggerKey struct{}

var (
	defaultLogger    *logrus.Entry = logrus.StandardLogger().WithField("go.version", runtime.Version)
	defaultLoggerMux sync.RWMutex
)

func WithContextLogger(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}
func GetLoggerWithField(ctx context.Context, key, value interface{}, keys ...string) Logger {
	return getLogrusLogger(ctx, keys).WithField(fmt.Sprint(key), value)
}

func GetLoggerWithFields(ctx context.Context, fields map[interface{}]interface{}, keys ...string) Logger {
	flds := make(logrus.Fields, len(fields))
	for k, v := range fields {
		flds[fmt.Sprint(k)] = v
	}
	return getLogrusLogger(ctx, keys).WithFields(flds)
}

func GetContextLogger(ctx context.Context, keys ...interface{}) Logger {
	return getLogrusLogger(ctx, keys)
}

func getLogrusLogger(ctx context.Context, keys ...interface{}) *logrus.Entry {
	var logEntry *logrus.Entry

	// 判断是否已存在一个logger
	loggerInterface := ctx.Value(loggerKey{})
	if loggerInterface != nil {
		if lger, ok := loggerInterface.(*logrus.Entry); ok {
			logEntry = lger
		}
	}

	if logEntry == nil {
		fields := logrus.Fields{}
		instanceId := ctx.Value("instant.id")
		if instanceId != nil {
			fields["instant.id"] = instanceId
		}

		defaultLoggerMux.Lock()
		logEntry = defaultLogger.WithFields(fields)
		defaultLoggerMux.Unlock()
	}

	fields := logrus.Fields{}
	for _, key := range keys {
		fieldVal := ctx.Value(key)
		if fieldVal != nil {
			fields[fmt.Sprint(key)] = fieldVal
		}
	}
	return logEntry.WithFields(fields)
}

func SetDefaultLogger(logger interface{}) {
	entry, ok := logger.(*logrus.Entry)
	if !ok {
		return
	}
	defaultLoggerMux.Lock()
	defaultLogger = entry
	defaultLoggerMux.Unlock()
}
