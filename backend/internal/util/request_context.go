package util

import (
	"context"
)

// Constants declaration
const (
	RequestContextField = "RequestContext"
	RequestIDHeader     = "X-Request-Id"
)

// RequestContext structure consolidates context data related to each individual request that is handled by the API
type RequestContext struct {
	requestID      string
	metricBaseName string
	logger         *Logger
}

// GetRequestContext retrieves the RequestContext attached to the given context
func GetRequestContext(ctx context.Context) *RequestContext {
	if ctx == nil {
		return nil
	}

	reqCtx := ctx.Value(RequestContextField)

	if reqCtx == nil {
		return nil
	}

	return reqCtx.(*RequestContext)
}

// GetRequestID returns the ID of the request associates to this context
func (reqContext *RequestContext) GetRequestID() string {
	return reqContext.requestID
}
