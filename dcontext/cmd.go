package dcontext

import (
	"context"

	"github.com/blueseller/deploy.git/configure"
)

type commandKey struct{}
type commandUpStepKey struct{}

func WithCommand(ctx context.Context, command configure.Command) context.Context {
	return context.WithValue(ctx, commandKey{}, command)
}

func Command(ctx context.Context) configure.Command {
	val := ctx.Value(commandKey{})
	if v, ok := val.(configure.Command); ok {
		return v
	}
	return configure.Command{}
}

func WithCommandStep(ctx context.Context, step configure.CmdStep) context.Context {
	cmdSteps := CommandStep(ctx)
	if cmdSteps == nil {
		cmdSteps = make([]configure.CmdStep, 0)
	}
	cmdSteps = append(cmdSteps, configure.CmdStep(step))
	return context.WithValue(ctx, commandUpStepKey{}, cmdSteps)
}

func CommandStep(ctx context.Context) []configure.CmdStep {
	val := ctx.Value(commandUpStepKey{})
	upStep, ok := val.([]configure.CmdStep)
	if ok {
		return upStep
	}
	return nil
}

func CommandLastStep(ctx context.Context) configure.CmdStep {
	val := ctx.Value(commandUpStepKey{})
	upStep, ok := val.([]configure.CmdStep)
	if ok {
		return upStep[len(upStep)-1]
	}
	return configure.CmdStep(0)
}
