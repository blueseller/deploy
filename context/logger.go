package context

import (
	"context"
	"fmt"
	"runtime"
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	defaultLogger   *logrus.Entry = logrus.StandardLogger().WithField("go.version", runtime.Version())
	defaultLoggerMu sync.RWMutex
)

// Logger provides a leveled-logging interface.
type Logger interface {
	// standard logger methods
	Print(args ...interface{})
	Printf(format string, args ...interface{})
	Println(args ...interface{})

	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatalln(args ...interface{})

	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
	Panicln(args ...interface{})

	// Leveled methods, from logrus
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Debugln(args ...interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Errorln(args ...interface{})

	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Infoln(args ...interface{})

	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Warnln(args ...interface{})

	WithError(err error) *logrus.Entry
}

type loggerKey struct{}

func WithLogger(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

func GetLoggerWithField(ctx context.Context, key, value interface{}, keys ...interface{}) Logger {
	return getLogrusLogger(ctx, keys...).WithField(fmt.Sprint(key), value)
}

func GetLoggerWithFields(ctx context.Context, fields map[interface{}]interface{}, keys ...interface{}) Logger {
	lfields := make(logrus.Fields, len(fields))
	for key, value := range fields {
		lfields[fmt.Sprint(key)] = value
	}
	return getLogrusLogger(ctx, lfields...)
}

func GetLogger(ctx context.Context, keys ...interface{}) Logger {
	return getLogrusLogger(ctx, keys...)
}

func SetDefaultLogger(logger Logger) {
	entry, ok := logger.(*logrus.Entry)
	if !ok {
		return
	}
	defaultLoggerMu.Lock()
	defaultLogger = entry
	defaultLoggerMu.Unlock()
}

// 获取 logger
func getLogrusLogger(ctx context.Context, keys ...interface{}) *logrus.Entry {
	var logger *logrus.Entry

	loggerInterface := context.Value(loggerKey{})

	if loggerInterface != nil {
		if lgr, ok := loggerInterface.(*logrus.Entry); ok {
			logger = lgr
		}
	}

	if logger == nil {

		fields := logrus.Fields{}
		instanceId := context.Value(ctx, "instance.id")
		if instanceId != nil {
			fields["instance.id"] = instanceId
		}

		defaultLoggerMu.RLock()
		logger = defaultLogger.WithFields(fields)
		defaultLoggerMu.Unlock()
	}

	fields := logrus.Fields{}
	for _, key := range keys {
		v := context.Value(ctx, key)
		if v != nil {
			fields[fmt.Sprint(key)] = ""
		}
	}

	return defaultLogger.WithFields(fields)
}
