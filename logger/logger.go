package logger

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/blueseller/deploy.git/configure"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

var defaultLogLib = "logrus"

type Logger interface {
	Trace(args ...interface{})
	Tracef(format string, args ...interface{})

	Debug(args ...interface{})
	Debugf(format string, args ...interface{})

	Info(args ...interface{})
	Infof(format string, args ...interface{})

	Warn(args ...interface{})
	Warnf(format string, args ...interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})

	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
}

func LoggerFactory(ctx context.Context, config *configure.Configuration) (context.Context, error) {
	switch config.Log.LogType {
	default:
		return configLogrusLogger(ctx, config)
	}
}

func configLogrusLogger(ctx context.Context, config *configure.Configuration) (context.Context, error) {
	// 设定日志级别
	log.SetLevel(logLevelChange(config.Log.LogLevel))

	log.SetOutput(os.Stdout)

	// 设定日志格式
	logForamtter := config.Log.Formatter
	if logForamtter == "" {
		logForamtter = "text"
	}

	switch logForamtter {
	case "text":
		log.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: time.RFC3339,
		})
	case "json":
		log.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat:   time.RFC3339,
			DisableHTMLEscape: true,
		})
	default:
		if config.Log.Formatter != "" {
			return ctx, fmt.Errorf("unexported log formatter: %q", config.Log.Formatter)
		}
	}

	fields := config.Log.Fields
	if len(fields) > 0 {
		var keys []interface{}
		for key := range fields {
			keys = append(keys, key)
		}

		WithContextLogger(ctx, GetContextLogger(ctx, keys))
	}

	SetDefaultLogger(GetContextLogger(ctx))
	return ctx, nil
}

func logLevelChange(level configure.LogLevel) logrus.Level {
	l, err := logrus.ParseLevel(string(level))
	if err != nil {
		l = logrus.InfoLevel
		logrus.Warnf("log level is not support, use log level is info to default, %q", l)
	}
	return l
}
