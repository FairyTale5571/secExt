package logger

import (
	"context"
	"github.com/fairytale5571/secExt/pkg/ip"
	"os"
	"runtime"
	"time"

	"github.com/evalphobia/logrus_sentry"
	"github.com/fairytale5571/secExt/pkg/helpers"
	"github.com/sirupsen/logrus"
)

type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})

	DebugCtx(ctx context.Context, args ...interface{})
	InfoCtx(ctx context.Context, args ...interface{})
	WarnCtx(ctx context.Context, args ...interface{})
	ErrorCtx(ctx context.Context, args ...interface{})
	FatalCtx(ctx context.Context, args ...interface{})

	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})

	DebugfCtx(ctx context.Context, format string, args ...interface{})
	InfofCtx(ctx context.Context, format string, args ...interface{})
	WarnfCtx(ctx context.Context, format string, args ...interface{})
	ErrorfCtx(ctx context.Context, format string, args ...interface{})
	FatalfCtx(ctx context.Context, format string, args ...interface{})
}

type Wrapper struct {
	lg    *logrus.Logger
	entry *logrus.Entry
}

func New(service string) *Wrapper {
	log := &Wrapper{
		lg: logrus.New(),
	}

	hook, err := logrus_sentry.NewSentryHook("https://75311a6f34fd40bbb8cf762330b75eb5@o482351.ingest.sentry.io/5982259", []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
	})
	if err == nil {
		log.lg.Hooks.Add(hook)
		hook.Timeout = 1 * time.Second
	}

	log.lg.SetFormatter(&logrus.JSONFormatter{})
	log.lg.SetOutput(os.Stdout)
	log.lg.SetLevel(logrus.DebugLevel)
	log.entry = log.lg.WithFields(logrus.Fields{
		"service": service,
		"arch":    runtime.GOARCH,
		"isAdmin": helpers.IsAdmin(),
		"ip":      ip.GetIp(),
	})
	return log
}

func (logger *Wrapper) Debug(args ...interface{}) {
	logger.entry.Debug(args...)
}

func (logger *Wrapper) Info(args ...interface{}) {
	logger.entry.Info(args...)
}

func (logger *Wrapper) Warn(args ...interface{}) {
	logger.entry.Warn(args...)
}

func (logger *Wrapper) Error(args ...interface{}) {
	logger.entry.Error(args...)
}

func (logger *Wrapper) Fatal(args ...interface{}) {
	logger.entry.Fatal(args...)
}

func (logger *Wrapper) DebugCtx(ctx context.Context, args ...interface{}) {
	logger.entry.Debug(args...)
}

func (logger *Wrapper) InfoCtx(ctx context.Context, args ...interface{}) {
	logger.entry.Info(args...)
}

func (logger *Wrapper) WarnCtx(ctx context.Context, args ...interface{}) {
	logger.entry.Warn(args...)
}

func (logger *Wrapper) ErrorCtx(ctx context.Context, args ...interface{}) {
	logger.entry.Error(args...)
}

func (logger *Wrapper) FatalCtx(ctx context.Context, args ...interface{}) {
	logger.entry.Fatal(args...)
}

func (logger *Wrapper) Debugf(format string, args ...interface{}) {
	logger.entry.Debugf(format, args...)
}

func (logger *Wrapper) Infof(format string, args ...interface{}) {
	logger.entry.Infof(format, args...)
}

func (logger *Wrapper) Warnf(format string, args ...interface{}) {
	logger.entry.Warnf(format, args...)
}

func (logger *Wrapper) Errorf(format string, args ...interface{}) {
	logger.entry.Errorf(format, args...)
}

func (logger *Wrapper) Fatalf(format string, args ...interface{}) {
	logger.entry.Fatalf(format, args...)
}

func (logger *Wrapper) DebugfCtx(ctx context.Context, format string, args ...interface{}) {
	logger.entry.Debugf(format, args...)
}

func (logger *Wrapper) InfofCtx(ctx context.Context, format string, args ...interface{}) {
	logger.entry.Infof(format, args...)
}

func (logger *Wrapper) WarnfCtx(ctx context.Context, format string, args ...interface{}) {
	logger.entry.Warnf(format, args...)
}

func (logger *Wrapper) ErrorfCtx(ctx context.Context, format string, args ...interface{}) {
	logger.entry.Errorf(format, args...)
}

func (logger *Wrapper) FatalfCtx(ctx context.Context, format string, args ...interface{}) {
	logger.entry.Fatalf(format, args...)
}
