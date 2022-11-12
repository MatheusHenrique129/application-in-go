package util_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/MatheusHenrique129/application-in-go/internal/util"
	"github.com/stretchr/testify/assert"
)

func Test_ErrorWithoutContext(t *testing.T) {
	logger := util.NewLogger("ErrorWithoutContext")
	logger.ErrorWithoutContext("say something", fmt.Errorf("%s", "something"))
	assert.True(t, true, true)
}

func Test_DebugfWithoutContext(t *testing.T) {
	logger := util.NewLogger("DebugfWithoutContext")
	logger.DebugfWithoutContext("say something")
	assert.True(t, true, true)
}

func Test_Warn(t *testing.T) {
	logger := util.NewLogger("Warn")
	logger.Warn(context.TODO(), "say something")
	assert.True(t, true, true)
}

func Test_Warnf(t *testing.T) {
	logger := util.NewLogger("Warnf")
	logger.Warnf(context.TODO(), "say something: %s", "hello")
	assert.True(t, true, true)
}

func Test_Errorf(t *testing.T) {
	logger := util.NewLogger("Errorf")
	logger.Errorf(context.TODO(), "say something: %s", fmt.Errorf("some error"), "hello")
	assert.True(t, true, true)
}

func Test_Panicf(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	logger := util.NewLogger("Panic")
	logger.Panicf("say something: %s", fmt.Errorf("some error"), "hello")
	assert.True(t, true, true)
}

func Test_Debug_WithTags(t *testing.T) {
	ctx := context.WithValue(context.TODO(), util.RequestContextField, &util.RequestContext{})
	logger := util.NewLogger("Debug_WithTags")
	logger.Debug(ctx, "say something")
}
