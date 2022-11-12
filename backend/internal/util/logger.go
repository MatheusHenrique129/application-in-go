package util

import (
	"context"
	"fmt"

	"github.com/MatheusHenrique129/application-in-go/libraries/logger"
	"github.com/MatheusHenrique129/application-in-go/libraries/logging"
)

type Logger struct {
	component string
}

func (l *Logger) Info(ctx context.Context, message string) {
	tags := l.getTags(ctx)
	msg := l.getBaseMessage(message)

	logger.Info(msg, tags...)
}

func (l *Logger) Debug(ctx context.Context, message string) {
	tags := l.getTags(ctx)
	msg := l.getBaseMessage(message)

	logger.Debug(msg, tags...)
}

func (l *Logger) Warn(ctx context.Context, message string) {
	tags := l.getTags(ctx)
	msg := l.getBaseMessage(message)

	logger.Warn(msg, tags...)
}

func (l *Logger) Error(ctx context.Context, message string, err error) {
	tags := l.getTags(ctx)
	msg := l.getBaseMessage(message)

	logger.Error(msg, err, tags...)
}

func (l *Logger) Infof(ctx context.Context, message string, args ...interface{}) {
	msg := fmt.Sprintf(message, args...)
	l.Info(ctx, msg)
}

func (l *Logger) Debugf(ctx context.Context, message string, args ...interface{}) {
	msg := fmt.Sprintf(message, args...)
	l.Debug(ctx, msg)
}

func (l *Logger) Warnf(ctx context.Context, message string, args ...interface{}) {
	msg := fmt.Sprintf(message, args...)
	l.Warn(ctx, msg)
}

func (l *Logger) Errorf(ctx context.Context, message string, err error, args ...interface{}) {
	msg := fmt.Sprintf(message, args...)
	l.Error(ctx, msg, err)
}

func (l *Logger) Panicf(message string, err error, args ...interface{}) {
	logger.Panicf(l.getBaseMessage(message), err, args...)
}

/**
Logs without context can only be used in app package.
*/

func (l *Logger) ErrorWithoutContext(message string, err error) {
	l.Error(context.TODO(), message, err)
}

func (l *Logger) DebugfWithoutContext(message string, args ...interface{}) {
	l.Debugf(context.TODO(), message, args...)
}

func (l *Logger) DebugWithoutContext(message string) {
	l.Debug(context.TODO(), message)
}

func (l *Logger) InfofWithoutContext(message string, args ...interface{}) {
	l.Infof(context.TODO(), message, args...)
}

func (l *Logger) getBaseMessage(message string) string {
	return fmt.Sprintf("[%s] %s", l.component, message)
}

func (l *Logger) getTags(ctx context.Context) []string {
	tags := make([]string, 0)

	requestContext := GetRequestContext(ctx)

	// adding tags to be logged
	if requestContext != nil {
		tags = append(tags, fmt.Sprintf("request_id:%s", requestContext.GetRequestID()))
	}

	return tags
}

func (l *Logger) Logf(classification logging.Classification, message string, args ...interface{}) {
	switch classification {
	case logging.Warn:
		l.Warnf(context.Background(), message, args...)
	default:
		l.Infof(context.Background(), message, args...)
	}
}

func NewLogger(component string) *Logger {
	return &Logger{
		component: component,
	}
}
