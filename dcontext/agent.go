package dcontext

import (
	"context"
)

type agentMasterIpKey struct{}
type agentMasterPortKey struct{}

func WithAgentMasterIp(ctx context.Context, ip string) context.Context {
	return context.WithValue(ctx, agentMasterIpKey{}, ip)
}

func AgentMasterIp(ctx context.Context) string {
	val := ctx.Value(agentMasterIpKey{})
	if v, ok := val.(string); ok {
		return v
	}
	return ""
}

func WithAgentMasterPort(ctx context.Context, port string) context.Context {
	return context.WithValue(ctx, agentMasterPortKey{}, port)
}

func AgentMasterPort(ctx context.Context) string {
	val := ctx.Value(agentMasterPortKey{})
	if v, ok := val.(string); ok {
		return v
	}
	return ""
}
