package chief

import (
	"context"
	"net"

	commandPb "github.com/blueseller/deploy.git/api/agent/command/v1"
	"github.com/blueseller/deploy.git/dcontext"
	"github.com/blueseller/deploy.git/internal/local"
	"github.com/blueseller/deploy.git/logger"
	"google.golang.org/grpc"
)

func Run(ctx context.Context) {
	ip := dcontext.AgentMasterIp(ctx)
	port := dcontext.AgentMasterPort(ctx)

	// 获取本机ip地址
	if ip == "" {
		ip = local.GetLocalIP()
	}

	if ip == "" {
		logger.GetContextLogger(ctx).Fatalf("获取本机IP地址失败,请使用参数传入一个IP地址 ")
	}
	dcontext.WithAgentMasterIp(ctx, ip)

	server := grpc.NewServer()

	cmdSrv := NewCommandService(ctx)

	commandPb.RegisterStreamCommandSerivceServer(server, cmdSrv)

	logger.GetContextLogger(ctx).Debugf("agent master start linten addr is %s", ip+":"+port)
	lis, err := net.Listen("tcp", ip+":"+port)
	//lis, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		logger.GetContextLogger(ctx).Fatalf("net.Listen err: %v,addr is %s", err, ip+":"+port)
	}

	logger.GetContextLogger(ctx).Debugf("listening addr %s is success", ip+":"+port)
	go cmdSrv.TestCase()

	server.Serve(lis)
}
