package http

import (
	"context"
	"go-backend/conf"
)

type contextKey string

const (
	configKey    = contextKey("config")
	componentKey = contextKey("component")
	requestIdKey = contextKey("request_id")
)

func (ck contextKey) String() string {
	return "http context value " + string(ck)
}

func WithConfig(ctx context.Context, config *conf.Configuration) context.Context {
	return context.WithValue(ctx, configKey, config)
}

func GetConfig(ctx context.Context) *conf.Configuration {
	obj := ctx.Value(configKey)
	if obj == nil {
		return nil
	}
	return obj.(*conf.Configuration)
}

func WithRequestId(ctx context.Context, requestId string) context.Context {
	return context.WithValue(ctx, requestIdKey, requestId)
}

func GetRequestId(ctx context.Context) string {
	obj := ctx.Value(requestIdKey)
	if obj == nil {
		return ""
	}
	return obj.(string)
}

func WithComponent(ctx context.Context, component string) context.Context {
	return context.WithValue(ctx, componentKey, component)
}

func GetComponent(ctx context.Context) string {
	obj := ctx.Value(componentKey)
	if obj == nil {
		return ""
	}
	return obj.(string)
}
