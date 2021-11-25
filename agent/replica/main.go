package replica

import (
	"context"

	"github.com/blueseller/deploy.git/logger"
	"google.golang.org/grpc"
)

func Run(ctx context.Context, addr string) {
	// 创建连接
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		logger.GetContextLogger(ctx).Fatalf("connenct agent master error:%s", err.Error())
	}

	logger.GetContextLogger(ctx).Debugf("create connecting this %s", addr)
	// Dispatch some capacity
	Dispatch(ctx, conn)
}

func Dispatch(ctx context.Context, conn *grpc.ClientConn) error {
	//1. 心跳连接
	//StartHeartBeat(ctx, conn)

	//2. command
	StartCommand(ctx, conn)

	//3. 服务器信息上报

	return nil
}
