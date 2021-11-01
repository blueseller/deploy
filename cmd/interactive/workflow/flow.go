package workflow

import (
	"context"
	"time"

	"github.com/blueseller/deploy/cmdflow"
	"github.com/blueseller/deploy/configure"
	"github.com/sirupsen/logrus"
)

func StartInteractive(ctx context.Context, config *configure.Configuration) {
	flow := cmdflow.NewCmdFlow(config)
	err := flow.InitCmd(ctx, config.CmdFlow)
	if err != nil {
		logrus.Fatalf("init cmd flow data is error, %v", err)
	}
	for {
		// 获取现在可执行的命令
		flow.GetWorkflowCmd(ctx)

		time.Sleep(1 * time.Second)
	}
}
