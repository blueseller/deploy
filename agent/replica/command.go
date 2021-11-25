package replica

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/blueseller/deploy.git/agent/replica/action"
	commandPb "github.com/blueseller/deploy.git/api/agent/command/v1"
	typesPb "github.com/blueseller/deploy.git/api/agent/types"
	"github.com/blueseller/deploy.git/internal/local"
	"github.com/blueseller/deploy.git/logger"
	"google.golang.org/grpc"
)

var cmdResultCh chan *commandPb.Cmd

func initSendChannel() {
	cmdResultCh = make(chan *commandPb.Cmd, 1000)
}

func PutCmdResult(msg *commandPb.Cmd) {
	cmdResultCh <- msg
}

func StartCommand(ctx context.Context, conn *grpc.ClientConn) error {
	client := commandPb.NewStreamCommandSerivceClient(conn)
	for {
		// send msg
		stream, err := client.Command(ctx)
		if err != nil {
			logger.GetContextLogger(ctx).Errorf("get command stream is error:%s", err.Error())
			time.Sleep(10 * time.Second)
			continue
		}

		initSendChannel()

		wg := sync.WaitGroup{}
		wg.Add(2)
		go func() {
			receive(ctx, wg, stream)
			wg.Done()
			close(cmdResultCh)
		}()

		go func() {
			send(ctx, wg, stream)
			wg.Done()
		}()

		//  send a register cmd
		registerCmd := getRegisterCmd()
		PutCmdResult(registerCmd)
		wg.Wait()
	}
	return nil
}

func getRegisterCmd() *commandPb.Cmd {
	cmd := &commandPb.Cmd{
		AgentId: &typesPb.AgentId{Ip: local.GetLocalIP()},
		CmdType: commandPb.CmdType_CLIENT_REGISTER,
	}
	return cmd
}

func receive(ctx context.Context, wg sync.WaitGroup, stream commandPb.StreamCommandSerivce_CommandClient) {
	var err error
	for {
		select {
		default:
			var cmd *commandPb.Cmd
			cmd, err = stream.Recv()
			if err != nil {
				goto EXIT
			}
			logger.GetContextLogger(ctx).Infof("cmd: ", cmd.AgentId, cmd.CmdId, string(cmd.Payload))

			// do this cmd
			execCmd(ctx, cmd)
		}
	}

EXIT:
	logger.GetContextLogger(ctx).Errorf("recv msg is close, error:%v", err)
}

func send(ctx context.Context, wg sync.WaitGroup, stream commandPb.StreamCommandSerivce_CommandClient) {
	var err error
	for {
		select {
		//	case <-sendCloseCh:
		//		logger.GetContextLogger(ctx).Warnf("send goroutinue is finish, please retry to connect.")
		//		goto EXIT
		case cmd := <-cmdResultCh:
			// send result
			err = stream.Send(cmd)
			if err != nil {
				goto EXIT
			}
		}
	}
EXIT:
	logger.GetContextLogger(ctx).Errorf("send msg is close, error: %v", err)
}

// TODO
func execCmd(ctx context.Context, cmd *commandPb.Cmd) error {
	switch cmd.CmdType {
	case commandPb.CmdType_LOG_AGGREGATE:
		logAction := action.NewLogAction()
		list := logAction.GetLogCatalogList(ctx, "")

		result := &commandPb.Result{
			ErrCode: 0,
			OutPut:  string(marshalStructToBytes(list)),
		}
		cmd.Result = result
		PutCmdResult(cmd)
	}
	return nil
}

func marshalStructToBytes(val interface{}) []byte {
	u, _ := json.Marshal(val)
	return u
}
