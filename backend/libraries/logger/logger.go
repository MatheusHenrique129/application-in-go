package logger

import (
	"fmt"
	"io"
	"os"
	"runtime/debug"
	"strings"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

const tagMessageFormat = "%s - %s"

func init() {
	Log = &logrus.Logger{
		Out:       os.Stdout,
		Formatter: &MercuryFormatter{DisableColors: true},
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.DebugLevel,
	}
}

func SetLogLevel(logLevel string) {
	if level, error := logrus.ParseLevel(logLevel); error != nil {
		panic(error)
	} else {
		Log.Level = level
	}
}

func Print(e interface{}) {
	Log.Printf("%s: %s", e, debug.Stack())
}

func Debug(message string, tags ...string) {
	if Log.Level >= logrus.DebugLevel {
		tags = append(tags, "level:debug")
		entry, message := buildLogEntryWithFieldsAndMessage(tags, message)

		entry.Debug(message)
	}
}

func Info(message string, tags ...string) {
	if Log.Level >= logrus.InfoLevel {
		tags = append(tags, "level:info")
		entry, message := buildLogEntryWithFieldsAndMessage(tags, message)

		entry.Info(message)
	}
}

func Warn(message string, tags ...string) {
	if Log.Level >= logrus.WarnLevel {
		tags = append(tags, "level:warn")
		entry, message := buildLogEntryWithFieldsAndMessage(tags, message)

		entry.Warn(message)
	}
}

func Error(message string, err error, tags ...string) {
	if Log.Level >= logrus.ErrorLevel {
		tags = append(tags, "level:error")

		msg := fmt.Sprintf("%s - ERROR: %v", message, err)
		entry, msg := buildLogEntryWithFieldsAndMessage(tags, msg)

		entry.Error(msg)
	}
}

func Panic(message string, err error, tags ...string) {
	if Log.Level >= logrus.PanicLevel {
		tags = append(tags, "level:panic")

		msg := fmt.Sprintf("%s - PANIC: %v", message, err)
		entry, msg := buildLogEntryWithFieldsAndMessage(tags, msg)

		entry.Panic(msg)
	}
}

func Debugf(format string, args ...interface{}) {
	if Log.Level >= logrus.DebugLevel {
		Debug(fmt.Sprintf(format, args...))
	}
}

func Infof(format string, args ...interface{}) {
	if Log.Level >= logrus.InfoLevel {
		Info(fmt.Sprintf(format, args...))
	}
}

func Warnf(format string, args ...interface{}) {
	if Log.Level >= logrus.WarnLevel {
		Warn(fmt.Sprintf(format, args...))
	}
}

func Errorf(format string, err error, args ...interface{}) {
	if Log.Level >= logrus.ErrorLevel {
		Error(fmt.Sprintf(format, args...), err)
	}
}

func Panicf(format string, err error, args ...interface{}) {
	if Log.Level >= logrus.PanicLevel {
		Panic(fmt.Sprintf(format, args...), err)
	}
}

func GetOut() io.Writer {
	return Log.Out
}

func buildLogEntryWithFieldsAndMessage(tags []string, message string) (*logrus.Entry, string) {
	fields, err := getFields(tags)
	if err != nil {
		message = fmt.Sprintf(tagMessageFormat, message, err.Error())
	}

	return Log.WithFields(fields), message
}

func getFields(tags []string) (logrus.Fields, error) {
	fields := make(logrus.Fields)
	wrongTags := []string{}

	var err error

	for _, value := range tags {
		values := strings.SplitN(value, ":", 2)

		if len(values) < 2 {
			wrongTags = append(wrongTags, value)
			continue
		}

		fields[strings.TrimSpace(values[0])] = strings.TrimSpace(values[1])
	}

	if len(wrongTags) > 0 {
		err = fmt.Errorf("Error parsing tags (%s)", strings.Join(wrongTags, ","))
	}

	return fields, err
}
