package dcontext

import "context"

var defaultCtx context.Context

func init() {
	defaultCtx = context.Background()
}

func SetDContext(ctx context.Context) {
	defaultCtx = ctx
}

func GetDContext() context.Context {
	return defaultCtx
}
