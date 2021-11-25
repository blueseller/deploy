package chief

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/blueseller/deploy.git/internal/uuid"
	"github.com/blueseller/deploy.git/logger"

	commandPb "github.com/blueseller/deploy.git/api/agent/command/v1"
	typesPb "github.com/blueseller/deploy.git/api/agent/types"
)

type ConnID string

var sendCmdCh chan *commandPb.Cmd

type CommandService struct {
	ctx        context.Context
	cmdSendMap map[ConnID]commandPb.StreamCommandSerivce_CommandServer
	agentIdMap map[string]ConnID
	mu         sync.RWMutex
}

func NewCommandService(ctx context.Context) *CommandService {
	cmdSrv := new(CommandService)
	cmdSrv.ctx = ctx
	cmdSrv.cmdSendMap = make(map[ConnID]commandPb.StreamCommandSerivce_CommandServer)
	cmdSrv.agentIdMap = make(map[string]ConnID)
	sendCmdCh = make(chan *commandPb.Cmd, 1000)

	go cmdSrv.startSendCmd(ctx)
	return cmdSrv
}

func (c *CommandService) addConnectStream(key string, stream commandPb.StreamCommandSerivce_CommandServer) {
	c.mu.Lock()
	c.cmdSendMap[ConnID(key)] = stream
	c.mu.Unlock()
}

func (c *CommandService) deleteConnectStream(key string) {
	c.mu.Lock()
	delete(c.cmdSendMap, ConnID(key))
	c.mu.Unlock()
}

func (c *CommandService) Command(allStr commandPb.StreamCommandSerivce_CommandServer) error {
	id := uuid.GenUuid()
	c.addConnectStream(id, allStr)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for {
			data, err := allStr.Recv()
			if err != nil {
				logger.GetContextLogger(c.ctx).Errorf("receive msg is error: %s", err.Error())
				goto EXIT
			}

			c.doReceiveMsg(c.ctx, ConnID(id), data)
		}
	EXIT:
		wg.Done()
	}()

	// 存储sendMap
	//go func() {
	//	for {
	//		select {
	//		case cmd := <-cmdCh:
	//			err := allStr.Send(cmd)
	//			if err != nil {
	//				logger.GetContextLogger(c.ctx).Errorf("send msg is error: %s", err.Error())
	//				goto EXIT
	//			}
	//		}
	//	}
	//EXIT:
	//	wg.Done()
	//}()
	wg.Wait()
	c.deleteConnectStream(id)
	return nil
}

func (c *CommandService) doReceiveMsg(ctx context.Context, connID ConnID, data *commandPb.Cmd) {
	switch data.CmdType {
	case commandPb.CmdType_CLIENT_REGISTER:
		var ip string
		agentId := data.GetAgentId()
		if agentId != nil {
			ip = agentId.GetIp()
		}

		if ip != "" {
			c.agentIdMap[ip] = connID
		}
		return
	}
}

func (c *CommandService) startSendCmd(ctx context.Context) {
	for {
		select {
		case cmd := <-sendCmdCh:
			ip := getSendIp(cmd.AgentId)
			if ip == "" {
				logger.GetContextLogger(ctx).Warnf("send cmd is error, send ip address is empty.")
				return
			}

			stream := c.getStreamByIp(ip)
			if stream != nil {
				err := stream.Send(cmd)
				if err != nil {
					logger.GetContextLogger(ctx).Errorf("send msg to ip %s is error : %s", ip, err.Error())
				}
			}
		}
	}
}

func getSendIp(agentId *typesPb.AgentId) string {
	var ip string
	if agentId != nil {
		ip = agentId.GetIp()
	}
	return ip
}

func (c *CommandService) getStreamByIp(ip string) commandPb.StreamCommandSerivce_CommandServer {
	c.mu.RLock()
	defer c.mu.RUnlock()
	connId := c.agentIdMap[ip]
	stream := c.cmdSendMap[connId]
	return stream
}

func Push(cmd *commandPb.Cmd) {
	sendCmdCh <- cmd
}

func (c *CommandService) TestCase() {
	i := 0
	for {
		time.Sleep(10 * time.Second)

		if len(c.agentIdMap) <= 0 {
			continue
		}

		cmdData := new(commandPb.Cmd)
		for ip, connId := range c.agentIdMap {
			agentId := new(typesPb.AgentId)
			agentId.Ip = ip
			cmdData.AgentId = agentId
			cmdData.CmdId = string(connId)
			break
		}

		cmdData.CmdType = commandPb.CmdType_LOG_AGGREGATE
		cmdData.Payload = []byte(fmt.Sprintf("this is num %d", i))

		logger.GetContextLogger(context.TODO()).Debugf("put cmd to client num %d", i)

		Push(cmdData)
		i++
	}
}
