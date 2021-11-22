package replica

import (
	"context"
	"fmt"
	"time"

	commandPb "github.com/blueseller/deploy.git/api/agent/command/v1"
	"github.com/blueseller/deploy.git/logger"
	"google.golang.org/grpc"
)

var cmdResultCh chan *commandPb.Cmd

func init() {
	cmdResultCh = make(chan *commandPb.Cmd, 1000)
}

func GetCmdResult() *commandPb.Cmd {
	return <-cmdResultCh
}

func PutCmdResult(msg *commandPb.Cmd) {
	cmdResultCh <- msg
}

func StartCommand(ctx context.Context, conn *grpc.ClientConn) error {
	client := commandPb.NewStreamCommandSerivceClient(conn)
	for {
		// send msg
		reqCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		stream, err := client.Command(reqCtx)
		if err != nil {
			logger.GetContextLogger(ctx).Fatalf("get command stream is error:%s", err.Error())
		}

		go receive(ctx, stream)

		go send(ctx, stream)
	}
	return nil
}

func receive(ctx context.Context, stream commandPb.StreamCommandSerivce_CommandClient) error {
	var err error
	for {
		var cmd *commandPb.Cmd
		cmd, err = stream.Recv()
		if err != nil {
			break
		}

		// do this cmd
		logger.GetContextLogger(ctx).Infof("cmd: ", cmd.AgentId, cmd.CmdId, string(cmd.Payload))
	}

	return fmt.Errorf("recv msg is error %s", err.Error())
}

func send(ctx context.Context, stream commandPb.StreamCommandSerivce_CommandClient) error {
	var err error
	for {
		res := GetCmdResult()
		// send result
		err = stream.Send(res)
		if err != nil {
			break
		}
	}
	return fmt.Errorf("send msg is error %s", err.Error())
}
