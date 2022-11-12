package util_test

import (
	"testing"
	"time"

	"github.com/MatheusHenrique129/application-in-go/internal/util"
	"github.com/stretchr/testify/assert"
)

func TestCurrentUtcDateTime(t *testing.T) {
	assert.NotNil(t, util.NewTimeHelper().GetCurrentUtcDateTime())
}

func TestSleep(t *testing.T) {
	before := time.Now()
	ms := int64(2)

	util.NewTimeHelper().Sleep(time.Duration(ms) * time.Millisecond)

	after := time.Now()

	assert.True(t, after.Sub(before).Milliseconds() >= ms)
}

func TestGetCurrentUnixMilliseconds(t *testing.T) {
	timeHelper := util.NewTimeHelper()
	timeOld := timeHelper.GetCurrentUnixMilliseconds()
	timeHelper.Sleep(time.Duration(2) * time.Millisecond)
	timeNew := timeHelper.GetCurrentUnixMilliseconds()
	assert.GreaterOrEqual(t, timeNew, timeOld)
}
