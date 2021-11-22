package dcontext

import "context"

type agentMasterIpKey struct{}
type agentMasterPortKey struct{}

func WithAgentMasterIp(ctx context.Context, ip string) context.Context {
	return context.WithValue(ctx, agentMasterIpKey{}, ip)
}

func WithAgentMasterPort(ctx context.Context, port string) context.Context {
	return context.WithValue(ctx, agentMasterPortKey{}, port)
}
