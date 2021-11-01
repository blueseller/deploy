package dcontext

import (
	"context"

	"github.com/blueseller/deploy/configure"
)

type commandKey struct{}
type commandUpStepKey struct{}

func WithCommand(ctx context.Context, command string) {
	ctx = context.WithValue(ctx, commandKey{}, command)

}

func Command(ctx context.Context) interface{} {
	return ctx.Value(commandKey{})
}

func WithCommandUpStep(ctx context.Context, upstep int) {
	ctx = context.WithValue(ctx, commandKey{}, upstep)
}

func CommandUpStep(ctx context.Context) configure.CmdStep {
	val := ctx.Value(commandKey{})
	upStep, ok := val.(configure.CmdStep)
	if ok {
		return upStep
	}
	return configure.CmdStep(0)
}
