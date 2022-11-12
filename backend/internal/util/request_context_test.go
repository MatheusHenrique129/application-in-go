package util_test

import (
	"context"
	"testing"

	"github.com/MatheusHenrique129/application-in-go/internal/util"
	"github.com/stretchr/testify/assert"
)

func Test_GetRequestContext_IfNilReturnNil(t *testing.T) {
	assert.Nil(t, util.GetRequestContext(nil))
}

func Test_GetRequestContext_IfNotNilReturnContext(t *testing.T) {
	ctx := context.WithValue(context.TODO(), util.RequestContextField, &util.RequestContext{})
	assert.NotNil(t, util.GetRequestContext(ctx))
}

func Test_GetRequestID_IsEmpty(t *testing.T) {
	ctx := context.WithValue(context.TODO(), util.RequestContextField, &util.RequestContext{})
	rCtx := util.GetRequestContext(ctx)
	assert.Empty(t, rCtx.GetRequestID())
}
