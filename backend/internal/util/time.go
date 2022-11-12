package util

import (
	"time"
)

// TimeHelper wraps time.Time. Useful for mock time.Time.
type TimeHelper interface {
	GetCurrentUtcDateTime() time.Time
	Sleep(d time.Duration)
	GetCurrentUnixMilliseconds() int64
}

// timeHelper represents time utilities.
type timeHelper struct{}

// GetCurrentUtcDateTime returns the current UTC time.
func (t *timeHelper) GetCurrentUtcDateTime() time.Time {
	return time.Now().UTC()
}

func (t *timeHelper) Sleep(d time.Duration) {
	time.Sleep(d)
}

func (t *timeHelper) GetCurrentUnixMilliseconds() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// NewTimeHelper constructs a TimeHelper.
func NewTimeHelper() TimeHelper {
	return &timeHelper{}
}
